package slack

import (
	"fmt"

	"github.com/cheunn-panaa/eol-checker/plugins"
)

func (plugin *plugins.Plugin) getSendMessage() {

	fmt.Println(plugin.Config.Slack.WebhookURL)

}
