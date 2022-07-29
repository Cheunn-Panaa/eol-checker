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

const (
	// The name of our config file
	defaultConfigFilename = "eol-cli"

	// The environment variable prefix of all environment variables bound to our command line flags.
	// For example, --number is bound to EOL_NUMBER.
	envPrefix = "EOL"

	defaultConfigExtension = "yaml"

	DEV_VERSION = "0.0.3-dev"

	author = "Cheunn NOURRY <cnourry.dev@protonmail.com>"
)

var (
	configFile      string
	disabledMessage bool
	verbose         bool // whether to display request info
	// to be populated by linker/builds
	version string

	showVers bool // whether to print version info or not

	// rootCmd represents the base command when called without any subcommands.
	rootCmd = &cobra.Command{
		Use:   "eol-cli",
		Short: "Check if your application runtime is out of date",
		Long: "EOL is a CLI library for project version management. \n This application is a tool to check wether or not your application version is out of date. \n" +
			"For now it checks on endoflife.date endpoints but is subject to change in the future",
		SilenceUsage: true,
		Version:      version,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {

			//TODO: verbose impl -> preflight ?
			//TODO: find a way to make bindFlags(cmd) work
			// Binding cobra and viper together
			//	if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			//		debug.Verbose = verbose
			//	}
		},
		// print version or help, or continue, depending on flag settings
		PreRunE: preFlight,
		Run: func(cmd *cobra.Command, args []string) {
			commands.Run()
		},
	}
)

// Execute adds all child commands to the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initializeConfig)
	configFileName := fmt.Sprintf("%s.%s", defaultConfigFilename, defaultConfigExtension)

	// Cli specific
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", configFileName, "Configuration file to use")
	rootCmd.Flags().BoolVarP(&showVers, "version", "v", false, "Display the current version of this CLI")

	rootCmd.PersistentFlags().BoolVarP(&disabledMessage, "disable-message", "d", false, "Disables the notifications on all plugins")
	viper.Set("disable-message", &disabledMessage)

	// Generic CLI conf
	viper.SetDefault("author", author)
	viper.SetDefault("license", "AGPL-3.0")
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

		// Search config in home directory with name "eol-cli" (without extension).
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

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error: Failed to read config file -", configFile)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		fmt.Println("Unable to decode into struct - ", err)
		os.Exit(1)
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

func preFlight(cmd *cobra.Command, args []string) error {
	// if --version is passed print the version info
	if showVers {
		fmt.Printf("eol-cli %s \n", version)
		return fmt.Errorf("")
	}

	return nil
}
