// @autor: @jeffotoni
package ses

import (
	"errors"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	fmts "github.com/jeffotoni/gconcat"

	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var timeout = time.Duration(2 * time.Second)

// SesEmail  email struct
// We will assemble our data map with this structure
type SesEmail struct {
	ses      *ses.SES
	Profiles map[string]*profile
}

// profile
type profile struct {
	From          *string
	Sfrom         string
	ReplyTo       []*string
	ReturnPath    *string
	ReturnPathArn *string
	SourceArn     *string
	Region        string
}

var (
	AWS_REGION            = os.Getenv("AWS_REGION")
	AWS_IDENTITY          = os.Getenv("AWS_IDENTITY")
	AWS_ACCESS_KEY_ID     = os.Getenv("AWS_ACCESS_KEY_ID")
	AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
)

func init() {
	if len(AWS_REGION) == 0 ||
		len(AWS_IDENTITY) == 0 {
		log.Println("Error need to export the environment variables: AWS_REGION | AWS_IDENTITY | It's not mandatory => AWS_ACCESS_KEY_ID | AWS_SECRET_ACCESS_KEY")
	}
}

// Setup a profile to use with Send
// With this function, we were able to make the subtitle work correctly
func (this *SesEmail) SetSetupProfile(name string, from string, replyTo []string, returnPath string, returnPathArn string, sourceArn string, region string) bool {

	//
	// or data map
	//
	this.Profiles = map[string]*profile{}

	//
	// profiles
	//
	this.Profiles[name] = &profile{

		From:    aws.String(from),
		Sfrom:   from,
		ReplyTo: []*string{},
		//ReturnPath:    aws.String(returnPath),
		//ReturnPathArn: aws.String(returnPathArn),
		//SourceArn:     aws.String(sourceArn),
		Region: region,
	}

	//
	// for
	//
	for _, d := range replyTo {
		this.Profiles[name].ReplyTo = append(this.Profiles[name].ReplyTo, aws.String(d))
	}

	return true
}

func AwsSesSetProfile(Region string, IdentityArn string, From string, Info string) *profile {

	IdentityARN := fmts.Concat("arn:aws:ses:", Region, ":", IdentityArn, ":identity/", From)

	//
	//
	//
	FromX := fmts.ConcatStr(Info, " <", From, ">")

	//
	//
	//
	ReturnPathx := IdentityARN

	//
	//
	//
	ReturnPathxArm := IdentityARN

	//
	// config email
	//
	sender := new(SesEmail)

	//
	//
	//
	sender.SetSetupProfile("default", FromX, []string{From},
		From,
		ReturnPathx,
		ReturnPathxArm, Region)

	prof := sender.Profiles["default"]
	if prof == nil {
		log.Println("Error profiles: ", prof)
		return nil
	}

	return prof
}

// EmailTo string, Cc string, Bc string, Html string, Subject string
//
// func (pf *profile) Send(EmailTo string, Html string, Subject string, Cc string, Bcc string) error {
func (pf *profile) Send(paramses ...string) error {

	var EmailTo string
	var Cc string
	var Bcc string
	var HTML string
	var Subject string

	if len(paramses) > 0 && len(paramses) <= 5 {

		if len(paramses) == 5 {

			EmailTo = paramses[0]

			HTML = paramses[1]

			Subject = paramses[2]

			Cc = paramses[3]

			Bcc = paramses[4]

		} else if len(paramses) == 4 {

			EmailTo = paramses[0]

			HTML = paramses[1]

			Subject = paramses[2]

			Cc = paramses[3]

			Bcc = ""

		} else if len(paramses) == 3 {

			EmailTo = paramses[0]

			HTML = paramses[1]

			Subject = paramses[2]

			Cc = ""

			Bcc = ""

		} else if len(paramses) == 2 {

			EmailTo = paramses[0]

			HTML = paramses[1]

			Subject = ""

			Cc = ""

			Bcc = ""

		} else if len(paramses) == 1 {

			EmailTo = paramses[0]

			HTML = ""

			Subject = ""

			Cc = ""

			Bcc = ""
		}

	} else {
		return errors.New("Error Parameters is missing")
	}

	//
	//
	//
	if EmailTo == "" {
		return errors.New("Error Parameters EmailTo Required")
	}

	//
	//
	//
	DestinationV := &ses.Destination{}
	_ = DestinationV
	//
	//
	//
	ToAddressesMail := []*string{}

	//
	//
	//
	CcAddressesMail := []*string{}

	//
	//
	//
	BccAddressesMail := []*string{}

	//
	//
	//
	EmailTo = strings.Trim(EmailTo, " ")

	//
	//
	//
	arrayMailTo := strings.Split(EmailTo, ",")

	for i := range arrayMailTo {

		//
		//
		//
		mailClean := strings.TrimSpace(arrayMailTo[i])

		//
		//
		//
		ToAddressesMail = append(ToAddressesMail, aws.String(mailClean))
	}

	//
	//
	//
	if Cc != "" {

		//
		//
		//
		Cc = strings.Trim(Cc, " ")

		//
		//
		//
		arrayMailCc := strings.Split(Cc, ",")

		for i := range arrayMailCc {

			//
			//
			//
			mailCcClean := strings.TrimSpace(arrayMailCc[i])

			//
			//
			//
			CcAddressesMail = append(CcAddressesMail, aws.String(mailCcClean))
		}
	}

	//
	//
	//
	if Bcc != "" {

		//
		//
		//
		Bcc = strings.Trim(Bcc, " ")

		//
		//
		//
		arrayMailBCc := strings.Split(Bcc, ",")

		for i := range arrayMailBCc {

			//
			//
			//
			mailBccClean := strings.TrimSpace(arrayMailBCc[i])

			//
			//
			//
			BccAddressesMail = append(BccAddressesMail, aws.String(mailBccClean))
		}
	}

	//
	//
	//
	if Cc != "" && Bcc == "" {

		DestinationV = &ses.Destination{ // Required

			CcAddresses: CcAddressesMail,

			ToAddresses: ToAddressesMail,
		}

	} else if Cc == "" && Bcc != "" {

		DestinationV = &ses.Destination{ // Required

			BccAddresses: BccAddressesMail,

			ToAddresses: ToAddressesMail,
		}
	} else if Cc != "" && Bcc != "" {

		DestinationV = &ses.Destination{ // Required

			BccAddresses: BccAddressesMail,

			CcAddresses: CcAddressesMail,

			ToAddresses: ToAddressesMail,
		}
	} else {

		DestinationV = &ses.Destination{ // Required

			ToAddresses: ToAddressesMail,
		}
	}

	params := &ses.SendEmailInput{

		//
		//
		//
		Destination: DestinationV,

		Message: &ses.Message{ // Required
			Body: &ses.Body{ // Required
				Html: &ses.Content{
					Data:    aws.String(HTML), // Required
					Charset: aws.String("utf-8"),
				},

				//,
				// Text: &ses.Content{
				//     Data:    aws.String("MessageData"), // Required
				//     Charset: aws.String("Charset"),
				// },
			},
			Subject: &ses.Content{ // Required
				Data:    aws.String(Subject), // Required
				Charset: aws.String("utf-8"),
			},
		},

		Source:           pf.From,
		ReplyToAddresses: pf.ReplyTo,
		//ReturnPath:       pf.ReturnPath,
		//ReturnPathArn:    pf.ReturnPathArn,
		//SourceArn:        pf.SourceArn,
	}

	var sess *session.Session
	var err error
	if len(AWS_ACCESS_KEY_ID) > 0 && len(AWS_SECRET_ACCESS_KEY) > 0 {
		sess, err = session.NewSession(&aws.Config{
			Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, ""),
			Region:      aws.String(pf.Region),
		})
	} else {
		sess, err = session.NewSession(&aws.Config{
			Region: aws.String(pf.Region),
		})
	}

	if err != nil {
		return errors.New(fmts.ConcatStr("Error Error creating session:", err.Error()))
	}

	svc := ses.New(sess)
	_, err = svc.SendEmail(params)
	if err != nil {
		return errors.New(fmts.ConcatStr("error aws sendEmail:", err.Error()))

	}
	return nil
}

