// @autor: @jeffotoni
package v2

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const boundary = "--_GoBoundary"

type Client struct {
	region string
	key    string
	secret string
}

func NewClient(region, key, secret string) *Client {
	return &Client{
		region: region,
		key:    key,
		secret: secret,
	}
}

func (c *Client) Send(ctx context.Context, data DataEmail) error {
	if err := data.Validate(); err != nil {
		return err
	}

	mime, err := getMIMEEmail(data)
	if err != nil {
		return err
	}

	params := &ses.SendRawEmailInput{
		RawMessage: &ses.RawMessage{
			Data: []byte(mime),
		},
	}

	s, err := c.getSession()
	if err != nil {
		return err
	}

	_, err = s.SendRawEmail(params)
	return err
}

func (c *Client) getSession() (*ses.SES, error) {
	var sess *session.Session
	var err error
	if len(c.key) > 0 && len(c.secret) > 0 {
		sess, err = session.NewSession(&aws.Config{
			HTTPClient: &http.Client{
				Transport: &http.Transport{
					DisableKeepAlives: true,
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			},
			DisableSSL:  aws.Bool(true),
			Credentials: credentials.NewStaticCredentials(c.key, c.secret, ""),
			Region:      aws.String(c.region),
		})
	} else {
		sess, err = session.NewSession(&aws.Config{
			Region: aws.String(c.region),
		})
	}
	if err != nil {
		return nil, err
	}

	return ses.New(sess), nil
}

func getMIMEEmail(data DataEmail) (string, error) {
	var emailBody bytes.Buffer

	// Headers
	header := fmt.Sprintf(
		"From: %s\nTo: %s\nBcc: %s\nCc: %s\nSubject: %s\nMIME-Version: 1.0\nContent-Type: multipart/mixed; boundary=\"%s\"\n\n",
		data.From,
		strings.Join(data.ToAddresses, ","),
		strings.Join(data.BccAddresses, ","),
		strings.Join(data.CcAddresses, ","),
		data.Title,
		boundary,
	)
	emailBody.WriteString(header)

	// HTML Body
	emailBody.WriteString(fmt.Sprintf("--%s\n", boundary))
	emailBody.WriteString("Content-Type: text/html; charset=UTF-8\n")
	emailBody.WriteString("Content-Transfer-Encoding: base64\n\n")
	emailBody.WriteString(base64.StdEncoding.EncodeToString([]byte(data.MsgHTML)))
	emailBody.WriteString("\n")

	// Attachments
	for _, v := range data.Attachments {

		emailBody.WriteString(fmt.Sprintf("--%s\n", boundary))
		emailBody.WriteString(fmt.Sprintf("Content-Type: application/octet-stream; name=\"%s\"\n", v.Name))
		emailBody.WriteString("Content-Description: file\n")
		emailBody.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"; size=%d;\n", v.Name, len(v.Data)))
		emailBody.WriteString("Content-Transfer-Encoding: base64\n\n")
		emailBody.WriteString(base64.StdEncoding.EncodeToString(v.Data))
		emailBody.WriteString("\n")
	}

	// Final boundary
	emailBody.WriteString(fmt.Sprintf("--%s--\n", boundary))

	return emailBody.String(), nil
}
