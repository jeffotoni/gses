// @autor: @jeffotoni
package v2

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	ses "github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

type Client struct {
	ses    *ses.Client
	region string
	key    string
	secret string
}

func NewClient(region, key, secret string) *Client {
	return &Client{
		ses:    &ses.Client{},
		region: region,
		key:    key,
		secret: secret,
	}
}

func (c *Client) Send(ctx context.Context, data DataEmail) error {
	if err := data.Validate(); err != nil {
		return err
	}

	destination := types.Destination{
		ToAddresses:  data.ToAddresses,
		BccAddresses: data.BccAddresses,
		CcAddresses:  data.CcAddresses,
	}

	var emailBody bytes.Buffer

	header := fmt.Sprintf(
		"From: %s\nTo: %s\nBcc: %s\nCc: %s\nSubject: %s\nMIME-Version: 1.0\nContent-Type: multipart/mixed; boundary=\"%s\"\n\n",
		data.From,
		strings.Join(data.ToAddresses, ","),
		strings.Join(data.BccAddresses, ","),
		strings.Join(data.CcAddresses, ","),
		data.Title,
		"--_GoBoundary",
	)
	emailBody.WriteString(header)

	emailBody.WriteString("----_GoBoundary\n")
	emailBody.WriteString("Content-Type: text/html; charset=UTF-8\n")
	emailBody.WriteString("Content-Transfer-Encoding: base64\n\n")
	emailBody.WriteString(base64.StdEncoding.EncodeToString([]byte(data.MsgHTML)))
	emailBody.WriteString("\n")

	for _, v := range data.Attachments {
		emailBody.WriteString("----_GoBoundary\n")
		emailBody.WriteString(fmt.Sprintf("Content-Type: text/csv; name=\"%s\"\n", v.Name)) //fname[len(fname)-1]))
		emailBody.WriteString("Content-Description: file\n")
		emailBody.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"; size=%d;\n", v.Name, len(v.Name))) //fname[len(fname)-1], len(fdata)))
		emailBody.WriteString("Content-Transfer-Encoding: base64\n\n")
		emailBody.WriteString(base64.StdEncoding.EncodeToString(v.Data))
		emailBody.WriteString("\n")
	}

	emailBody.WriteString("----_GoBoundary--\n")

	params := &ses.SendEmailInput{
		Destination: &destination,
		Content: &types.EmailContent{
			Raw: &types.RawMessage{Data: emailBody.Bytes()},
		},
	}

	svc := ses.NewFromConfig(aws.Config{
		Credentials: credentials.NewStaticCredentialsProvider(
			c.key,
			c.secret,
			"",
		),
		Region: c.region,
	})
	if svc == nil {
		return ErrNilSVC

	}

	_, err := svc.SendEmail(ctx, params)
	return err
}
