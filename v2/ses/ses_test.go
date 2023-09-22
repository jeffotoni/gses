package ses

import (
	"context"
	"os"
	"testing"

	"github.com/jeffotoni/gses/v2/models"
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

	htmlBody := `<h1>Hello World</h1>`

	req := models.DataEmail{
		ToAddresses:  []string{AWS_TO1},
		From:         AWS_FROM,
		FromMsg:      "message",
		Title:        "title",
		MsgHTML:      htmlBody,
		BccAddresses: []string{AWS_TO1, AWS_TO2},
		CcAddresses:  []string{AWS_TO1},
		Attachments: []models.Attachment{
			{Data: []byte("text 1"), Name: "file1.txt"},
			{Data: []byte("text 2"), Name: "file2.txt"},
		},
	}

	if err := c.Send(context.Background(), req); err != nil {
		t.Error(err)
	}
}
