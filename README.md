# MTG Explorer

Collecting all of the cards currently in standard can take a long time. It may even exceed a connection's ttl.

MTG Explorer attempts to collect all sets in standard, then collect set cards in parallel.
Collected cards are then stored in a local redis instance which speeds up future operations.

# Goose migration
Example temp env vars followed by goose command.
GOOSE_DRIVER=postgres GOOSE_DBSTRING="user=postgres dbname=postgres sslmode=disable" goose status
