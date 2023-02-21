package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server ServerConfig `viper:"server"`
	DB     DBConfig     `viper:"db"`
}

type ServerConfig struct {
	AppVersion   string
	Port         string
	JwtSecretKey string
	SSL          bool
	CSRF         bool
	Debug        bool
}

type DBConfig struct {
	DBName    string
	User      string
	Password  string
	Addr      string
	Net       string
	ParseTime bool
	Collation string
	Location  string
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
