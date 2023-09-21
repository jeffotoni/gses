package models

import (
	"fmt"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// DataEmail ..
type DataEmail struct {
	ToAddresses  []string     `json:"toAddresses"`
	From         string       `json:"from"`
	FromMsg      string       `json:"frommsg"`
	Title        string       `json:"titulo"`
	MsgHTML      string       `json:"msghtml"`
	Charset      string       `json:"charset"`
	BccAddresses []string     `json:"bccAddresses"`
	CcAddresses  []string     `json:"ccAddresses"`
	Attachments  []Attachment `json:"attachments"`
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
