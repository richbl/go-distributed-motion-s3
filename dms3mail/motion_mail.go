// Package dms3mail implements a mailer service for dms3clients
package dms3mail

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"math"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/richbl/go-distributed-motion-s3/dms3libs"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/gomail.v2"
)

// Init configs the library and configuration for dms3mail
func Init(configPath string) {

	dms3libs.LoadLibConfig(filepath.Join(configPath, "dms3libs", "dms3libs.toml"))
	dms3libs.LoadComponentConfig(&mailConfig, filepath.Join(configPath, "dms3mail", "dms3mail.toml"))

	dms3libs.SetLogFileLocation(mailConfig.Logging)
	dms3libs.CreateLogger(mailConfig.Logging)
	dms3libs.LogInfo("dms3mail " + dms3libs.GetProjectVersion() + " started")

	dms3libs.CheckFileLocation(configPath, "dms3mail", &mailConfig.FileLocation, mailConfig.Filename)

	GenerateEventEmail()

}

// GenerateEventEmail is the entry point for this package, first calling parseEventArgs to interpret
// the Motion event, and then calling generateSMTPEmail to create and send the email
func GenerateEventEmail() {

	var eventDetails structEventDetails

	eventDetails.parseEventArgs()
	eventDetails.generateSMTPEmail()

}

// parseEventArgs creates an event by parsing the following command line arguments passed in via the
// Motion on_picture_save or the on_movie_end command:
// ARGV[0] count of pixels changed
// ARGV[1] media filename
func (eventDetails *structEventDetails) parseEventArgs() {

	// parse command line arguments passed from Motion command
	pixels := flag.Int("pixels", 0, "count of pixels detected in the event")
	filename := flag.String("filename", "", "fullpath filename of the event media file")

	flag.Parse()

	if flag.NFlag() != 2 {
		dms3libs.LogFatal("only " + strconv.Itoa(flag.NFlag()) + " argument(s) passed... exiting")
	} else if !dms3libs.IsFile(*filename) {
		dms3libs.LogFatal("filename not found... exiting")
	}

	eventDetails.eventMedia = *filename

	// get image dimensions and calculate percent of image change
	width, height := dms3libs.GetImageDimensions(eventDetails.eventMedia)
	eventDetails.eventChange = fmt.Sprintf("%d", int(math.Ceil((float64(*pixels) / float64(width*height) * 100))))

	eventDetails.eventDate = getEventDetails(eventDetails.eventMedia)
	eventDetails.clientName = cases.Title(language.English, cases.NoLower).String(dms3libs.GetDeviceHostname())

}

// getEventDetails creates the following event details based on filename:
//
//	eventNumber - Motion-generated event number
//	eventDate - Motion-generated event datetime
//
// This method assumes that filename follows the default Motion file-naming convention of
// [%v-]%Y%m%d%H%M%S (for movies) or [%v-]%Y%m%d%H%M%S-%q (for pictures), where:
//
//	[%v] - event number (as of Motion 4.3.2, no longer included in filename by default)
//	%Y%m%d%H%M%S - ISO 8601 date, with hours, minutes, seconds notion
//	%q - frame number (value ignored)
func getEventDetails(filename string) (eventDate string) {

	var index int
	file := path.Base(filename)
	sepCount := strings.Count(file, "-")

	if sepCount > 0 && sepCount < 3 {
		index = sepCount - 1
	} else {
		dms3libs.LogFatal("unexpected Motion filenaming convention: missing separators")
	}

	res := strings.Split(file, "-")

	if len(res[index]) != 14 {
		dms3libs.LogFatal("unexpected Motion filenaming convention: incorrect string length")
	}

	year, _ := strconv.Atoi(res[index][0:4])
	month, _ := strconv.Atoi(res[index][4:6])
	day, _ := strconv.Atoi(res[index][6:8])
	hour, _ := strconv.Atoi(res[index][8:10])
	min, _ := strconv.Atoi(res[index][10:12])
	sec, _ := strconv.Atoi(res[index][12:14])

	return time.Date(year, time.Month(month), day, hour, min, sec, 0, time.UTC).Format("15:04:05 on 2006-01-02")

}

// createEmailBody loads the email template (HTML) and parses elements
func (elements emailTemplateElements) createEmailBody() string {

	t := template.Must(template.New(mailConfig.Filename).ParseFiles(filepath.Join(mailConfig.FileLocation, mailConfig.Filename)))

	var tpl bytes.Buffer

	if err := t.Execute(&tpl, elements); err != nil {
		dms3libs.LogFatal(err.Error())
	}

	return tpl.String()

}

// generateSMTPEmail generates and mails an email message based on configuration options
// See https://github.com/go-gomail/gomail for mail package options
func (eventDetails *structEventDetails) generateSMTPEmail() {

	mail := gomail.NewMessage()
	mail.SetHeader("From", mailConfig.Email.From)
	mail.SetHeader("To", mailConfig.Email.To)
	mail.SetHeader("Subject", "DMS3 Client "+eventDetails.clientName+": Event Detected at "+eventDetails.eventDate)

	headerImage := filepath.Join(mailConfig.FileLocation, "assets", "img", "dms3logo.png")
	footerImage := filepath.Join(mailConfig.FileLocation, "assets", "img", "dms3github.png")

	elements := &emailTemplateElements{
		Header: filepath.Base(headerImage),
		Event:  filepath.Base(eventDetails.eventMedia),
		Date:   eventDetails.eventDate,
		Client: eventDetails.clientName,
		Change: eventDetails.eventChange,
		Footer: filepath.Base(footerImage),
	}

	mail.SetBody("text/html", elements.createEmailBody())

	mail.Embed(headerImage)
	mail.Embed(eventDetails.eventMedia)
	mail.Embed(footerImage)

	dialer := gomail.NewDialer(mailConfig.SMTP.Address, mailConfig.SMTP.Port, mailConfig.SMTP.Username, mailConfig.SMTP.Password)

	if err := dialer.DialAndSend(mail); err != nil {
		dms3libs.LogFatal(err.Error())
	} else {
		dms3libs.LogInfo("successfully processed and sent email")
	}

}
