package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AwsRegion    string        `mapstructure:"AWS_REGION"`
	AwsIdentity  string        `mapstructure:"AWS_IDENTITY"`
	AwsAccess    string        `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecret    string        `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AwsFrom      string        `mapstructure:"AWS_FROM"`
	AwsMessage   string        `mapstructure:"AWS_MSG"`
	AwsInfo      string        `mapstructure:"AWS_INFO"`
	SendInterval time.Duration `mapstructure:"SEND_INTERVAL"`
}

func FromFile(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("viper.ReadInConfig: %v", err)
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("viper.Unmarshal: %v", err)
	}

	return &c, nil
}

func FromEnv() *Config {
	return nil
}
