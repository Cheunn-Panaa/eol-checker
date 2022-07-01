package cmd

import (
	"github.com/cheunn-panaa/eol-checker/pkg/commands"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var disabledMessage bool

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "check end of life for every project present in a configfile",
	Run: func(cmd *cobra.Command, args []string) {
		commands.RunAll()
	},
}

func init() {

	allCmd.Flags().BoolVarP(&disabledMessage, "disable-message", "d", false, "Disables the notifications on all plugins")
	viper.Set("disable-message", &disabledMessage)
	rootCmd.AddCommand(allCmd)
}
