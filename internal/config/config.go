package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ApplicationConfig `mapstructure:"application"`
	DatabaseConfig    `mapstructure:"database"`
	RedisConfig       `mapstructure:"redis"`
	CorsConfig        `mapstructure:"cors"`
	AWSConfig         `mapstructure:"aws"`
}

type ApplicationConfig struct {
	Port        int    `mapstructure:"port"`
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
