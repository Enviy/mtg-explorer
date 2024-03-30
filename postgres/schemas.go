package postgres

// schema intial DB setup, define tables.
var schema = `
CREATE TABLE IF NOT EXISTS card (
    id text unique NOT NULL,
    multiverseid text,
    name text,
    names text[],
    manaCost text,
    cmc float64,
    colors text[],
    colorIdentity text[],
    type text,
    types text[],
    supertypes text[],
    subtypes text[],
    rarity text,
    set text,
    setName text,
    text text,
    flavor text,
    artist text,
    number text,
    power text,
    toughness text,
    loyalty text,
    layout text,
    variations text[],
    imageUrl text,
    border text,
    timeshifted bool,
    rulings text[],
    source text,
    legalities text[]
);
`
