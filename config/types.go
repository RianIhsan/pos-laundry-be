package config

import "time"

type Config struct {
	Server      ServiceConfig
	Postgres    PostgresConfig
	Logger      LoggerConfig
	RateLimiter RateLimiterConfig
}

type ServiceConfig struct {
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Mode         string
	SSL          bool
	JWTSecretKey string
	EncryptKey   string
}

type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Dbname   string
	SSLMode  string
}

type LoggerConfig struct {
	Level       string
	Caller      bool
	Encoding    string
	Development bool
}

type RateLimiterConfig struct {
	Enabled bool // Global enable/disable flag
	API     RateLimitConfig
}

type RateLimitConfig struct {
	RequestsPerMinute int
	Burst             int
}
