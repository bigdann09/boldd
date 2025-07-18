package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	ApplicationConfig `mapstructure:"application"`
	DatabaseConfig    `mapstructure:"database"`
	RedisConfig       `mapstructure:"redis"`
	CorsConfig        `mapstructure:"cors"`
	AWSConfig         `mapstructure:"aws"`
	JSWConfig         `mapstructure:"jwt"`
	MailConfig        `mapstructure:"mail"`
}

type ApplicationConfig struct {
	Port        int    `mapstructure:"port"`
	URL         string `mapstructure:"url"`
	Timezone    string `mapstructure:"timezone"`
	Environment string `mapstructure:"environment"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	SSLMode  string `mapstructure:"sslmode"`
}

type CorsConfig struct {
	AllowedOrigins []string `mapstructure:"allowed_origins"`
}

type RedisConfig struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	Protocol int    `mapstructure:"protocol"`
}

type AWSConfig struct {
	Credential string
}

type JSWConfig struct {
	Key           string `mapstructure:"key"`
	AccessExpiry  int    `mapstructure:"access_expiry"`
	RefreshExpiry int    `mapstructure:"refresh_expiry"`
}

type MailConfig struct {
	From     string `mapstructure:"from"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

func Load(path string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Config file not found at %s: %v", viper.ConfigFileUsed(), err)
		} else {
			log.Fatalf("Error reading config file: %v", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to decode into struct %v", err)
	}

	return &cfg, nil
}

func LoadConfigPath() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		return "", errors.New("development environment not provided")
	}

	var path string
	switch env {
	case "docker_development":
		path = "/app/boldd"
	case "development":
		path = "$HOME/.config/boldd/"
	case "production":
		path = "/etc/boldd"
	default:
		path = "$HOME/.config/boldd/"
	}

	return path, nil
}
