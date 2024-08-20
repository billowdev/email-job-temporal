package domain

type EmailDto struct {
	Sender       string
	Receiver     string
	Subject      string
	HTMLTemplate string
	CC           []string
}
