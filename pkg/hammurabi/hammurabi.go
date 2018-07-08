package hammurabi

import (
	"fmt"
	"math/rand"

	"github.com/pkg/errors"
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
	LandsToBuy    int
	BushelsToFeed int
	LandsToSeed   int
}

// StateDelta captures the changes made to the game state from the transition of the previous year to the next.
type StateDelta struct {
	PeopleStarved   int
	PeopleKilled    int
	PeopleAdded     int
	BushelsInfested int
	HasRat          bool
	HasPlague       bool
}

// Game represents a Hammurabi game
type Game struct {
	State  *GameState
	Delta  *StateDelta
	Action *GameAction
	Year   int
}

// Hammurabi represents the minimal interface a Hammurabi game must have.
type Hammurabi interface {
	Transition() (int, *GameState, *StateDelta, error)
}

const (
	bushelsPerPerson = 20
	landsPerPerson   = 10
	bushelsPerLand   = 1

	minRatPercentage = 0.1
	maxRatPercentage = 0.3

	uprisingThreshold = 0.45

	plagueChance = 0.15
	ratChance    = 0.40

	minLandPrice = 17
	maxLandPrice = 26

	minLandProfit = 1
	maxLandProfit = 6

	minNewcomers = 2
	maxNewcomers = 5

	initialPopulation = 100
	initialBushels    = 2800
	initialLands      = 1000
	initialLandPrice  = 22
	initialLandProfit = 3

	initialNewPeople       = 5
	initialBushelsInfested = 200
)

// NewGame creates a new game with the maximum number of years, aka turns.
func NewGame(maxYear int) *Game {
	// Create a fixed initial state delta
	delta := &StateDelta{
		PeopleStarved:   0,
		PeopleKilled:    0,
		PeopleAdded:     initialNewPeople,
		BushelsInfested: initialBushelsInfested,
		HasRat:          true,
		HasPlague:       false,
	}

	// Create a fixed initial state
	state := &GameState{
		Bushels:    initialBushels,
		Population: initialPopulation,
		Lands:      initialLands,
		LandPrice:  initialLandPrice,
		LandProfit: initialLandProfit,
	}

	// Initialize a new game
	return &Game{
		State: state,
		Delta: delta,
		Year:  1,
	}
}

// Transition transitions the given game state to the next.
func (g *Game) Transition() (nextYear int, nextState *GameState, delta *StateDelta, err error) {
	// Get the state and action
	state := g.State
	action := g.Action

	// Validate parameters
	if state == nil {
		err = &nilGameState{}
		return
	}
	if action == nil {
		err = &nilGameAction{}
		return
	}
	if err = errors.Wrap(validate(action), "validation failed"); err != nil {
		return
	}

	// Initialize the next game state and delta
	nextState = &GameState{}
	delta = &StateDelta{}

	// Make a copy of the current state
	*nextState = *state

	// Perform land tranding action
	if err = errors.Wrap(validateLandTradingAction(nextState, action), "land trading validation failed"); err != nil {
		return
	}
	nextState.Bushels = nextState.Bushels - state.LandPrice*action.LandsToBuy
	nextState.Lands = nextState.Lands + action.LandsToBuy

	// Perform the action to feed the population
	if err = errors.Wrap(validateFeedingAction(nextState, action), "feeding validation failed"); err != nil {
		return
	}
	numPeopleFed := min(nextState.Population*bushelsPerPerson, action.BushelsToFeed) / bushelsPerPerson
	nextState.Bushels = nextState.Bushels - numPeopleFed*bushelsPerPerson
	nextState.Population = numPeopleFed
	delta.PeopleStarved = state.Population - nextState.Population
	if u, e := isUprising(state.Population, delta.PeopleStarved); u || e != nil {
		if u {
			err = &Uprising{Year: g.Year, PeopleStarved: delta.PeopleStarved, Percentage: float32(delta.PeopleStarved) / float32(state.Population) * 100.0}
		} else {
			err = errors.Wrap(e, "uprising validation failed")
		}
		return
	}

	// Perform the action to plant seed
	if err = errors.Wrap(validateSeedingAction(nextState, action), "seeding validation failed"); err != nil {
		return
	}

	// Calculate the harvestable lands and profits
	maxLandsToHarvest := min(nextState.Population/landsPerPerson, nextState.Lands)
	landsHarvested := min(maxLandsToHarvest, action.LandsToSeed)
	nextState.Bushels = nextState.Bushels + landsHarvested*(state.LandProfit-bushelsPerLand)

	// Now add randomized events to the mix
	// Plague
	delta.PeopleKilled, err = getPeopleKilledByPlague(nextState.Population)
	if err != nil {
		err = errors.Wrap(err, "validation failed")
		return
	}
	delta.HasPlague = delta.PeopleKilled != 0
	nextState.Population = nextState.Population - delta.PeopleKilled

	// Rat
	delta.BushelsInfested, err = getInfestedBushels(nextState.Bushels)
	if err != nil {
		err = errors.Wrap(err, "validation failed")
		return
	}
	delta.HasRat = delta.BushelsInfested != 0
	nextState.Bushels = nextState.Bushels - delta.BushelsInfested

	// Newcomers
	newcomers, err := getNewcomers()
	if err != nil {
		err = errors.Wrap(err, "validation failed")
		return
	}
	nextState.Population = nextState.Population + newcomers
	delta.PeopleAdded = newcomers

	// Generate next year land price and profit
	nextState.LandPrice, err = getNextLandPrice()
	if err != nil {
		err = errors.Wrap(err, "validation failed")
		return
	}
	nextState.LandProfit, err = getNextLandProfit()
	if err != nil {
		err = errors.Wrap(err, "validation failed")
		return
	}

	// Increment year
	nextYear = g.Year + 1

	// Update the game year
	g.Year = nextYear
	g.State = nextState
	g.Delta = delta

	// Done
	return
}

