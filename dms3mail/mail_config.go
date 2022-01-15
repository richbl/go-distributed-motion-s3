// Package dms3mail configuration structures and variables
//
package dms3mail

import "github.com/richbl/go-distributed-motion-s3/dms3libs"

// mailConfig contains dms3mail configuration settings read from TOML file
var mailConfig *structEmailSettings

// motion mail configuration parameters
type structEmailSettings struct {
	Filename     string
	FileLocation string
	Email        *structEmail
	SMTP         *structSMTP
	Logging      *dms3libs.StructLogging
}

// email composition parameters
type structEmail struct {
	To   string
	From string
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

// event details to be sent via email
type structEventDetails struct {
	eventMedia  string
	eventDate   string
	eventChange string
	clientName  string
}

// HTML tokens used in email template
type emailTemplateElements struct {
	Header string
	Event  string
	Date   string
	Client string
	Change string
	Footer string
}
