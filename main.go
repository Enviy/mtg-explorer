package main

import (
	"fmt"

	"github.com/Enviy/mtg-explorer/mtg"
)

func main() {
	fmt.Println("mtg explorer")
	defer fmt.Println("complete.")
	standardSets, err := mtg.StandardSets()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Number of sets in standard: %v\n\n", len(standardSets))
	for key, value := range standardSets {
		fmt.Println(key, value)
	}

	// Get standard cards.
	cards, err := mtg.StandardCards()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Cards in standard: %d\n", len(cards))
}
