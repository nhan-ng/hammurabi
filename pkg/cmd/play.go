package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	h "github.com/nhan-ng/hammurabi/pkg/hammurabi"
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
	game := h.NewGame(maxYears)
	reader := bufio.NewReader(os.Stdin)
	for year := 1; year <= maxYears; {
		game.DisplayGameState(year)

		// Ask for input
		_, err := game.ReadActionInput(reader)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Once again?")
			continue
		}

		// Transition to the new state
		nextYear, _, _, err := game.Transition()
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
	}
}
