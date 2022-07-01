package cmd

import (
	"fmt"

	"github.com/cheunn-panaa/eol-checker/internal/utils"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Returns the current version of the CLI.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("eol-cli/" + utils.GetVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
