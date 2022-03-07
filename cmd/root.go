package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/cheunn-panaa/eol-checker/configs"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var configFile string

const (
	// The name of our config file
	defaultConfigFilename = "eol-config.yaml"

	// The environment variable prefix of all environment variables bound to our command line flags.
	// For example, --number is bound to EOL_NUMBER.
	envPrefix = "EOL"
)

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:          "dol-eol",
	Short:        "Check if your application runtime is out of date",
	SilenceUsage: true,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//TODO: verbose impl
		// Binding cobra and viper together
		//	if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
		//		debug.Verbose = verbose
		//	}
		initializeConfig(cmd)
	},
}

// Execute adds all child commands to the root command.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&configFile, "config", "", fmt.Sprintf("config file (default is %s)", defaultConfigFilename))
}

// initializeConfig reads in config file and ENV variables if set.
func initializeConfig(cmd *cobra.Command) {
	v := viper.New()
	if configFile == "" {
		configFile = defaultConfigFilename
	}

	configs.LoadConfiguration(configFile, envPrefix, v)

	// Bind the current command's flags to viper
	bindFlags(cmd, v)
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Environment variables can't have dashes in them, so bind them to their equivalent
		// keys with underscores, e.g. --disable-messaging to EOL_DISABLE_MESSAGING
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			v.BindEnv(f.Name, fmt.Sprintf("%s_%s", envPrefix, envVarSuffix))
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
