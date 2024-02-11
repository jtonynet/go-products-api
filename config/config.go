package config

import "github.com/spf13/viper"

type API struct {
	// API General
	Name       string `mapstructure:"API_NAME"`
	Port       string `mapstructure:"API_PORT"`
	TagVersion string `mapstructure:"API_TAG_VERSION"`
	Env        string `mapstructure:"API_ENV"`
	Host       string `mapstructure:"API_HOST"`

	// API Metrics to Prometheus
	MetricsEnabled              bool `mapstructure:"API_METRICS_ENABLED"`
	MetricsRefreshIntervalInSec int  `mapstructure:"API_METRICS_REFRESH_INTERVAL_IN_SEC"`
}

type Database struct {
	//DataBase Conn General
	Host                    string `mapstructure:"DATABASE_HOST"`
	Port                    string `mapstructure:"DATABASE_PORT"`
	User                    string `mapstructure:"DATABASE_USER"`
	Pass                    string `mapstructure:"DATABASE_PASSWORD"`
	DB                      string `mapstructure:"DATABASE_DB"`
	RetryMaxElapsedTimeInMs int    `mapstructure:"DATABASE_RETRY_MAX_ELAPSED_TIME_IN_MS"`

	//DataBase Conn Metrics to Prometheus
	MetricsEnabled              bool   `mapstructure:"DATABASE_METRICS_ENABLED"`
	MetricsDBName               string `mapstructure:"DATABASE_METRICS_NAME"`
	MetricsRefreshIntervalInSec uint32 `mapstructure:"DATABASE_METRICS_REFRESH_INTERVAL_IN_SEC"`
	MetricStartServer           bool   `mapstructure:"DATABASE_METRICS_START_SERVER"`
	MetricServerPort            uint32 `mapstructure:"DATABASE_METRICS_SERVER_PORT"`
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
