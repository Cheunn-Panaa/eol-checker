package configs

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

// SetConfiguration is constructor
func SetConfiguration(conf *Configuration) {
	Config = conf
}

func GetConfig() *Configuration {
	return Config
}
