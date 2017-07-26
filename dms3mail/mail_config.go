package dms3mail

import (
	"go-distributed-motion-s3/dms3libs"
)

// LogLevel sets the log levels for application logging using the following table:
//
//		0 - OFF, no logging
//		1 - FATAL, report fatal events
//		2 - INFO, report informational events
//    4 - DEBUG, report debugging events
//
const LogLevel = 2

// LogDevice determines to what device logging should be set using the following table:
//
//		0 - STDOUT (terminal)
//		1 - log file
//
// Ignored if LogLevel == 0
//
const LogDevice = 0

// LogFilename is the logging filename
// Ignored if LogLevel == 0 or LogDevice == 0
//
const LogFilename = "dms_mail.log"

// LogLocation is the location of logfile (full path)
// By default, this is in the local folder (e.g., /etc/dms3/dms3server)
// Ignored if LogLevel == 0 or LogDevice == 0
//
var LogLocation = dms3libs.GetPackageDir()

// EmailFrom is the email sender
const EmailFrom = "motion@businesslearninginc.com"

// EmailTo is the email recipient
const EmailTo = "user@gmail.com"

// EmailBody is the email body
// Note that reserved words use the syntax !ALLCAPS and are parsed and replaced
//
const EmailBody = "Motion detected an event of importance. The event (!EVENT) shows !PIXELS pixels changed, and was captured by Camera !CAMERA."

// SMTPAddress is the SMTP address of the receipient
const SMTPAddress = "smtp.gmail.com"

// SMTPPort is the port used by the recipient email account
const SMTPPort = 587

// SMTPDomain is the recieving email domain
const SMTPDomain = "localhost"

// SMTPUsername is the username of the recipient
const SMTPUsername = "user"

// SMTPPassword is the password of the recipient
const SMTPPassword = "password"

// SMTPAuthentication is the email server authentication scheme
const SMTPAuthentication = "plain"

// SMTPEnableStartTLSAuto indicates whether TLS is used
const SMTPEnableStartTLSAuto = true
