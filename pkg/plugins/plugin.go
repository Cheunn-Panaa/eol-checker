package plugins

import (
	"fmt"

	"github.com/cheunn-panaa/eol-checker/configs"
	"github.com/cheunn-panaa/eol-checker/internal/utils"
	"github.com/cheunn-panaa/eol-checker/pkg/api"
)

// Plugins container
type Plugins interface {
	AddPlugin(name string, plugin *Plugin) error
	GetPlugin(name string) (Plugin, error)
	GetAllPlugins() (map[string]*Plugin, error)
}

// Plugin interface
type Plugin interface {
	SendMessage([]PluginsMessage) interface{}
}

//pluginsContainer di container
type pluginsContainer struct {
	plugins map[string]*Plugin
}

//PluginsMessqge generic type for each plugin
type PluginsMessage struct {
	Name         string
	Cycle        string
	LatestCycle  utils.StringOrInt
	Release      string
	EOL          utils.StringOrBool
	Latest       string
	Link         string
	LTS          bool
	Discontinued string
}

// AddPlugin method to add a plugin to plugin container
func (sc pluginsContainer) AddPlugin(name string, plugin *Plugin) error {
	if _, ok := sc.plugins[name]; ok {
		return fmt.Errorf("'%s' plugin name is already set", name)
	}

	sc.plugins[name] = plugin

	return nil
}

// GetPlugin method
func (sc pluginsContainer) GetPlugin(name string) (Plugin, error) {
	if _, ok := sc.plugins[name]; ok {
		return *sc.plugins[name], nil
	}

	return nil, fmt.Errorf("'%s' plugin name is not set", name)
}

// GetAllPlugins retrieve all listed plugins
func (sc pluginsContainer) GetAllPlugins() (map[string]*Plugin, error) {
	if len(sc.plugins) > 0 {
		return sc.plugins, nil
	}

	return nil, fmt.Errorf("plugin list is empty")
}

// newPluginsContainer constructor
func newPluginsContainer() *Plugins {
	var sc Plugins = pluginsContainer{plugins: map[string]*Plugin{}}
	return &sc
}

// Load plugins at start
func Load(config *configs.Configuration) (Plugins, error) {
	sc := *newPluginsContainer()

	if err := sc.AddPlugin("slack", initSlackPlugin(config)); err != nil {
		return nil, err
	}

	return sc, nil
}

//MessageBuilder will generate a message from given objects
func MessageBuilder(project *api.ProjectCycle, product *configs.Product) PluginsMessage {
	return PluginsMessage{
		Name:        product.Name,
		Cycle:       product.Version,
		LatestCycle: project.LatestCycle,
		Release:     project.Release,
		EOL:         project.EOL,
		Latest:      project.Latest,
		LTS:         project.LTS,
	}
}
