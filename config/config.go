package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server  ServerConfig  `viper:"server"`
	DB      DBConfig      `viper:"db"`
	Redis   RedisConfig   `viper:"redis"`
	Session SessionConfig `viper:"session"`
	Cookie  CookieConfig  `viper:"cookie"`
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
	Debug     bool
}

type RedisConfig struct {
	Addr         string
	Password     string
	DB           int
	DefaultDB    string
	MinIdleConns int
	PoolSize     int
	PoolTimeout  int
}

type SessionConfig struct {
	Prefix string
	Name   string
	Expire int
}

type CookieConfig struct {
	Name     string
	MaxAge   int
	Path     string
	Domain   string
	Secure   bool
	HTTPOnly bool
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
