// Package postgres provides methods for interacting with the postgres mtg db.
package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Enviy/mtg-explorer/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Gateway defines methods for interacting with postgres.
type Gateway interface {
}

type gateway struct {
	db     *sqlx.DB
	txOpts *sql.TxOptions
}

// New is the Gateway interface constructor.
// Not designd for cloud execution. Intended for CLI.
func New() (Gateway, error) {
	connStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("MTG_POSTGRES_USER"),
		os.Getenv("MTG_POSTGRES_DB"),
		os.Getenv("MTG_POSTGRES_PASSWORD"),
		os.Getenv("MTG_POSTGRES_SSL"),
	)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("sql Open %w", err)
	}

	return &gateway{
		db:     db,
		txOpts: &sql.TxOptions{Isolation: sql.LevelSerializable},
	}, nil
}

// GetCards is a rudementary filter method. It stacks AND statements with
// fields and values provided in parameter f. Does not support LIKE.
func (g *gateway) GetCards(ctx context.Context, f model.Card) ([]model.Card, error) {
	// Convert struct to map string interface.
	queryMap, err := toMap(f)
	if err != nil {
		return nil, fmt.Errorf("toMap %w", err)
	}

	// Build the query.
	query := "SELECT * FROM card"
	var index int
	var prefix string
	for key := range queryMap {
		if index == 0 {
			prefix = "WHERE"
		}
		query += fmt.Sprintf(" %s %s=:%s", prefix, key, key)
		index += 1
	}

	// Begin transaction.
	tx, err := g.db.BeginTxx(ctx, g.txOpts)
	if err != nil {
		return nil, fmt.Errorf("BeginTxx %w", err)
	}
	defer tx.Rollback()

	// Prepare query as built statement.
	preparedStmnt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("PrepareNamedContext %w", err)
	}

	// Execute query with params.
	var cards []model.Card
	if err := preparedStmnt.SelectContext(ctx, &cards, queryMap); err != nil {
		return nil, fmt.Errorf("SelectContext %w", err)
	}

	// End transaction.
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("Commit %w", err)
	}

	return cards, nil
}

// toMap parses a struct to a map accounting for sql.Nullx types.
// Supports using a single struct for reading and writing rows.
func toMap(p interface{}) (map[string]interface{}, error) {
	queryBytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	var queryMap map[string]interface{}
	if err := json.Unmarshal(queryBytes, &queryMap); err != nil {
		return nil, err
	}

	// Resolve map values to underlying values.
	// ONLY handles sql.Nullx type values.
	nullables := []string{
		"String",
		"Bool",
		"Int64",
		"Byte",
		"Float64",
		"Int16",
		"Int32",
		"Time",
	}
	// Parse resulting maps of sql.Nullx types. Remove invalids.
	for key, value := range queryMap {
		if asMap, ok := value.(map[string]interface{}); ok {
			if !asMap["Valid"].(bool) {
				delete(queryMap, key)
				continue
			}

			for _, nullable := range nullables {
				if subVal, ok := asMap[nullable]; ok {
					queryMap[key] = subVal
				}
			}
		}
	}

	return queryMap, nil
}
