package views

import (
	"github.com/madjlzz/madlens/models"
	"log"
)

const (
	AlertLevelError   string = "danger"
	AlertLevelWarning string = "warning"
	AlertLevelInfo    string = "info"
	AlertLevelSuccess string = "success"

	// AlertMessageGeneric is displayed when any random error
	// is encountered by our backend.
	AlertMessageGeneric string = "Something went wrong. Please try again, and contact us if the problem persists."
)

// Alert is used to render Bootstrap alert messages in templates.
type Alert struct {
	Level   string
	Message string
}

// Data is the top level structure that views expect data
// to come in.
type Data struct {
	Alert *Alert
	User  *models.User
	Yield interface{}
}

func (d *Data) SetAlert(err error) {
	if pErr, ok := err.(PublicError); ok {
		d.Alert = &Alert{
			Level:   AlertLevelError,
			Message: pErr.Public(),
		}
	} else {
		log.Println(err)
		d.Alert = &Alert{
			Level:   AlertLevelError,
			Message: AlertMessageGeneric,
		}
	}
}

func (d *Data) AlertError(message string) {
	d.Alert = &Alert{
		Level:   AlertLevelError,
		Message: message,
	}
}

type PublicError interface {
	error
	Public() string
}
