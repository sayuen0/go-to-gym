package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

// Config represents the configuration of whole application
type Config struct {
	Server  ServerConfig  `viper:"server"`
	DB      DBConfig      `viper:"db"`
	Redis   RedisConfig   `viper:"redis"`
	Session SessionConfig `viper:"session"`
	Cookie  CookieConfig  `viper:"cookie"`
}

// ServerConfig represents the configuration of application server
type ServerConfig struct {
	AppVersion   string
	Port         string
	JwtSecretKey string
	SSL          bool
	CSRF         bool
	Debug        bool
	Pepper       string
}

// DBConfig represents the configuration of DB connection
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

// RedisConfig represents the configuration of Redis connection
type RedisConfig struct {
	Addr         string
	Password     string
	DB           int
	DefaultDB    string
	MinIdleConns int
	PoolSize     int
	PoolTimeout  int
}

// SessionConfig represents the configuration of session
type SessionConfig struct {
	Prefix string
	Name   string
	Expire int
}

// CookieConfig represents the configuration of cookie
type CookieConfig struct {
	Name     string
	MaxAge   int
	Path     string
	Domain   string
	Secure   bool
	HTTPOnly bool
}

// LoadConfig reads file defined by filename and converts it to viper.Viper object
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

// ParseConfig parses viper.Viper into a Config object
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
