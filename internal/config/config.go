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
	CorsConfig        `mapstructure:"cors"`
	AWSConfig         `mapstructure:"aws"`
	JWTConfig         `mapstructure:"jwt"`
	MailConfig        `mapstructure:"mail"`
	RedisConfig
	GoogleOAuthConfig
	DatabaseConfig
	CloudinaryConfig
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
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type AWSConfig struct {
	Credential string
}

type JWTConfig struct {
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

type GoogleOAuthConfig struct {
	ClientID     string
	ClientSecret string
	CallbackURL  string
}

type CloudinaryConfig struct {
	CloudName string
	Key       string
	Secret    string
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

	log.Println("Load application core configs")
	cfg.LoadMailConfig()
	cfg.LoadDatabaseConfig()
	cfg.LoadRedisConfig()
	cfg.LoadGoogleConfig()
	cfg.LoadCloudinaryConfig()

	return &cfg, nil
}

func LoadConfigPath() (string, error) {
	if os.Getenv("DOCKER_ENV") != "true" {
		if err := godotenv.Load(); err != nil {
			return "", err
		}
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

func (cfg *Config) LoadGoogleConfig() {
	cfg.GoogleOAuthConfig = GoogleOAuthConfig{
		ClientID:     viper.GetString("GOOGLE_CLIENT_ID"),
		ClientSecret: viper.GetString("GOOGLE_CLIENT_SECRET"),
		CallbackURL:  viper.GetString("GOOGLE_CALLBACK_URI"),
	}
}

func (cfg *Config) LoadDatabaseConfig() {
	cfg.DatabaseConfig = DatabaseConfig{
		Port:     viper.GetInt("DATABASE_PORT"),
		Host:     viper.GetString("DATABASE_HOST"),
		Database: viper.GetString("DATABASE_NAME"),
		Username: viper.GetString("DATABASE_USER"),
		Password: viper.GetString("DATABASE_PASS"),
		SSLMode:  viper.GetString("DATABASE_SSLMODE"),
	}
}

func (cfg *Config) LoadRedisConfig() {
	cfg.RedisConfig = RedisConfig{
		DB:       viper.GetInt("REDIS_DB"),
		Host:     viper.GetString("REDIS_HOST"),
		Password: viper.GetString("REDIS_PASSWORD"),
		Port:     viper.GetString("REDIS_PORT"),
	}
}

func (cfg *Config) LoadCloudinaryConfig() {
	cfg.CloudinaryConfig = CloudinaryConfig{
		CloudName: viper.GetString("CLOUDINARY_CLOUD_NAME"),
		Key:       viper.GetString("CLOUDINARY_API_KEY"),
		Secret:    viper.GetString("CLOUDINARY_API_SECRET"),
	}
}

func (cfg *Config) LoadMailConfig() {
	cfg.MailConfig = MailConfig{
		Host:     viper.GetString("MAIL_HOST"),
		Port:     viper.GetInt("MAIL_PORT"),
		From:     viper.GetString("MAIL_FROM"),
		Username: viper.GetString("MAIL_USERNAME"),
		Password: viper.GetString("MAIL_PASSWORD"),
	}
}
