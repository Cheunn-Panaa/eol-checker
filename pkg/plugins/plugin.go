package plugins

import (
	"fmt"

	"github.com/cheunn-panaa/eol-checker/configs"
)

// Plugins container
type Plugins interface {
	AddPlugin(name string, plugin *Plugin) error
	GetPlugin(name string) (Plugin, error)
}

// Plugin interface
type Plugin interface {
	SendMessage() interface{}
}

type pluginsContainer struct {
	plugins map[string]*Plugin
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

// newPluginsContainer constructor
func newPluginsContainer() *Plugins {
	var sc Plugins = pluginsContainer{plugins: map[string]*Plugin{}}
	return &sc
}

// Load plugins at start
func Load(config configs.Config) (Plugins, error) {
	sc := *newPluginsContainer()

	if err := sc.AddPlugin("slack", initSlackPlugin(config)); err != nil {
		return nil, err
	}

	return sc, nil
}
