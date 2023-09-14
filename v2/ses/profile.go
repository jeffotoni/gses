package ses

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	fmts "github.com/jeffotoni/gconcat"
)

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

// Setup a profile to use with Send
// With this function, we were able to make the subtitle work correctly
func (s *SesEmail) createProfile(name string,
	from string,
	replyTo []string,
	returnPath string,
	returnPathArn string,
	sourceArn string,
	region string) error {

	//
	// profiles
	//
	profile := profile{

		From:    aws.String(from),
		Sfrom:   from,
		ReplyTo: []*string{},
		//ReturnPath:    aws.String(returnPath),
		//ReturnPathArn: aws.String(returnPathArn),
		//SourceArn:     aws.String(sourceArn),
		Region: region,
	}

	if _, ok := s.Profiles[name]; ok {
		return errors.New("profile name exist")
	}

	s.Profiles[name] = &profile

	for _, d := range replyTo {
		s.Profiles[name].ReplyTo = append(s.Profiles[name].ReplyTo, aws.String(d))
	}

	return nil
}

func (s *SesEmail) AddProfile(
	// name of profile
	profileName string,

	//
	// region aws ex: us-east-1
	//
	Region string,

	//
	// https://console.aws.amazon.com/ses
	// Identity ARN: arn:aws:ses:region-aws:xxxx:identity/yourmail@domain.com
	//
	IdentityArn string,

	//
	// Mail that it will send, it has to be configured on your SES
	//
	From string,

	//
	// Message that will be displayed in from ex: "text info here <emailfrom@domain.com>"
	//
	Info string) (*profile, error) {

	//
	// Identity ARN: arn:aws:ses:region-aws:xxxx:identity/yourmail@domain.com
	//
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
	//
	//

	err := s.createProfile(
		profileName,
		FromX, []string{From},
		From,
		ReturnPathx,
		ReturnPathxArm,
		Region,
	)

	return s.Profiles[profileName], err
}
