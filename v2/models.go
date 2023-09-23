package sesv2

import (
	"fmt"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type DataEmail struct {
	// Required at least 1
	ToAddresses []string

	// Required
	From string

	// Required
	FromMsg string

	// Required
	Title string

	// Required
	MsgHTML string

	Charset      string
	BccAddresses []string
	CcAddresses  []string
	Attachments  []Attachment
}

type Attachment struct {
	Data []byte
	Name string
}

func (d *DataEmail) Validate() error {
	switch 0 {
	case len(d.ToAddresses):
		return ErrInvalidTo
	case len(d.From):
		return ErrInvalidFrom
	case len(d.FromMsg):
		return ErrInvalidMessage
	case len(d.Title):
		return ErrInvalidTitle
	case len(d.MsgHTML):
		return ErrInvalidMessageHTML
	}

	title, err := removeAccents(d.Title)
	if err != nil {
		return err
	}

	fromMsg, err := removeAccents(d.FromMsg)
	if err != nil {
		return err
	}

	d.Title = title
	d.FromMsg = fromMsg

	return nil
}

func removeAccents(s string) (string, error) {
	isMn := newisMn()

	t := transform.Chain(norm.NFD, runes.Remove(isMn), norm.NFC)

	result, _, err := transform.String(t, s)
	if err != nil {
		return "", fmt.Errorf("removeAccents.transform.String: %v", err)
	}
	return result, nil
}

type isMn struct{}

func newisMn() *isMn {
	return &isMn{}
}

func (i isMn) Contains(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

// MsgEmailBody ..
type MsgEmailBody struct {
	Status   string `json:"status"`
	Msg      string `json:"msg"`
	Duration string `json:"duration"`
}
