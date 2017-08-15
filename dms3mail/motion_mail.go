package dms3mail

import (
	"flag"
	"go-distributed-motion-s3/dms3libs"
	"path"
	"strconv"
	"strings"
	"time"

	"gopkg.in/gomail.v2"
)

type structEventDetails struct {
	eventNumber    string
	eventMedia     string
	eventDate      string
	cameraNumber   int
	pixelsDetected int
}

// Init configs the library and configuration for dms3mail
func Init() {

	dms3libs.LoadLibConfig("/etc/distributed-motion-s3/dms3libs/dms3libs.toml")
	LoadMailConfig("/etc/distributed-motion-s3/dms3mail/dms3mail.toml")

	cfg := MailConfig.Logging
	dms3libs.CreateLogger(cfg.LogLevel, cfg.LogDevice, cfg.LogLocation, cfg.LogFilename)

	GenerateEventEmail()

}

// GenerateEventEmail is the entry point for this package, first calling parseEvent to interpret
// the Motion event, and then calling generateSMTPEmail to create and send the email
//
func GenerateEventEmail() {

	var eventDetails structEventDetails

	parseEvent(&eventDetails)
	generateSMTPEmail(&eventDetails)

}

// parseEvent creates an event by parsing the following command line arguments passed in via the
//  Motion on_picture_save or the on_movie_end command:
//
//  ARGV[0] pixels detected
//  ARGV[1] media filename
//  ARGV[2] device (camera) number
//
func parseEvent(eventDetails *structEventDetails) {

	pixels := flag.Int("pixels", 0, "count of pixels detected in the event")
	filename := flag.String("filename", "", "fullpath filename of the event media file")
	camera := flag.Int("camera", 0, "camera number that captured the event")

	flag.Parse()

	if flag.NFlag() != 3 {
		dms3libs.LogFatal("only " + strconv.Itoa(flag.NFlag()) + " argument(s) passed... exiting")
	}

	if !dms3libs.IsFile(*filename) {
		dms3libs.LogFatal("filename not found... exiting")
	}

	eventDetails.cameraNumber = *camera
	eventDetails.eventMedia = *filename
	eventDetails.pixelsDetected = *pixels
	eventDetails.eventNumber, eventDetails.eventDate = getEventDetails(*filename)

}

// getEventDetails creates the following event details based on filename:
//
//   eventNumber - Motion-generated event number
//   eventDate - Motion-generated event datetime
//
// NOTE: this method assumes that filename follows the default Motion file-naming convention of
//  %v-%Y%m%d%H%M%S (for movies) or %v-%Y%m%d%H%M%S-%q (for pictures), where:
//
//   %v - Motion-generated event number
//   %Y%m%d%H%M%S - ISO 8601 date, with hours, minutes, seconds notion
//   %q - frame number (value ignored)
//
func getEventDetails(filename string) (eventNumber string, eventDate string) {

	file := path.Base(filename)
	sepCount := strings.Count(file, "-")

	if sepCount != 1 && sepCount != 2 {
		dms3libs.LogFatal("bad file-naming convention... exiting")
	}

	res := strings.Split(file, "-")
	year, _ := strconv.Atoi(res[1][0:4])
	month, _ := strconv.Atoi(res[1][4:6])
	day, _ := strconv.Atoi(res[1][6:8])
	hour, _ := strconv.Atoi(res[1][8:10])
	min, _ := strconv.Atoi(res[1][10:12])
	sec, _ := strconv.Atoi(res[1][12:14])

	return res[0], time.Date(year, time.Month(month), day, hour, min, sec, 0, time.UTC).Format("2006-01-02 at 15:04:05")

}

// createEmailBody performs a placeholder replacement in the email body with eventDetails elements
func createEmailBody(eventDetails *structEventDetails) string {

	var replacements = map[string]string{
		"!EVENT":  eventDetails.eventNumber,
		"!PIXELS": strconv.Itoa(eventDetails.pixelsDetected),
		"!CAMERA": strconv.Itoa(eventDetails.cameraNumber),
	}

	processedEmailBody := MailConfig.EmailBody

	for key, val := range replacements {
		processedEmailBody = strings.Replace(processedEmailBody, key, val, -1)
	}

	return processedEmailBody

}

// generateSMTPEmail generates and mails an email message based on configuration options
// See https://github.com/go-gomail/gomail for mail package options
//
func generateSMTPEmail(eventDetails *structEventDetails) {

	mail := gomail.NewMessage()
	mail.SetHeader("From", MailConfig.EmailFrom)
	mail.SetHeader("To", MailConfig.EmailTo)
	mail.SetHeader("Subject", "Motion Detected on Camera #"+strconv.Itoa(eventDetails.cameraNumber)+" at "+eventDetails.eventDate)
	mail.SetBody("text/html", createEmailBody(eventDetails))
	mail.Attach(eventDetails.eventMedia)

	dialer := gomail.NewDialer(MailConfig.SMTPAddress, MailConfig.SMTPPort, MailConfig.SMTPUsername, MailConfig.SMTPPassword)

	if err := dialer.DialAndSend(mail); err != nil {
		dms3libs.LogFatal(err.Error())
	} else {
		dms3libs.LogInfo("successfully processed and sent email")
	}

}
