package cmd

import (
	"github.com/spf13/cobra"

	h "../hammurabi"
)

const (
	shortPlayCmdDesc = "Plays Hammurabi interactively"
	longPlayCmdDesc  = "Plays Hammurabi interactively turn by turn."
)

func newPlayCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "play",
		Short: shortPlayCmdDesc,
		Long:  longPlayCmdDesc,
		Run:   runPlayCmd,
	}
}

func runPlayCmd(cmd *cobra.Command, args []string) {
	maxYear := 1
	s, d := h.NewGame(maxYear)
	for i := 0; i < maxYear; i++ {
		h.DisplayGameState(i, s, d)

		// Ask for input
	}
}
