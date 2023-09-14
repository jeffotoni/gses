// @autor: @jeffotoni
package ses

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	fmts "github.com/jeffotoni/gconcat"
	"github.com/jeffotoni/gses/config"
	"github.com/jeffotoni/gses/models"
	"github.com/pkg/errors"
)

// SesEmail  email struct
// We will assemble our data map with this structure
type SesEmail struct {
	config   *config.Config
	ses      *ses.SES
	Profiles map[string]*profile
}

// Ses COnstructor
func NewSesEmail(config *config.Config) *SesEmail {
	return &SesEmail{
		config:   config,
		ses:      &ses.SES{},
		Profiles: make(map[string]*profile),
	}
}

func splitAddr(s string) []string {
	sp := strings.Split(s, ",")

	items := make([]string, len(sp))
	for i := range sp {
		items = append(items, items[i])
	}

	return items
}

// EmailTo string, Cc string, Bc string, Html string, Subject string
//
// func (pf *profile) Send(EmailTo string, Html string, Subject string, Cc string, Bcc string) error {
func (s *SesEmail) send(pf *profile, data *models.DataEmail) error {

	DestinationV := &ses.Destination{}

	for _, v := range splitAddr(data.To) {

		DestinationV.ToAddresses = append(DestinationV.ToAddresses, aws.String(v))
	}

	for _, addr := range splitAddr(data.CcAddresses) {

		DestinationV.CcAddresses = append(DestinationV.CcAddresses, aws.String(addr))
	}

	for _, addr := range splitAddr(data.BccAddresses) {

		DestinationV.BccAddresses = append(DestinationV.BccAddresses, aws.String(addr))
	}

	params := &ses.SendEmailInput{

		Destination: DestinationV,

		Message: &ses.Message{ // Required
			Body: &ses.Body{ // Required
				Html: &ses.Content{
					Data:    aws.String(data.MsgHTML), // Required
					Charset: aws.String("utf-8"),
				},

				//,
				// Text: &ses.Content{
				//     Data:    aws.String("MessageData"), // Required
				//     Charset: aws.String("Charset"),
				// },
			},
			Subject: &ses.Content{ // Required
				Data:    aws.String(data.Title), // Required
				Charset: aws.String("utf-8"),
			},
		},

		Source:           pf.From,
		ReplyToAddresses: pf.ReplyTo,
		//ReturnPath:       pf.ReturnPath,
		//ReturnPathArn:    pf.ReturnPathArn,
		//SourceArn:        pf.SourceArn,
	}

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(s.config.AwsAccess, s.config.AwsSecret, ""),
		Region:      aws.String(pf.Region),
	})

	// sess, err = session.NewSession(&aws.Config{
	// 	Region: aws.String(pf.Region),
	// })

	if err != nil {
		return errors.New(fmts.ConcatStr("send.NewSession:", err.Error()))
	}

	svc := ses.New(sess)
	_, err = svc.SendEmail(params)
	if err != nil {
		return errors.New(fmts.ConcatStr("send.SendMail:", err.Error()))

	}
	return nil
}

// SendEmailSes ..
func (s *SesEmail) SendEmailSes(profileName string, data *models.DataEmail) error {

	if err := data.Validate(); err != nil {
		return err
	}

	if len(s.Profiles) == 0 {
		return errors.Wrap(ErrNoProfileSet, "SendEmailSes len s.Profiles == 0")
	}

	profile, ok := s.Profiles[profileName]
	if !ok {
		return errors.Wrap(ErrProfileNotSearched, "SendEmailSes.Profiles[profileName]")
	}

	return s.send(profile, data)
}
