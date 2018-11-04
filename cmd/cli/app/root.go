package app

import (
	"github.com/spf13/cobra"
)

const (
	shortRootCmdDesc = "Hammurabi"
	longRootCmdDesc  = `Long`
)

// NewHammurabiCmd creates a new root cmd for Hammurabi
func NewHammurabiCmd() *cobra.Command {
	// Create the root command
	cmd := &cobra.Command{
		Use:   "hammurabi",
		Short: shortRootCmdDesc,
		Long:  longRootCmdDesc,
		Run: func(cmd *cobra.Command, args []string) {
			// Do main stuff
		},
	}

	// Add other commands
	cmd.AddCommand(newPlayCmd())

	return cmd
}