// SendEmailSes ..
func (config *Config) SendEmailSesNew(To, From, FromMsg, Titulo, MsgHTML string) bool {

	To = strings.TrimSpace(To)
	From = strings.TrimSpace(From)
	FromMsg = strings.TrimSpace(FromMsg)
	Titulo = strings.TrimSpace(Titulo)

	FromMsg = removeAccents(FromMsg)
	Titulo = removeAccents(Titulo)

	//
	//
	//
	var err error

	S := AwsSesSetProfile(

		config.AWS_REGION,

		config.AWS_IDENTITY,

		From, // AWS_FROM

		FromMsg, // AWS_MSG
	)

	err = S.Send(

		To,

		MsgHTML,

		Titulo,
	)

	if err != nil {
		log.Println("Error sending SES email: ", To, " error: ", err)
		return false

	}
	return true

}

// SendEmailSes ..
func SendEmailSes(To, From, FromMsg, Titulo, MsgHTML string) bool {

	To = strings.TrimSpace(To)
	From = strings.TrimSpace(From)
	FromMsg = strings.TrimSpace(FromMsg)
	Titulo = strings.TrimSpace(Titulo)

	FromMsg = removeAccents(FromMsg)
	Titulo = removeAccents(Titulo)

	//
	//
	//
	var err error

	S := AwsSesSetProfile(

		AWS_REGION,

		AWS_IDENTITY,

		From, // AWS_FROM

		FromMsg, // AWS_MSG
	)

	err = S.Send(

		To,

		MsgHTML,

		Titulo,
	)

	if err != nil {
		log.Println("Error sending SES email: ", To, " error: ", err)
		return false

	}
	return true

}

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

func removeAccents(s string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
