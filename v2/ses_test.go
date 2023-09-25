package v2

import (
	"context"
	"log"
	"os"
	"testing"
)

func TestSendEmail(t *testing.T) {
	var (
		AWS_REGION            = os.Getenv("AWS_REGION")
		AWS_ACCESS_KEY_ID     = os.Getenv("AWS_ACCESS_KEY_ID")
		AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")

		AWS_FROM = os.Getenv("AWS_FROM")
		AWS_TO1  = os.Getenv("AWS_TO1")
		AWS_TO2  = os.Getenv("AWS_TO2")
	)

	c := NewClient(AWS_REGION, AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY)

	htmlBody := `<h1>Hello World Send Email SES</h1>`

	fname1 := "file1.txt"
	data1, err := os.ReadFile(fname1)
	if err != nil {
		log.Fatal(err)
	}

	req := DataEmail{
		ToAddresses:  []string{AWS_TO1},
		From:         AWS_FROM,
		FromMsg:      "message",
		Title:        "My Test Send Email",
		MsgHTML:      htmlBody,
		BccAddresses: []string{AWS_TO1, AWS_TO2},
		CcAddresses:  []string{AWS_TO1},
		Attachments: []Attachment{
			{Data: data1, Name: "file1.txt"},
			{Data: data1, Name: "file2.txt"},
		},
	}

	if err := c.Send(context.Background(), req); err != nil {
		t.Error(err)
	}
}
