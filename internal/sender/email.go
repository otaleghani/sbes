package sender

// Structure of the email
type Email struct {
	// Authentication
	SmtpHost string
	SmtpPort int
	Username string
	Password string
	Oauth    string

	// Message
	From        string
	MailingList []string
	Subject     string
	Body        string
	MsgType     string
	Campaign    string
	Domain      string
}
