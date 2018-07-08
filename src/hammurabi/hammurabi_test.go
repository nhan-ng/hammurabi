package hammurabi

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestValidateAction(t *testing.T) {
	testCases := []struct {
		name        string
		action      *GameAction
		expectError bool
	}{
		{
			name:        "Negative bushels to feed",
			action:      &GameAction{BushelsToFeed: -1},
			expectError: true,
		},
		{
			name:        "Negative lands to seed",
			action:      &GameAction{LandsToSeed: -1},
			expectError: true,
		},
		{
			name: "Positive bushels to feed and lands to seed",
			action: &GameAction{
				BushelsToFeed: 1,
				LandsToSeed:   1,
			},
			expectError: false,
		},
		{
			name: "Zero bushels to feed and lands to seed",
			action: &GameAction{
				BushelsToFeed: 0,
				LandsToSeed:   0,
			},
			expectError: false,
		},
	}

	// Act
	for _, tc := range testCases {
		actualError := errors.Cause(validate(tc.action))
		if tc.expectError && actualError == nil {
			t.Errorf("[%s] expects error, but got nil", tc.name)
		}
		if !tc.expectError && actualError != nil {
			t.Errorf("[%s] expects no error, but got an error", tc.name)
		}
	}
}

func TestValidateLandTradingAction(t *testing.T) {
	testCases := []struct {
		name          string
		state         *GameState
		action        *GameAction
		expectedError error
	}{
		{
			name:   "Insufficient lands to sell",
			state:  &GameState{Lands: 10},
			action: &GameAction{LandsToBuy: -20},
			expectedError: &insufficientLandsToSell{
				currentLands:   10,
				requestedLands: 20,
			},
		},
		{
			name:   "Insufficient bushels to buy",
			state:  &GameState{Bushels: 100, LandPrice: 20},
			action: &GameAction{LandsToBuy: 100},
			expectedError: &insufficientBushelsToBuyLands{
				currentBushels:  100,
				requiredBushels: 2000,
			},
		},
		{
			name:          "Sufficient lands to sell",
			state:         &GameState{Lands: 20},
			action:        &GameAction{LandsToBuy: -10},
			expectedError: nil,
		},
		{
			name:          "Sufficient bushels to buy",
			state:         &GameState{Bushels: 2000, LandPrice: 10},
			action:        &GameAction{LandsToBuy: 10},
			expectedError: nil,
		},
	}

	// Act
	for _, tc := range testCases {
		actualError := errors.Cause(validateLandTradingAction(tc.state, tc.action))

		// Assert the error
		if !reflect.DeepEqual(actualError, tc.expectedError) {
			t.Errorf("[%s] expected:\n%#v\ngot:\n%#v", tc.name, tc.expectedError, actualError)
		}
	}
}

func TestValidateFeedingAction(t *testing.T) {
	testCases := []struct {
		name          string
		state         *GameState
		action        *GameAction
		expectedError error
	}{
		{
			name: "Action asks for more bushels than affordable",
			state: &GameState{
				Population: 10,
				Bushels:    bushelsPerPerson * 5, // Only enough to feed 5
			},
			action: &GameAction{
				BushelsToFeed: bushelsPerPerson * 10, // Asking to feed 10
			},
			expectedError: &insufficientBushelsToFeed{
				currentBushels:   bushelsPerPerson * 5,
				requestedBushels: bushelsPerPerson * 10,
			},
		},
		{
			name: "Action asks for less bushels than required",
			state: &GameState{
				Population: 10,
				Bushels:    bushelsPerPerson * 20, // Enough bushels to feed 20
			},
			action: &GameAction{
				BushelsToFeed: bushelsPerPerson * 9,
			},
			expectedError: nil,
		},
		{
			name: "Action asks for more bushels than maximum required",
			state: &GameState{
				Population: 10,
				Bushels:    bushelsPerPerson * 20, // Enough to feed 20
			},
			action: &GameAction{
				BushelsToFeed: bushelsPerPerson * 30, // Asking to feed 30
			},
			expectedError: nil,
		},
	}

	// Act
	for _, tc := range testCases {
		actualError := errors.Cause(validateFeedingAction(tc.state, tc.action))

		// Assert the error
		if !reflect.DeepEqual(actualError, tc.expectedError) {
			t.Errorf("[%s] expected:\n%#v\ngot:\n%#v", tc.name, tc.expectedError, actualError)
		}
	}
}

func TestValidateSeedingAction(t *testing.T) {
	testCases := []struct {
		name          string
		state         *GameState
		action        *GameAction
		expectedError error
	}{
		{
			name: "Action asks for more bushels to seed than affordable",
			state: &GameState{
				Population: 100 * landsPerPerson, // Enough people to farm 100 lands
				Lands:      10,
				Bushels:    bushelsPerLand * 5, // Only afforable to seed 5 lands
			},
			action: &GameAction{
				LandsToSeed: 10,
			},
			expectedError: &insufficientBushelsToSeed{
				currentBushels:   bushelsPerLand * 5,
				requestedBushels: bushelsPerLand * 10,
			},
		},
		{
			name: "Action asks for enough bushels",
			state: &GameState{
				Population: 100 * landsPerPerson, // Enough people to farm 100 lands
				Lands:      10,
				Bushels:    bushelsPerLand * 10, // Affordable to seed 10 lands
			},
			action: &GameAction{
				LandsToSeed: 5,
			},
			expectedError: nil,
		},
		{
			name: "Action asks for more bushels than maximum required",
			state: &GameState{
				Population: 100 * landsPerPerson, // Enough people to farm 100 lands
				Lands:      10,
				Bushels:    bushelsPerLand * 20, // Affordable to seed 20 lands
			},
			action: &GameAction{
				LandsToSeed: 50,
			},
			expectedError: nil,
		},
	}

	// Act
	for _, tc := range testCases {
		actualError := errors.Cause(validateSeedingAction(tc.state, tc.action))

		// Assert the error
		if !reflect.DeepEqual(actualError, tc.expectedError) {
			t.Errorf("[%s] expected:\n%#v\ngot:\n%#v", tc.name, tc.expectedError, actualError)
		}
	}
}
