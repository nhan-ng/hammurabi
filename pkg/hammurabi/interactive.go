package hammurabi

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	requiredInput int = 3
)

const (
	intro = `
Congratulations, you are the newest ruler of ancient Samaria, elected for %d-year term of office. Your duties are to dispsense food,
direct farming, and buy and sell land as needed to support your people. Watch out for rat infestations and the plague! Gain is the general
currency, measured in bushels. The following will help you in your decisions:

- Each person needs at least %d bushels of grain per year to survive.
- Each person can farm at most %d acres of land.
- It takes %d bushel of grain to farm an acre of land.
- The mark price for land fluctuates yearly.

Rule wisely and you will be showered with appreciation at the end of your term. Rule poorly and you will be kicked out of office!
	`
)

// InteractiveHammurabi represents the minimal interface for an interactive Hammurabi game.
type InteractiveHammurabi interface {
	DisplayIntro(year int) error
	DisplayGameState(year int) error
	ReadActionInput(reader *bufio.Reader) (*GameAction, error)
	Hammurabi
}

// NewInteractiveHammurabi creates a new game with the maximum number of years, aka turns.
func NewInteractiveHammurabi(maxYear int) InteractiveHammurabi {
	// Initialize a new game
	return newGame(maxYear)
}

// DisplayIntro displays introduction text of the game.
func (g *game) DisplayIntro(year int) error {
	// Validate
	if year < 1 && year > g.year {
		return &valueOutOfRange{kind: "year", reason: fmt.Sprintf("Should be within range [%d, %d].", 0, g.year)}
	}

	fmt.Printf(intro, year, bushelsPerPerson, landsPerPerson, bushelsPerLand)
	return nil
}

// DisplayGameState displays textual representation of the game state and state delta.
func (g *game) DisplayGameState(year int) error {
	// Get the current state from the given year
	if year < 1 && year > g.year {
		return &valueOutOfRange{kind: "year", reason: fmt.Sprintf("Should be within range [%d, %d].", 0, g.year)}
	}

	// Get the previous delta and the current state
	delta := g.delta
	state := g.state

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
func (g *game) ReadActionInput(reader *bufio.Reader) (action *GameAction, err error) {
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
		err = &invalidInput{}
		return
	}

	// Parse the input
	action = &GameAction{}
	action.LandsToBuy, err = strconv.Atoi(input[0])
	if err != nil {
		err = errors.Wrap(err, "validation failed")
		return
	}
	action.BushelsToFeed, err = strconv.Atoi(input[1])
	if err != nil {
		err = errors.Wrap(err, "validation failed")
		return
	}
	action.LandsToSeed, err = strconv.Atoi(input[2])
	if err != nil {
		err = errors.Wrap(err, "validation failed")
		return
	}

	// Otherwise set the action
	g.action = action
	return
}
