package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

func NewAppConfig(configPath string) (*Config, error) {
	v := viper.New()

	v.AutomaticEnv()

	if _, err := os.Stat(".env"); err == nil {
		v.SetConfigFile(".env")
		v.SetConfigType("env")
		if err := v.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("failed to read .env file: %w", err)
		}
	}

	cfg := new(Config)

	// Server
	cfg.Server.Host = v.GetString("SERVER_HOST")
	cfg.Server.Port = v.GetInt("SERVER_PORT")
	cfg.Server.ReadTimeout = time.Duration(v.GetInt("SERVER_READ_TIMEOUT")) * time.Second
	cfg.Server.WriteTimeout = time.Duration(v.GetInt("SERVER_WRITE_TIMEOUT")) * time.Second
	cfg.Server.Mode = v.GetString("SERVER_MODE")
	cfg.Server.SSL = v.GetBool("SERVER_SSL")
	cfg.Server.JWTSecretKey = v.GetString("SERVER_JWT_SECRET_KEY")
	cfg.Server.EncryptKey = v.GetString("SERVER_ENCRYPT_KEY")

	// Postgres
	cfg.Postgres.User = v.GetString("POSTGRES_USER")
	cfg.Postgres.Password = v.GetString("POSTGRES_PASSWORD")
	cfg.Postgres.Host = v.GetString("POSTGRES_HOST")
	cfg.Postgres.Port = v.GetInt("POSTGRES_PORT")
	cfg.Postgres.Dbname = v.GetString("POSTGRES_DBNAME")
	cfg.Postgres.SSLMode = v.GetString("POSTGRES_SSL_MODE")

	// Logger
	cfg.Logger.Level = v.GetString("LOGGER_LEVEL")
	cfg.Logger.Caller = v.GetBool("LOGGER_CALLER")
	cfg.Logger.Encoding = v.GetString("LOGGER_ENCODING")
	cfg.Logger.Development = v.GetBool("LOGGER_DEVELOPMENT")

	cfg.RateLimiter.Enabled = getBoolWithDefault(v, "RATE_LIMITER_ENABLED", true)

	cfg.RateLimiter.API.RequestsPerMinute = getIntWithDefault(v, "RATE_LIMIT_API_REQUESTS_PER_MINUTE", 100)
	cfg.RateLimiter.API.Burst = getIntWithDefault(v, "RATE_LIMIT_API_BURST", 20)

	return cfg, nil
}

func getIntWithDefault(v *viper.Viper, key string, defaultValue int) int {
	if v.IsSet(key) {
		return v.GetInt(key)
	}
	return defaultValue
}

func getBoolWithDefault(v *viper.Viper, key string, defaultValue bool) bool {
	if v.IsSet(key) {
		return v.GetBool(key)
	}
	return defaultValue
}

func getStringWithDefault(v *viper.Viper, key string, defaultValue string) string {
	if v.IsSet(key) {
		return v.GetString(key)
	}
	return defaultValue
}
