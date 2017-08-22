package dms3mail

import (
	"go-distributed-motion-s3/dms3libs"
)

// MailConfig contains dms3mail configuration settings read from TOML file
var MailConfig *structSettings

// motion mail configuration parameters
type structSettings struct {
	Email   *structEmail
	SMTP    *structSMTP
	Logging *dms3libs.StructLogging
}

// email composition parameters
type structEmail struct {
	To   string
	From string
	Body string
}

// SMTP mailer parameters
type structSMTP struct {
	Address            string
	Port               int
	Domain             string
	Username           string
	Password           string
	Authentication     string
	EnableStartTLSAuto bool
}
