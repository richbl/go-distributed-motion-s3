// Package dms3mail configuration structures and variables
//
package dms3mail

import "github.com/richbl/go-distributed-motion-s3/dms3libs"

// mailConfig contains dms3mail configuration settings read from TOML file
var mailConfig *structSettings

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
