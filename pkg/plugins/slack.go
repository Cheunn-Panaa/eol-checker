package plugins

import (
	"fmt"

	"github.com/cheunn-panaa/eol-checker/configs"
)

type SlackPlugin struct {
	conf *configs.Slack
}

// SendMessage method
func (s SlackPlugin) SendMessage() interface{} {
	fmt.Println("plugin.Config.Slack.WebhookURL")
	return nil
}

func initSlackPlugin(config configs.Config) *Plugin {
	var plugin Plugin

	plugin = SlackPlugin{conf: &config.Plugins.Slack}
	return &plugin
}