func validate(action *GameAction) error {
	if action.BushelsToFeed < 0 {
		return &valueOutOfRange{kind: "BushelsToFeed", reason: "Must be non-negative"}
	}
	if action.LandsToSeed < 0 {
		return &valueOutOfRange{kind: "LandsToSeed", reason: "Must be non-negative"}
	}
	return nil
}

func validateLandTradingAction(state *GameState, action *GameAction) error {
	// Validate if we have enough to sell
	if action.LandsToBuy < 0 && state.Lands+action.LandsToBuy < 0 {
		return &insufficientLandsToSell{currentLands: state.Lands, requestedLands: -action.LandsToBuy}
	}

	// Validate if we have enough bushels to buy
	if action.LandsToBuy > 0 {
		nextBushels := state.Bushels - state.LandPrice*action.LandsToBuy
		if nextBushels < 0 {
			return &insufficientBushelsToBuyLands{currentBushels: state.Bushels, requiredBushels: state.Bushels - nextBushels}
		}
	}

	// All good
	return nil
}

func validateFeedingAction(state *GameState, action *GameAction) error {
	// Validate if we have enough bushels to feed
	maxBushelsRequired := state.Population * bushelsPerPerson
	bushelsToFeed := min(maxBushelsRequired, action.BushelsToFeed)
	if bushelsToFeed > state.Bushels {
		return &insufficientBushelsToFeed{currentBushels: state.Bushels, requestedBushels: bushelsToFeed}
	}
	return nil
}

func isUprising(population, starved int) (ret bool, err error) {
	// Validate
	if population < 0 {
		err = &valueOutOfRange{kind: "population", reason: "Must be non-negative"}
	} else if starved < 0 {
		err = &valueOutOfRange{kind: "starved", reason: "Must be non-negative"}
	} else if starved > population {
		err = &valueOutOfRange{kind: "starved", reason: fmt.Sprintf("Must be smaller than population %d", population)}
	}
	if err != nil {
		return
	}

	// Return true if the starved population is more than uprising threshold
	ret = (float32(starved) / float32(population)) > uprisingThreshold
	return
}

func validateSeedingAction(state *GameState, action *GameAction) error {
	// Validate that we have enough bushels to seed
	maxBushelsRequired := min(state.Population*landsPerPerson, state.Lands) * bushelsPerLand
	requestedBushels := min(maxBushelsRequired, action.LandsToSeed*bushelsPerLand)
	if state.Bushels < requestedBushels {
		return &insufficientBushelsToSeed{currentBushels: state.Bushels, requestedBushels: requestedBushels}
	}
	return nil
}

func getPeopleKilledByPlague(population int) (ret int, err error) {
	// Validate
	if population < 0 {
		err = &valueOutOfRange{kind: "population", reason: "Must be non-negative"}
		return
	}

	// Early exit if there is no plague
	if rand.Float32() > plagueChance {
		ret = 0
	} else {
		// Otherwise, kill half the population
		ret = population / 2
	}

	return
}

func getInfestedBushels(bushels int) (ret int, err error) {
	// Validate
	if bushels < 0 {
		err = &valueOutOfRange{kind: "bushels", reason: "Must be non-negative"}
		return
	}

	// Early exit there is no rat infestation
	if rand.Float32() > ratChance {
		ret = 0
	} else {
		// Otherwise infest randomly between 10% to 40% of the bushels
		ret = int(float32(bushels) * (rand.Float32()*0.3 + 0.1))
	}
	return
}

func getNewcomers() (int, error) {
	return randIntInRangeInclusive(minNewcomers, maxNewcomers)
}

func getNextLandPrice() (int, error) {
	return randIntInRangeInclusive(minLandPrice, maxLandPrice)
}

func getNextLandProfit() (int, error) {
	return randIntInRangeInclusive(minLandProfit, maxLandProfit)
}

func randIntInRangeInclusive(lowInclusive, highInclusive int) (ret int, err error) {
	// Validate and return corner case
	if lowInclusive > highInclusive {
		err = &valueOutOfRange{kind: "lowInclusive", reason: fmt.Sprintf("Must be smaller than highInclusive %d", highInclusive)}
		return
	} else if lowInclusive == highInclusive {
		ret = lowInclusive
		return
	}

	// Return a random number between low and high inclusively
	ret = lowInclusive + rand.Intn(highInclusive-lowInclusive+1)
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
