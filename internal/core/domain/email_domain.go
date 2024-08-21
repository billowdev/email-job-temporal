package domain

type EmailDto struct {
	Sender       string   `json:"sender"`
	Receiver     string   `json:"receiver"`
	Subject      string   `json:"subject"`
	HTMLTemplate string   `json:"html_template"`
	CC           []string `json:"cc"`
}
