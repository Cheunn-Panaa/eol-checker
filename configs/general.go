package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config *Configuration

//Config base type for everything
type Configuration struct {
	Default  Default   `mapstructure:"default"`
	Products []Product `mapstructure:"products"`
	Plugins  Plugins   `mapstructure:"plugins"`
}

//Plugins is the type that stores every plugins conf
type Plugins struct {
	Slack Slack `mapstructure:"slack"`
}

//Default is the base conf for the app
type Default struct {
	Url      string   `mapstructure:"url"`
	Alerting Alerting `mapstructure:"alert"`
}

//Alerting is the cadence you consider critical for your apps
type Alerting struct {
	Month int `mapstructure:"month"`
}

//Product is the default config for product/software eol
type Product struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Custom  Custom `mapstructure:"custom_url"`
}

//TODO: do it ?
//Custom custom config for eol
type Custom struct {
	Url    string `mapstructure:"url"`
	Method string `mapstructure:"method"`
	Body   string `mapstructure:"body"`
}

// LoadConfiguration is constructor
func LoadConfiguration(configFile string, envPrefix string, v *viper.Viper) {
	var configuration *Configuration

	v.SetConfigFile(configFile)
	// We are only looking in the current working directory.
	v.AddConfigPath(".")

	v.SetEnvPrefix(envPrefix)

	// read in environment variables that match
	v.AutomaticEnv()

	// If a config file is found, read it in. else returns an error
	if err := v.ReadInConfig(); err == nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Sprintf("Error reading config file, %s", err)
		}
	}

	err := v.Unmarshal(&configuration)
	if err != nil {
		fmt.Sprintf("Unable to decode into struct, %v", err)
	}

	Config = configuration
}

func GetConfig() *Configuration {
	return Config
}
