package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const boundary = "--_GoBoundary"

func getMIMEEmail(from, to, subject, htmlBody string, attachmentPaths []string) (string, error) {
	var emailBody bytes.Buffer

	// Headers
	header := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\nMIME-Version: 1.0\nContent-Type: multipart/mixed; boundary=\"%s\"\n\n", from, to, subject, boundary)
	emailBody.WriteString(header)

	// HTML Body
	emailBody.WriteString(fmt.Sprintf("--%s\n", boundary))
	emailBody.WriteString("Content-Type: text/html; charset=UTF-8\n")
	emailBody.WriteString("Content-Transfer-Encoding: base64\n\n")
	emailBody.WriteString(base64.StdEncoding.EncodeToString([]byte(htmlBody)))
	emailBody.WriteString("\n")

	// Attachments
	for _, filePath := range attachmentPaths {
		fileName := strings.Split(filePath, "/")
		fileBytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			return "", err
		}

		emailBody.WriteString(fmt.Sprintf("--%s\n", boundary))
		emailBody.WriteString(fmt.Sprintf("Content-Type: application/octet-stream; name=\"%s\"\n", fileName[len(fileName)-1]))
		emailBody.WriteString("Content-Description: file\n")
		emailBody.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"; size=%d;\n", fileName[len(fileName)-1], len(fileBytes)))
		emailBody.WriteString("Content-Transfer-Encoding: base64\n\n")
		emailBody.WriteString(base64.StdEncoding.EncodeToString(fileBytes))
		emailBody.WriteString("\n")
	}

	// Final boundary
	emailBody.WriteString(fmt.Sprintf("--%s--\n", boundary))

	return emailBody.String(), nil
}

func main() {
	// Suponha que você já configurou sua sessão e cliente SES...
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("SES_REGION")),
	})
	if err != nil {
		fmt.Println("Erro ao criar a sessão AWS:", err)
		return
	}

	svc := ses.New(sess)

	from := os.Getenv("SES_EMAIL_FROM")
	to := os.Getenv("SES_EMAIL_TO")
	subject := "Assunto do email TESTE TDC-V2"
	htmlBody := "<h1>Email com Anexo</h1><p>Olá!</p>"
	attachments := []string{"./desenho1.pdf", "./desenho2.pdf"}

	mimeEmail, err := getMIMEEmail(from, to, subject, htmlBody, attachments)
	if err != nil {
		fmt.Println("Erro ao criar o email:", err)
		return
	}

	input := &ses.SendRawEmailInput{
		RawMessage: &ses.RawMessage{
			Data: []byte(mimeEmail),
		},
	}

	_, err = svc.SendRawEmail(input)
	if err != nil {
		fmt.Println("Erro ao enviar o email:", err)
		return
	}

	fmt.Println("Email enviado com sucesso!")
}
