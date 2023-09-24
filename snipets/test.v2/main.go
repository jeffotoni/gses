package main

import (
	"context"
	"log"
	"os"

	sesv2 "github.com/jeffotoni/gses/v2"
)

var (
	AWS_REGION            = os.Getenv("AWS_REGION")
	AWS_ACCESS_KEY_ID     = os.Getenv("AWS_ACCESS_KEY_ID")
	AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")

	AWS_FROM = os.Getenv("AWS_FROM")
	AWS_TO1  = os.Getenv("AWS_TO1")
	AWS_TO2  = os.Getenv("AWS_TO2")
	AWS_MSG  = ""
)

func main() {
	c := sesv2.NewClient(
		AWS_REGION,
		AWS_ACCESS_KEY_ID,
		AWS_SECRET_ACCESS_KEY,
	)

	htmlBody := `<h1>Hello World</h1>`

	data1, err := os.ReadFile("file1.pdf")
	if err != nil {
		log.Fatal(err)
	}

	data2, err := os.ReadFile("file2.pdf")
	if err != nil {
		log.Fatal(err)
	}

	req := sesv2.DataEmail{
		ToAddresses:  []string{AWS_TO1},
		From:         AWS_FROM,
		FromMsg:      "message",
		Title:        "title",
		MsgHTML:      htmlBody,
		BccAddresses: []string{AWS_TO1, AWS_TO2},
		CcAddresses:  []string{AWS_TO1},
		Attachments: []sesv2.Attachment{
			{Data: data1, Name: "file1.pdf"},
			{Data: data2, Name: "file2.pdf"},
		},
	}

	if err := c.Send(context.Background(), req); err != nil {
		log.Fatal(err)
	}
}
