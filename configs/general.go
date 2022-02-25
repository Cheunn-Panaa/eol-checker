package configs

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Default  Default   `yaml:"default"`
	Products []Product `yaml:"products"`
	Plugins  Plugins   `yaml:"plugins"`
}

type Plugins struct {
	Slack Slack `yaml:"slack"`
}

type Default struct {
	Url      string   `yaml:"url"`
	Alerting Alerting `yaml:"alert"`
}
type Alerting struct {
	Month int `yaml:"month"`
}

type Product struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Custom  Custom `yaml:"custom_url"`
}

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
