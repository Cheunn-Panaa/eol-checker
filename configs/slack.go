package configs

//Slack base config for slack's plugin
type Slack struct {
	WebhookURL  string       `mapstructure:"webhook_url"`
	Attachments []Attachment `mapstructure:"attachments"`
	UserName    string       `mapstructure:"user_name"`
	IconEmoji   string       `mapstructure:"icon_emoji"`
}

//An Attachement Slack attachments conf
type Attachment struct {
	Author struct {
		Name string `mapstructure:"name"`
		Link string `mapstructure:"link"`
		Icon string `mapstructure:"icon"`
	} `mapstructure:"author"`
	Color      string `mapstructure:"color"`
	Fallback   string `mapstructure:"fallback"`
	Footer     string `mapstructure:"footer"`
	FooterIcon string `mapstructure:"footer_icon"`
	Pretext    string `mapstructure:"pretext"`
	Title      string `mapstructure:"title"`
	TitleLink  string `mapstructure:"title_link"`
	ThumbURL   string `mapstructure:"thumb_url"`
}
