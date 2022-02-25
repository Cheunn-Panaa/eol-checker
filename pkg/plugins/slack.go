package plugins

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cheunn-panaa/eol-checker/configs"
	"github.com/cheunn-panaa/eol-checker/pkg/utils"
)

type SlackPlugin struct {
	conf *configs.Slack
}

// Attachment for a Slack message
// https://api.slack.com/docs/message-attachments
type Attachment struct {
	AuthorIcon string   `json:"author_icon"`
	AuthorLink string   `json:"author_link"`
	AuthorName string   `json:"author_name"`
	Color      string   `json:"color"`
	Fallback   string   `json:"fallback"`
	Fields     []Field  `json:"fields"`
	FooterIcon string   `json:"footer_icon"`
	Footer     string   `json:"footer"`
	ImageURL   string   `json:"image_url"`
	MrkdwnIn   []string `json:"mrkdwn_in"`
	Pretext    string   `json:"pretext"`
	Text       string   `json:"text"`
	ThumbURL   string   `json:"thumb_url"`
	TitleLink  string   `json:"title_link"`
	Title      string   `json:"title"`
	Ts         int64    `json:"ts"`
}

// Field of an attachment
type Field struct {
	Short bool   `json:"short"`
	Title string `json:"title"`
	Value string `json:"value"`
}

// Payload is a Slack message with attachments
type Payload struct {
	Attachments []Attachment `json:"attachments"`
	LinkNames   bool         `json:"link_names"`
	Mrkdwn      bool         `json:"mrkdwn"`
	IconEmoji   string       `json:"icon_emoji"`
	Username    string       `json:"username"`
	Channel     string       `json:"channel"`
	Thread      string       `json:"thread_ts,omitempty"`
}

// SendMessage method
func (s SlackPlugin) SendMessage(productList []PluginsMessage) interface{} {

	attachments := s.generateAttachments(productList)

	payload := Payload{
		Attachments: attachments,
		LinkNames:   true,
		Mrkdwn:      true,
		Username:    s.conf.UserName,
		IconEmoji:   s.conf.IconEmoji,
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(data)
	request, err := http.NewRequest("POST", s.conf.WebhookURL, body)
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}

func (s SlackPlugin) generateAttachments(productList []PluginsMessage) []Attachment {
	var attachments []Attachment
	if s.conf.Attachments != nil {
		for _, product := range productList {
			attachments = append(attachments, Attachment{
				MrkdwnIn:   []string{"text", "pretext"},
				AuthorIcon: s.conf.Attachments[0].Author.Icon,
				AuthorName: s.conf.Attachments[0].Author.Name,
				Color:      colorHandler(product.EOL),
				Text:       generateText(product.EOL),
				Title:      s.conf.Attachments[0].Title,
				ThumbURL:   s.conf.Attachments[0].ThumbURL,
				Fields:     s.generateFields(product),
			})
		}
	}
	return attachments
}

func (s SlackPlugin) generateFields(product PluginsMessage) []Field {

	var fields []Field
	var message string
	if product.EOL.Bool {
		message = "No dates have been provided."
	} else {
		message = utils.DateFormat(product.EOL.String, "02 January 2006")
	}
	fields = addField(fields, "Product", product.Name, "true")
	fields = addField(fields, "Your version", product.Cycle, "true")
	fields = addField(fields, "Released in", utils.DateFormat(product.Release, "02 January 2006"), "true")
	fields = addField(fields, "End of Life", message, "true")
	fields = addField(fields, "Latest minor version", product.Latest, "true")
	fields = addField(fields, "Latest major version", product.LatestCycle.ToString(), "true")
	return fields
}
func addField(fields []Field, fieldTitle string, fieldValue string, fieldShort string) []Field {
	if fieldTitle != "" && fieldValue != "" {
		var short = false
		var err error

		if fieldShort != "" {
			if short, err = strconv.ParseBool(fieldShort); err != nil {
				fmt.Println("slack-notifier: Error: ", err.Error())
				short = false
			}
		}

		fields = append(fields, Field{
			Title: fieldTitle,
			Value: fieldValue,
			Short: short,
		})
	}

	return fields
}

func colorHandler(eol utils.StringOrBool) string {
	if eol.String != "" {

		eolDate, _ := time.Parse("2006-01-02", eol.String)
		if eolDate.Before(time.Now()) {
			return "danger"
		}
		if eolDate.Before(time.Now().AddDate(0, 12, 0)) {
			return "warning"
		}
	} else if eol.Bool {
		return "danger"
	}
	return "good"
}

func generateText(eol utils.StringOrBool) string {
	if eol.String != "" {

		eolDate, _ := time.Parse("2006-01-02", eol.String)
		if eolDate.Before(time.Now()) {
			return "This product has reached it's end of life"
		}
		if eolDate.Before(time.Now().AddDate(0, 12, 0)) {
			return "This product is near it's end of life we recommend you upgrade to latest stable version."
		}
	} else if eol.Bool {
		return "This product has reached it's end of life"
	}
	return "good"
}

func initSlackPlugin(config *configs.Config) *Plugin {
	var plugin Plugin
	plugin = SlackPlugin{conf: &config.Plugins.Slack}
	return &plugin
}
