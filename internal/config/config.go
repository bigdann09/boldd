package config

type Config struct {
	Application ApplicationConfig `mapstructure:"application"`
	Database    DatabaseConfig    `mapstructure:"database"`
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

func Load() (*Config, error) {
	return &Config{}, nil
}
