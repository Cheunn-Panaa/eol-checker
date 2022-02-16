package plugins

import (
	"github.com/rs/zerolog"
)

type Plugin struct {
	Logger *zerolog.Logger,
	Config *Config
}

func LoadPlugin(logger *zerolog.Logger) *Plugin {
	conf, err := loadConfig("plugin.conf.yaml")
	return &Plugin{
		Logger: logger,
		Config: conf,
	}
}



// loadConfig is constructor
func loadConfig(filename string) (config *Config, err error) {
	b, err := ioutil.ReadFile(filename)
	err = yaml.Unmarshal(b, &config)
	return
}
