package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type StateConfig struct {
	State string `mapstructure:"STATE"`
}

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	State                string        `mapstructure:"STATE"`
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	MigrationURL         string        `mapstructure:"MIGRATION_URL"`
	RedisAddress         string        `mapstructure:"REDIS_ADDRESS"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	RedisPort            string        `mapstructure:"REDIS_PORT"`
	RedisHost            string        `mapstructure:"REDIS_HOST"`
	RedisPassword        string        `mapstructure:"REDIS_PASSWORD"`
	RedisDatabase        int           `mapstructure:"REDIS_DATABASE"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	SupaUrl              string        `mapstructure:"SUPA_URL"`
	SupaKey              string        `mapstructure:"SUPA_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	ClientBaseUrl        string        `mapstructure:"CLIENT_BASE_URL"`
	ResendApiKey         string        `mapstructure:"RESEND_API_KEY"`
}

func loadAncScan(config *Config) (err error) {
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(config)
	if err != nil {
		return err
	}
	return nil
}

// LoadConfig reads configuration from file or environment variables.
func LoadState(path string) (config StateConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("state")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string, state string) (config Config, err error) {
	stateEnvFilePath := fmt.Sprintf("%s.env", state)
	viper.SetConfigName(stateEnvFilePath)
	err = loadAncScan(&config)
	if err != nil {
		return
	}
	viper.SetConfigName("shared.env")
	err = loadAncScan(&config)
	if err != nil {
		return
	}
	return
}
