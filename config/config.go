package config

import "github.com/spf13/viper"

type Config struct {
	Database DatabaseConfig `mapstructure:"DATABASE"`
	NSQ      NSQConfig      `mapstructure:"NSQ"`
}
type DatabaseConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
	Name     string `mapstructure:"NAME"`
}
type NSQConfig struct {
	Host  string `mapstructure:"HOST"`
	Port  string `mapstructure:"PORT"`
	Topic string `mapstructure:"TOPIC"`
}

func InitConfiguration() (*Config, error) {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	viper.AutomaticEnv()
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
