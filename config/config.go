package config

import "github.com/spf13/viper"

type Config struct {
	Database DatabaseConfig `mapstructure:"DATABASE"`
	NSQ      NSQConfig      `mapstructure:"NSQ"`
	Pusher   PusherConfig   `mapstructure:"PUSHER"`
}
type DatabaseConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
	Name     string `mapstructure:"NAME"`
}
type NSQConfig struct {
	Host   string `mapstructure:"HOST"`
	Port   string `mapstructure:"PORT"`
	Topic  string `mapstructure:"TOPIC"`
	Topic2 string `mapstructure:"TOPIC2"`
}
type PusherConfig struct {
	AppId   string `mapstructure:"APPID"`
	Key     string `mapstructure:"KEY"`
	Secret  string `mapstructure:"SECRET"`
	Cluster string `mapstructure:"CLUSTER"`
	Secure  bool   `mapstructure:"SECURE"`
	Channel string `mapstructure:"CHANNEL"`
	Event1  string `mapstructure:"EVENT1"`
	Event2  string `mapstructure:"EVENT2"`
	Event3  string `mapstructure:"EVENT3"`
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
