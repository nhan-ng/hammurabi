package hammurabi

import (
	"fmt"
)

// Uprising represents an error when more than 45% of the population is killed and the people overthrowed the player
type Uprising struct {
	Year          int
	PeopleStarved int
	Percentage    float32
}

// NilGameState represents an error when the game state is nil.
type NilGameState struct{}

// NilGameAction represents an error when the game action is nil.
type NilGameAction struct{}

// ValueOutOfRange represents an error when the value of given type is out of range.
type ValueOutOfRange struct {
	Type   string
	Reason string
}

// InsufficientBushelsToBuyLands represents an error when the player's action is requesting to buy land with bushels they don't have.
type InsufficientBushelsToBuyLands struct {
	CurrentBushels  int
	RequiredBushels int
}

// InsufficientLandsToSell represents an error when the player's action is requesting to sell lands that they don't have.
type InsufficientLandsToSell struct {
	CurrentLands   int
	RequestedLands int
}

// InsufficientBushelsToFeed represents an error when the player's action is requesting to feed the people with bushels they don't have.
type InsufficientBushelsToFeed struct {
	CurrentBushels   int
	RequestedBushels int
}

// InsufficientBushelsToSeed represents an error when the player's action is requesting to seed using bushels that they don't have.
type InsufficientBushelsToSeed struct {
	CurrentBushels   int
	RequestedBushels int
}

func (e *Uprising) Error() string {
	return fmt.Sprintf("In year %d, you starved %d people (%.2f%% of the population). The people overthrowed you.", e.Year, e.PeopleStarved, e.Percentage)
}

func (e *InsufficientBushelsToBuyLands) Error() string {
	return fmt.Sprintf("Insufficient bushels to buy lands. Having %d bushels but required %d bushels to buy.", e.CurrentBushels, e.RequiredBushels)
}

func (e *InsufficientLandsToSell) Error() string {
	return fmt.Sprintf("Insufficient lands to sell. Having %d acres of land but requested to sell %d acres.", e.CurrentLands, e.RequestedLands)
}

func (e *InsufficientBushelsToFeed) Error() string {
	return fmt.Sprintf("Insufficient bushels to feed people. Having %d bushels but requested for %d bushels to feed.", e.CurrentBushels, e.RequestedBushels)
}

func (e *InsufficientBushelsToSeed) Error() string {
	return fmt.Sprintf("Insufficient bushels to seed. Having %d bushels but requested %d bushels to seed.", e.CurrentBushels, e.RequestedBushels)
}

func (e *NilGameState) Error() string {
	return fmt.Sprintf("Game state is nil.")
}

func (e *NilGameAction) Error() string {
	return fmt.Sprintf("Game action is nil.")
}

func (e *ValueOutOfRange) Error() string {
	return fmt.Sprintf("Value of type '%s' is out of range. Reason %s.", e.Type, e.Reason)
}
