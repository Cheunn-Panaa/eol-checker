package cmd

import (
	"github.com/cheunn-panaa/eol-checker/pkg/commands"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	productList string
)

// allCmd represents the all command
var specificCmd = &cobra.Command{
	Use:   "specific",
	Short: "Check end of life for one specific language in your configfile",
	Run: func(cmd *cobra.Command, args []string) {
		commands.Run()
	},
}

func init() {

	specificCmd.Flags().BoolVarP(&disabledMessage, "disable-message", "d", false, "Disables the notifications on all plugins")
	rootCmd.PersistentFlags().StringVar(&productList, "products", "", "Select one or multiple specific language to check from the config file")
	viper.Set("disable-message", &disabledMessage)
	rootCmd.AddCommand(specificCmd)
}
