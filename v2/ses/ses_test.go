package ses

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	config "github.com/jeffotoni/gses/v2/config"
	"github.com/jeffotoni/gses/v2/models"
	"github.com/stretchr/testify/require"
)

var (
	envName     = "/.env"
	pathWorking = "./"
)

func createEnvContent(path string, content string) error {
	return os.WriteFile(path+envName, []byte(strings.TrimSpace(content)), 0644)
}

func delFile() error {
	return os.Remove(pathWorking + envName)
}

func TestSendEmail(t *testing.T) {

	expectedConfig := config.Config{
		AwsRegion:    "region",
		AwsIdentity:  "identity",
		AwsAccess:    "access",
		AwsSecret:    "secret",
		AwsFrom:      "from",
		AwsMessage:   "message",
		AwsInfo:      "info",
		SendInterval: time.Duration(10) * time.Second,
	}

	content := fmt.Sprintf(`
	AWS_REGION="%s"
AWS_IDENTITY="%s"
AWS_ACCESS_KEY_ID="%s"
AWS_SECRET_ACCESS_KEY="%s"
AWS_FROM="%s"
AWS_MSG="%s"
AWS_INFO="%s"
SEND_INTERVAL=%v`,
		expectedConfig.AwsRegion,
		expectedConfig.AwsIdentity,
		expectedConfig.AwsAccess,
		expectedConfig.AwsSecret,
		expectedConfig.AwsFrom,
		expectedConfig.AwsMessage,
		expectedConfig.AwsInfo,
		expectedConfig.SendInterval,
	)

	err := createEnvContent(pathWorking, content)
	require.NoError(t, err)

	defer delFile()

	cfg, err := config.FromFile(pathWorking)
	require.NoError(t, err)
	require.NotNil(t, cfg)

	defaultProfile := "default"
	fakeProfileNotExist := "fake-profile"

	ses := NewSesEmail(cfg)

	t.Run("PASS send mail", func(t *testing.T) {

		prof, err := ses.AddProfile(defaultProfile, cfg.AwsRegion, cfg.AwsIdentity, cfg.AwsFrom, cfg.AwsInfo)
		require.NoError(t, err)
		require.NotNil(t, prof)

		data := models.NewDataEmail("to", "from", "message", "title", "htmlMsgBody", "bccAddress", "ccAddress")

		err = ses.SendEmailSes(defaultProfile, data)
		require.NoError(t, err)
	})

	t.Run("FAIL no Profile Set", func(t *testing.T) {
		//No set profile

		// ses.AddProfile(defaultProfile, config.AwsRegion, config.AwsIdentity, config.AwsFrom, config.AwsInfo)

		data := models.NewDataEmail("to", "from", "message", "title", "htmlMsgBody", "bccAddress", "ccAddress")

		err := ses.SendEmailSes(fakeProfileNotExist, data)
		require.Error(t, err, ErrNoProfileSet)
	})

	t.Run("FAIL profile not searched", func(t *testing.T) {

		prof, err := ses.AddProfile(defaultProfile, cfg.AwsRegion, cfg.AwsIdentity, cfg.AwsFrom, cfg.AwsInfo)
		require.NoError(t, err)
		require.NotNil(t, prof)

		data := models.NewDataEmail("to", "from", "message", "title", "htmlMsgBody", "bccAddress", "ccAddress")

		//Profile param in function SendEmailSes no added
		err = ses.SendEmailSes(fakeProfileNotExist, data)
		require.Error(t, err, ErrProfileNotSearched)
	})
}
