package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	CORS     CORSConfig
}

type ServerConfig struct {
	Host         string
	Port         string
	RateLimit    int
	RateInterval int
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type CORSConfig struct {
	AllowedOrigins []string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// Read .env file if it exists
	_ = viper.ReadInConfig()

	// Server config
	viper.SetDefault("SERVER_HOST", "0.0.0.0")
	viper.SetDefault("SERVER_PORT", ":9999")
	viper.SetDefault("RATE_LIMIT", 60)
	viper.SetDefault("RATE_INTERVAL", 60)

	// Database config
	viper.SetDefault("DB_HOST", "127.0.0.1")
	viper.SetDefault("DB_PORT", "3307")
	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASSWORD", "root")
	viper.SetDefault("DB_NAME", "acheisuacara")

	// Redis config
	viper.SetDefault("REDIS_HOST", "localhost")
	viper.SetDefault("REDIS_PORT", "6379")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("REDIS_DB", 0)

	// CORS config
	viper.SetDefault("ALLOWED_ORIGINS", "http://localhost:5173,https://acheisuacara.com.br")

	config := &Config{
		Server: ServerConfig{
			Host:         viper.GetString("SERVER_HOST"),
			Port:         viper.GetString("SERVER_PORT"),
			RateLimit:    viper.GetInt("RATE_LIMIT"),
			RateInterval: viper.GetInt("RATE_INTERVAL"),
		},
		Database: DatabaseConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			DBName:   viper.GetString("DB_NAME"),
		},
		Redis: RedisConfig{
			Host:     viper.GetString("REDIS_HOST"),
			Port:     viper.GetString("REDIS_PORT"),
			Password: viper.GetString("REDIS_PASSWORD"),
			DB:       viper.GetInt("REDIS_DB"),
		},
		CORS: CORSConfig{
			AllowedOrigins: strings.Split(viper.GetString("ALLOWED_ORIGINS"), ","),
		},
	}

	return config, nil
}
