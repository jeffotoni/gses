package models

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// DataEmail ..
type DataEmail struct {
	To      string `json:"to"`
	From    string `json:"from"`
	FromMsg string `json:"frommsg"`
	Title   string `json:"titulo"`
	MsgHTML string `json:"msghtml"`

	BccAddresses string
	// The recipients to place on the CC: line of the message.
	CcAddresses string
}

func NewDataEmail(to, from, message, title, htmlMsgBody, bccAddress, ccAddress string) *DataEmail {
	return &DataEmail{
		To:           strings.TrimSpace(to),
		From:         strings.TrimSpace(from),
		FromMsg:      strings.TrimSpace(message),
		Title:        strings.TrimSpace(title),
		MsgHTML:      htmlMsgBody,
		BccAddresses: strings.TrimSpace(bccAddress),
		CcAddresses:  strings.TrimSpace(ccAddress),
	}
}

func (d *DataEmail) Validate() error {
	switch 0 {
	case len(d.To):
		return ErrInvalidTo{}
		// case len(d.From):
		// 	return ErrInvalidFrom{}
		// case len(d.FromMsg):
		// 	return ErrInvalidMessage{}
		// case len(d.Title):
		// 	return ErrInvalidTitle{}
		// case len(d.MsgHTML):
		// 	return ErrInvalidMessageHTML{}
	}

	var err error

	d.Title, err = removeAccents(d.Title)
	if err != nil {
		return err
	}

	d.FromMsg, err = removeAccents(d.FromMsg)
	if err != nil {
		return err
	}

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
