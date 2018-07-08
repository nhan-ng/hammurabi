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

// nilGameState represents an error when the game state is nil.
type nilGameState struct{}

// nilGameAction represents an error when the game action is nil.
type nilGameAction struct{}

// valueOutOfRange represents an error when the value of given type is out of range.
type valueOutOfRange struct {
	kind   string
	reason string
}

// insufficientBushelsToBuyLands represents an error when the player's action is requesting to buy land with bushels they don't have.
type insufficientBushelsToBuyLands struct {
	currentBushels  int
	requiredBushels int
}

// insufficientLandsToSell represents an error when the player's action is requesting to sell lands that they don't have.
type insufficientLandsToSell struct {
	currentLands   int
	requestedLands int
}

// insufficientBushelsToFeed represents an error when the player's action is requesting to feed the people with bushels they don't have.
type insufficientBushelsToFeed struct {
	currentBushels   int
	requestedBushels int
}

// insufficientBushelsToSeed represents an error when the player's action is requesting to seed using bushels that they don't have.
type insufficientBushelsToSeed struct {
	currentBushels   int
	requestedBushels int
}

// invalidInput represents an error when the input is invalid.
type invalidInput struct {
}

func (e *Uprising) Error() string {
	return fmt.Sprintf("In year %d, you starved %d people (%.2f%% of the population). The people overthrowed you.", e.Year, e.PeopleStarved, e.Percentage)
}

func (e *insufficientBushelsToBuyLands) Error() string {
	return fmt.Sprintf("Insufficient bushels to buy lands. Having %d bushels but required %d bushels to buy.", e.currentBushels, e.requiredBushels)
}

func (e *insufficientLandsToSell) Error() string {
	return fmt.Sprintf("Insufficient lands to sell. Having %d acres of land but requested to sell %d acres.", e.currentLands, e.requestedLands)
}

func (e *insufficientBushelsToFeed) Error() string {
	return fmt.Sprintf("Insufficient bushels to feed people. Having %d bushels but requested for %d bushels to feed.", e.currentBushels, e.requestedBushels)
}

func (e *insufficientBushelsToSeed) Error() string {
	return fmt.Sprintf("Insufficient bushels to seed. Having %d bushels but requested %d bushels to seed.", e.currentBushels, e.requestedBushels)
}

func (e *nilGameState) Error() string {
	return fmt.Sprintf("Game state is nil.")
}

func (e *nilGameAction) Error() string {
	return fmt.Sprintf("Game action is nil.")
}

func (e *valueOutOfRange) Error() string {
	return fmt.Sprintf("Value of type '%s' is out of range. Reason %s.", e.kind, e.reason)
}

func (e *invalidInput) Error() string {
	return "Invalid input"
}
