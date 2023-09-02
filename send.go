package ses

import (
	"errors"
)

type Config struct {
	AWS_REGION            string `json:"AWS_REGION,omitempty"`
	AWS_IDENTITY          string `json:"AWS_IDENTITY,omitempty"`
	AWS_ACCESS_KEY_ID     string `json:"AWS_ACCESS_KEY_ID,omitempty"`
	AWS_SECRET_ACCESS_KEY string `json:"AWS_SECRET_ACCESS_KEY,omitempty"`
}

func New(AWS_REGION, AWS_IDENTITY, AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY string) *Config {
	return &Config{
		AWS_REGION:            AWS_REGION,
		AWS_IDENTITY:          AWS_IDENTITY,
		AWS_ACCESS_KEY_ID:     AWS_ACCESS_KEY_ID,
		AWS_SECRET_ACCESS_KEY: AWS_SECRET_ACCESS_KEY,
	}
}

func (c *Config) SendEmailNew(To, From, FromMsg, Titulo, MsgHTML string) (err error) {
	err = valid(To, From, FromMsg, Titulo, MsgHTML)
	if err != nil {
		return err
	}
	var Demail = dataEmail{To: To, From: From, FromMsg: FromMsg, Titulo: Titulo, MsgHTML: MsgHTML}
	if c.SendEmailSesNew(Demail.To, Demail.From, Demail.FromMsg, Demail.Titulo, Demail.MsgHTML) {
		return nil
	}
	return errors.New(`An error occurried when sending the email`)
}

// SendEmail ..
func SendEmail(To, From, FromMsg, Titulo, MsgHTML string) (err error) {
	err = valid(To, From, FromMsg, Titulo, MsgHTML)
	if err != nil {
		return err
	}
	var Demail = dataEmail{To: To, From: From, FromMsg: FromMsg, Titulo: Titulo, MsgHTML: MsgHTML}
	if SendEmailSes(Demail.To, Demail.From, Demail.FromMsg, Demail.Titulo, Demail.MsgHTML) {
		return nil
	}
	return errors.New(`An error occurried when sending the email`)
}

func valid(To, From, FromMsg, Titulo, MsgHTML string) error {
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
	return nil
}
