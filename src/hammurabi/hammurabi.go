package hammurabi

import (
	"fmt"
)

// GameState describes the current resource state of the game.
type GameState struct {
	Bushels    int
	Population int
	Lands      int
	LandPrice  int
	LandProfit int
}

// GameAction describes the player's action given a game state.
type GameAction struct {
	LandToBuy    int
	BushelToFeed int
	LandToSeed   int
}

// StateDelta captures the changes made to the game state from the transition of the previous year to the next.
type StateDelta struct {
	PeopleStarved int
	PeopleKilled  int
	PeopleAdded   int
	BushelLost    int
	HasRat        bool
	HasPlague     bool
}

const (
	bushelsPerPerson  = 20
	maxLandsPerPerson = 10

	minRatPercentage = 0
	maxRatPercentage = 0.4

	plagueChance = 0.15

	minLandPrice = 17
	maxLandPrice = 26

	minNewPeople = 2
	maxNewPeople = 5

	initialPopulation = 100
	initialBushels    = 2800
	initialLands      = 1000
	initialLandPrice  = 22
	initialLandProfit = 3

	initialNewPeople  = 5
	initialBushelLost = 200
)

// NewGame creates a new game with the maximum number of years, aka turns.
func NewGame(maxYear int) (*GameState, *StateDelta) {
	// Create a fixed initial state delta
	delta := &StateDelta{
		PeopleStarved: 0,
		PeopleKilled:  0,
		PeopleAdded:   initialNewPeople,
		BushelLost:    initialBushelLost,
		HasRat:        true,
		HasPlague:     false,
	}

	// Create a fixed initial state
	state := &GameState{
		Bushels:    initialBushels,
		Population: initialPopulation,
		Lands:      initialLands,
		LandPrice:  initialLandPrice,
		LandProfit: initialLandProfit,
	}

	return state, delta
}

// DisplayGameState displays textual representation of the game state and state delta.
func DisplayGameState(year int, state *GameState, delta *StateDelta) {
	// Display general information
	fmt.Println("Hammurabi: I beg to report to you,")
	fmt.Printf("In Year %d, %d people starved.\n", year, delta.PeopleStarved)
	fmt.Printf("%d people came to the city.\n", delta.PeopleAdded)
	fmt.Printf("The city population is now %d.\n", state.Population)
	fmt.Printf("The city now owns %d acres.\n", state.Lands)
	fmt.Printf("You harvested %d bushels per acre.\n", state.LandProfit)
	if delta.HasRat {
		fmt.Printf("Rats ate %d bushels.\n", delta.BushelLost)
	}
	if delta.HasPlague {
		fmt.Printf("Plague killed %d people.\n", delta.PeopleKilled)
	}
	fmt.Printf("You now have %d bushels in store.\n", state.Bushels)
	fmt.Printf("Land is trading at %d bushels per acre.\n", state.LandPrice)
}

// Transition transitions the given game state to the next.
func Transition(currState *GameState, action *GameAction) (*GameState, *StateDelta) {
	// Do magic here
	return nil, nil
}
