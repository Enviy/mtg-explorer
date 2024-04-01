package model

// Card stores information about one single card.
type Card struct {
	// Name defines the name of the front of a card.
	// For split, double-faced and flip cards, the name of only one side.
	// Basically each ‘sub-card’ has its own record.
	Name string `json:"name" db:"name"`
	// Names, only used for split, flip and dual cards.
	// Will contain all the names on this card, front or back.
	Names []string `json:"names" db:"names"`
	// The ManaCost of a card. Consists of one or more mana symbols.
	// Use CMC and Colors to query.
	ManaCost string `json:"manaCost" db:"mana_cost"`
	// Converted mana cost(CMC). Always a number.
	CMC float64 `json:"cmc" db:"cmc"`
	// The card Colors. Usually derived from the casting cost.
	// Except for cards like the back of dual sided cards and Ghostfire.
	Colors []string `json:"colors" db:"colors"`
	// ColorIdentity defines card colors by color code.
	// Ex. [“Red”, “Blue”] becomes [“R”, “U”]
	ColorIdentity []string `json:"colorIdentity" db:"color_identity"`
	// Type defines card type. Seen in type line of printed card.
	// Note: The dash is a UTF8 "long dash" as per MTG rules.
	Type string `json:"type" db:"type"`
	// Types defines multiple entries for Type
	// Seen on left of the dash in a card type. Examples: Instant, Sorcery,\
	// Artifact, Creature, Enchantment, Land, Planeswalker.
	Types []string `json:"types" db:"types"`
	// Supertype. Appears to the far left of the card type.
	// Examples: Basic, Legendary, Snow, World, Ongoing.
	Supertypes []string `json:"supertypes" db:"supertypes"`
	// Subtypes. Appear after long dash following type.
	// Examples: Trap, Arcane, Equipment, Aura, Human, Rat, Squirrel.
	Subtypes []string `json:"subtypes" db:"subtypes"`
	// Rarity of card.
	// Examples: Common, Uncommon, Rare, Mythic Rare, Special, Basic Land.
	Rarity string `json:"rarity" db:"rarity"`
	// Set defines what expansion set the card belongs to by set code.
	Set string `json:"set" db:"set"`
	// SetName defines name of expansion set the card belongs to.
	SetName string `json:"setName" db:"set_name"`
	// Text defines oracle text of card.
	// May contain mana symbols and other symbols.
	// Text defines oracle text of card.
	// May contain mana symbols and other symbols.
	Text string `json:"text" db:"text"`
	// Flavor defines the flavor text of card.
	Flavor string `json:"flavor" db:"flavor"`
	// Artist defines the artist on the card.
	// This may not match the card, MTGJSON corrects card misprints.
	Artist string `json:"artist" db:"artist"`
	// Number defines card's set number. Appears bottom-center of the card.
	// NOTE: Set number can contain letters, strconv to int will error.
	Number string `json:"number" db:"number"`
	// Power defines power of creature cards.
	// NOTE: Power can contain non-int, strconv to int will error.
	Power string `json:"power" db:"power"`
	// Toughness defines toughness of creature cards.
	// NOTE: Toughness can contain non-int, strconv to int will error.
	Toughness string `json:"toughness" db:"toughness"`
	// Loyalty defines loyalty of planeswalker cards.
	Loyalty string `json:"loyalty" db:"loyalty"`
	// MultiverseID defines the ID of the card on Wizard’s Gatherer web page.
	// Cards from sets that do not exist on Gatherer will NOT have a MultiverseID.
	// Sets not on Gatherer: ATH, ITP, DKM, RQS, DPA and all sets with a 4 letter\
	// code that starts with a lowercase 'p’.
	MultiverseID string `json:"multiverseid" db:"multiverse_id"`
	// ImageURL defines URL for card image.
	// NOTE: Only for cards with a MultiverseID.
	ImageURL string `json:"imageUrl" db:"image_url"`
	// Reserved defines if card is reserved by Wizards Reprint Policy.
	Reserved bool `json:"reserved" db:"reserved"`
	// Rulings define rulings for the card.
	Rulings []*Ruling `json:"rulings" db:"rulings"`
	// ID defines unique identification number of the card.
	// ID calculated by SHA1 hash of setCode + cardName + cardImageName.
	ID string `json:"id" db:"id"`
	// Legalities defines formats this card is legal, restricted or banned in.
	// Objects defined as "format" and "legality" keys.
	Legalities []Legality `json:"legalities" db:"legalities"`
}

// Ruling contains additional rule information about the card.
type Ruling struct {
	// Date the information was released.
	Date string `json:"date" db:"date"`
	// Text of the ruling hint.
	Text string `json:"text" db:"text"`
}

// Legality stores information about legality notices for a specific format.
type Legality struct {
	// Format, such as Commander, Standard, Legacy, etc.
	Format string `json:"format" db:"format"`
	// Legality for the given format such as Legal, Banned or Restricted.
	Legality string `json:"legality" db:"legality"`
}
