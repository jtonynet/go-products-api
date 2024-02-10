package config

import "github.com/spf13/viper"

type API struct {
	Name       string `mapstructure:"API_NAME"`
	Port       string `mapstructure:"API_PORT"`
	TagVersion string `mapstructure:"API_TAG_VERSION"`
	Env        string `mapstructure:"API_ENV"`
}

type Database struct {
	Host string `mapstructure:"DATABASE_HOST"`
	Port string `mapstructure:"DATABASE_PORT"`
	User string `mapstructure:"DATABASE_USER"`
	Pass string `mapstructure:"DATABASE_PASSWORD"`
	DB   string `mapstructure:"DATABASE_DB"`
}

type Config struct {
	API      API      `mapstructure:",squash"`
	Database Database `mapstructure:",squash"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
