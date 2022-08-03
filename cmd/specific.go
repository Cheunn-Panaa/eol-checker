package cmd

import (
	"github.com/cheunn-panaa/eol-checker/pkg/commands"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	productList string
	versionList string
)

// allCmd represents the all command
var specificCmd = &cobra.Command{
	Use:   "specific",
	Short: "Check end of life for one specific language in your configfile",
	Run: func(cmd *cobra.Command, args []string) {
		commands.RunSpecific()
	},
}

func init() {

	specificCmd.PersistentFlags().StringVar(&productList, "products", "", "Select one or multiple specific products to check from the config file or standalone")
	specificCmd.PersistentFlags().StringVar(&versionList, "versions", "", "If specific product is defined then you can specify a specific version aswell")

	viper.BindPFlag("productList", specificCmd.PersistentFlags().Lookup("products"))
	viper.BindPFlag("versionList", specificCmd.PersistentFlags().Lookup("versions"))

	viper.Set("disable-message", &disabledMessage)
	rootCmd.AddCommand(specificCmd)
}
