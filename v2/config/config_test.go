package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

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

func TestFromFile(t *testing.T) {

	expectedConfig := Config{
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

	t.Run("Pass load config", func(t *testing.T) {
		err := createEnvContent(pathWorking, content)
		require.NoError(t, err)

		defer delFile()

		cfg, err := FromFile(pathWorking)
		require.NoError(t, err)

		if *cfg != expectedConfig {
			t.Errorf("expected config %v, got %v", expectedConfig, *cfg)
		}
	})

	t.Run("Error load config", func(t *testing.T) {

		expectedConfig.AwsFrom = "fake from cause forced error"

		err := createEnvContent(pathWorking, content)
		require.NoError(t, err)

		defer delFile()

		cfg, err := FromFile(pathWorking)
		require.NoError(t, err)

		if *cfg == expectedConfig {
			t.Errorf("expected config %v, got %v", expectedConfig, *cfg)
		}
	})

	t.Run("Error laod config non exist file", func(t *testing.T) {
		fileNoExist := "./no_exist"

		cfg, err := FromFile(fileNoExist)
		require.Error(t, err, errors.New("no such file or directory"))
		require.Nil(t, cfg)
	})
}
