package mtg

import (
	mtg "github.com/Enviy/mtg-sdk-go"
)

// StandardCards returns slice of cards in Standard.
func StandardCards() ([]*mtg.Card, error) {
	// cards is mtg.[]*Card
	return mtg.StandardCards()
}

// GetFormats .
func GetFormats() ([]string, error) {
	return mtg.GetFormats()
}

// StandardSets returns map of sets in Standard.
func StandardSets() (map[string]mtg.SetCode, error) {
	return mtg.StandardSets()
}

// GetSet returns a slice of Card for a given set.
func GetSet(setCode string) ([]*mtg.Card, error) {
	// mtg.CardSet is the set code column.
	// Where(column cardColumn, query string) Query
	query := mtg.NewQuery().Where(mtg.CardSet, setCode)
	return query.All()
}
