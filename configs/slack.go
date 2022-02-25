package configs

type Slack struct {
	WebhookURL  string       `yaml:"webhook_url"`
	Attachments []Attachment `yaml:"attachments"`
	UserName    string       `yaml:"user_name"`
	IconEmoji   string       `yaml:"icon_emoji"`
}
type Attachment struct {
	Author struct {
		Name string `yaml:"name"`
		Link string `yaml:"link"`
		Icon string `yaml:"icon"`
	} `yaml:"author"`
	Color      string `yaml:"color"`
	Fallback   string `yaml:"fallback"`
	Footer     string `yaml:"footer"`
	FooterIcon string `yaml:"footer_icon"`
	Pretext    string `yaml:"pretext"`
	Title      string `yaml:"title"`
	TitleLink  string `yaml:"title_link"`
	ThumbURL   string `yaml:"thumb_url"`
}
