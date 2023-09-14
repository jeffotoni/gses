package main

import (
	"log"

	"github.com/jeffotoni/gses/config"
	"github.com/jeffotoni/gses/models"
	"github.com/jeffotoni/gses/ses"
)

func main() {

	cfg, err := config.FromFile(".")
	if err != nil {
		log.Fatal(err)
	}

	ses := ses.NewSesEmail(cfg)

	defaultProfile := "default"
	// profile0 := "profile0"

	profile, err := ses.AddProfile(
		defaultProfile,
		cfg.AwsRegion,
		cfg.AwsIdentity,
		cfg.AwsFrom,
		cfg.AwsInfo,
	)

	if err != nil {
		log.Fatal(err)
	}

	_ = profile

	// ses.AddProfile(
	// 	profile0,
	// 	cfg.AwsRegion,
	// 	cfg.AwsIdentity,
	// 	cfg.AwsFrom,
	// 	cfg.AwsInfo,
	// )

	htmlBody := `<h1>Hello World</h1>`

	data := models.NewDataEmail(
		"to",
		"from",
		"message",
		"title",
		htmlBody,
		"bccAddress",
		"ccAddress",
	)

	if err := ses.SendEmailSes(defaultProfile, data); err != nil {
		log.Fatal(err)
	}

	// if err := ses.SendEmailSes(profile0, data); err != nil {
	// 	log.Fatal(err)
	// }
}
