package hammurabi

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const (
	requiredInput int = 3
)

// InteractiveHammurabi represents the minimal interface for an interactive Hammurabi game.
type InteractiveHammurabi interface {
	DisplayGameState(year int) error
	ReadActionInput(reader *bufio.Reader) (*GameAction, error)
	Hammurabi
}

// DisplayGameState displays textual representation of the game state and state delta.
func (g *Game) DisplayGameState(year int) error {
	// Get the current state from the given year
	if year < 1 && year > g.Year {
		return &ValueOutOfRange{Type: "year", Reason: fmt.Sprintf("Should be within range [%d, %d].", 0, g.Year)}
	}

	// Get the previous delta and the current state
	delta := g.Delta
	state := g.State

	// Display general information
	fmt.Println()
	fmt.Println("Hammurabi: I beg to report to you,")
	fmt.Printf("In Year %d, %d people starved.\n", year, delta.PeopleStarved)
	fmt.Printf("%d people came to the city.\n", delta.PeopleAdded)
	fmt.Printf("The city population is now %d.\n", state.Population)
	fmt.Printf("The city now owns %d acres.\n", state.Lands)
	fmt.Printf("You harvested %d bushels per acre.\n", state.LandProfit)
	if delta.HasRat {
		fmt.Printf("Rats ate %d bushels.\n", delta.BushelsInfested)
	}
	if delta.HasPlague {
		fmt.Printf("Plague killed %d people.\n", delta.PeopleKilled)
	}
	fmt.Printf("You now have %d bushels in store.\n", state.Bushels)
	fmt.Printf("Land is trading at %d bushels per acre.\n", state.LandPrice)

	// No error
	return nil
}

// ReadActionInput reads the input and parse it to GameAction
func (g *Game) ReadActionInput(reader *bufio.Reader) (action *GameAction, err error) {
	fmt.Println()
	fmt.Println("Input your action with the following format:")
	fmt.Println("[LandsToBuy] [BushelsToFeed] [LandsToSeed]")
	text, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	// Initialize game action
	input := strings.Fields(text)
	if len(input) != requiredInput {
		err = &InvalidInput{}
		return
	}

	// Parse the input
	action = &GameAction{}
	action.LandsToBuy, err = strconv.Atoi(input[0])
	if err != nil {
		return
	}
	action.BushelsToFeed, err = strconv.Atoi(input[1])
	if err != nil {
		return
	}
	action.LandsToSeed, err = strconv.Atoi(input[2])
	if err != nil {
		return
	}

	// Otherwise set the action
	g.Action = action
	return
}
