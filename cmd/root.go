package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/cheunn-panaa/eol-checker/configs"
	"github.com/cheunn-panaa/eol-checker/internal/utils"
	"github.com/cheunn-panaa/eol-checker/pkg/commands"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	configFile      string
	disabledMessage bool
)

const (
	// The name of our config file
	defaultConfigFilename = "eol-cli"

	// The environment variable prefix of all environment variables bound to our command line flags.
	// For example, --number is bound to EOL_NUMBER.
	envPrefix = "EOL"

	defaultConfigExtension = "yaml"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "eol-cli",
	Short: "Check if your application runtime is out of date",
	Long: "EOL is a CLI library for project version management. \n This application is a tool to check wether or not your application version is out of date. \n" +
		"For now it checks on endoflife.date endpoints but is subject to change in the future",
	SilenceUsage: true,
	Version:      utils.GetVersion(),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//TODO: verbose impl
		// Binding cobra and viper together
		//	if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
		//		debug.Verbose = verbose
		//	}
		//initializeConfig()
		cmd.SetVersionTemplate(utils.GetVersion())
		//TODO: find a way to make bindFlags(cmd) work
	},
	Run: func(cmd *cobra.Command, args []string) {
		commands.Run()
	},
}

// Execute adds all child commands to the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)

		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initializeConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", fmt.Sprintf("config file (default is %s)", defaultConfigFilename))
	rootCmd.PersistentFlags().BoolVarP(&disabledMessage, "disable-message", "d", false, "Disables the notifications on all plugins")
	viper.Set("disable-message", &disabledMessage)
}

// initializeConfig reads in config file and ENV variables if set.
func initializeConfig() {
	var configuration *configs.Configuration

	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Find home directory.
		dir, err := utils.GetApplicationDir()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".twitch-cli" (without extension).
		viper.AddConfigPath(dir)
		viper.SetConfigName(defaultConfigFilename)
		viper.SetConfigType(defaultConfigExtension)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// We are only looking in the current working directory.
	viper.AddConfigPath(".")

	viper.SetEnvPrefix(envPrefix)

	// read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in. else returns an error
	if err := viper.ReadInConfig(); err == nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Sprintf("Error reading config file, %s", err)
		}
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Sprintf("Unable to decode into struct, %v", err)
	}

	configs.SetConfiguration(configuration)

}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Environment variables can't have dashes in them, so bind them to their equivalent
		// keys with underscores, e.g. --disable-messaging to EOL_DISABLE_MESSAGING
		fmt.Printf("%s, %s \n", f.Name, f.Value)
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			viper.BindEnv(f.Name, fmt.Sprintf("%s_%s", envPrefix, envVarSuffix))
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && viper.IsSet(f.Name) {
			val := viper.Get(f.Name)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
