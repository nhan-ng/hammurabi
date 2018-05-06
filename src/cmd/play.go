package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	h "../hammurabi"
)

const (
	shortPlayCmdDesc = "Plays Hammurabi interactively"
	longPlayCmdDesc  = "Plays Hammurabi interactively turn by turn."
)

var (
	maxYears int
)

func newPlayCmd() *cobra.Command {
	playCmd := &cobra.Command{
		Use:   "play",
		Short: shortPlayCmdDesc,
		Long:  longPlayCmdDesc,
		Run:   runPlayCmd,
	}

	// Add flags
	// playCmd.Flags().IntVarP(&maxYears, "years", "y", 10, "Max number of years to play")
	playCmd.Flags().IntVarP(&maxYears, "years", "y", 10, "Max number of years to play")
	return playCmd
}

func runPlayCmd(cmd *cobra.Command, args []string) {
	state, delta := h.NewGame(maxYears)
	reader := bufio.NewReader(os.Stdin)
	for year := 0; year < maxYears; {
		h.DisplayGameState(year, state, delta)

		// Ask for input
		action, err := h.ReadActionInput(reader)
		if err != nil {
			log.Fatalln(err)
		}

		nextYear, nextState, nextDelta, err := h.Transition(year, state, action)
		if err != nil {
			fmt.Println(err)
			fmt.Println("There was an error. Please try again.")
			continue
		}

		year = nextYear
		state = nextState
		delta = nextDelta
	}
}
