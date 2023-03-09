package ses

import (
	"errors"
)

// SendEmail ..
func SendEmail(To, From, FromMsg, Titulo, MsgHTML string) (err error) {
	if len(To) == 0 {
		return errors.New(`Error: Parameter 'To' is required.`)
	}

	if len(From) == 0 {
		return errors.New(`Error: Parameter 'From' is required.`)
	}

	if len(FromMsg) == 0 {
		return errors.New(`Error: Parameter 'FromMsg' is required.`)
	}

	if len(Titulo) == 0 {
		return errors.New(`Error: Parameter 'Titulo' is required.`)
	}

	if len(MsgHTML) == 0 {
		return errors.New(`Error: Parameter 'MsgHtml' is required.`)
	}

	var Demail = dataEmail{To: To, From: From, FromMsg: FromMsg, Titulo: Titulo, MsgHTML: MsgHTML}
	if SendEmailSes(Demail.To, Demail.From, Demail.FromMsg, Demail.Titulo, Demail.MsgHTML) {
		return nil
	}

	return errors.New(`An error occurried when sending the email`)
}
