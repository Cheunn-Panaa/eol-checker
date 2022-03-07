package cmd

import (
	"github.com/cheunn-panaa/eol-checker/pkg/commands"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "check end of life for every project present in configfile",
	Run: func(cmd *cobra.Command, args []string) {
		commands.RunAll()
	},
}

func init() {
	RootCmd.AddCommand(allCmd)
}
