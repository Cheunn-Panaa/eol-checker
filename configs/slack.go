package configs

type Slack struct {
	WebhookURL string `yaml:"webhook_url"`
	UserName   string `yaml:"user_name"`
	IconEmoji  string `yaml:"icon_emoji"`
	Fallback   string `yaml:"fallback"`
	Color      string `yaml:"color"`
	Pretext    string `yaml:"pretext"`
	Author     struct {
		Name string `yaml:"name"`
		Link string `yaml:"link"`
		Icon string `yaml:"icon"`
	} `yaml:"author"`
	Title      string `yaml:"title"`
	TitleLink  string `yaml:"title_link"`
	ThumbURL   string `yaml:"thumb_url"`
	Footer     string `yaml:"footer"`
	FooterIcon string `yaml:"footer_icon"`
	Fields     struct {
		Field1 struct {
			Title string `yaml:"title"`
			Value string `yaml:"value"`
			Short bool   `yaml:"short"`
		} `yaml:"field1"`
		Field2 struct {
			Title string `yaml:"title"`
			Value string `yaml:"value"`
			Short bool   `yaml:"short"`
		} `yaml:"field2"`
	} `yaml:"fields"`
} `yaml:"slack"`