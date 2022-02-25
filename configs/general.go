package configs

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

//Config base type for everything
type Config struct {
	Default  Default   `yaml:"default"`
	Products []Product `yaml:"products"`
	Plugins  Plugins   `yaml:"plugins"`
}

//Plugins is the type that stores every plugins conf
type Plugins struct {
	Slack Slack `yaml:"slack"`
}

//Default is the base conf for the app
type Default struct {
	Url      string   `yaml:"url"`
	Alerting Alerting `yaml:"alert"`
}

//Alerting is the cadence you concider critical for your apps
type Alerting struct {
	Month int `yaml:"month"`
}

//Product is the default config for product/software eol
type Product struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Custom  Custom `yaml:"custom_url"`
}

//TODO: do it ?
//Custom custom config for eol
type Custom struct {
	Url    string `yaml:"url"`
	Method string `yaml:"method"`
	Body   string `yaml:"body"`
}

// NewConfig is constructor
func NewConfig(filename string) (config *Config, err error) {
	b, err := ioutil.ReadFile(filename)
	err = yaml.Unmarshal(b, &config)
	return
}
