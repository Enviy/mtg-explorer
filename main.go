package main

import (
	"fmt"

	mtg "github.com/Enviy/mtg-sdk-go"
)

func main() {
	fmt.Println("mtg explorer")
	defer fmt.Println("complete.")

	// Collect sets in standard.
	// standardSets keys are set names, values are set codes.
	standardSets, err := mtg.StandardSets()
	if err != nil {
		panic(err)
	}

	cards, err := mtg.StandardCards()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Number of sets in standard: %v\n\n", len(standardSets))
	for key, value := range standardSets {
		fmt.Println(key, value)
	}

	fmt.Printf("Sample mtg.Card struct:\n%+v\n", cards[0])

	/*
		fmt.Println("Enter set code to get set's cards: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		setName := scanner.Text()
	*/

	/*
		for card := range cards {
			if card.Set == setName {
				fmt.Printf("%+v\n", card)
			}
		}
	*/
}
