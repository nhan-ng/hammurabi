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
		fmt.Println()
		if err != nil {
			if e, ok := err.(*h.Uprising); ok {
				fmt.Println(e)
				fmt.Printf("Due to this extreme mismanagement, you have not only been impeached and thrown out of office, but you have also been declared 'National Fink'!")
				return
			}

			fmt.Println("O Great Hammurabi, surely you jest! We seem to have problem understanding your decisions!")
			fmt.Println(err)
			continue
		}

		year = nextYear
		state = nextState
		delta = nextDelta
	}
}
