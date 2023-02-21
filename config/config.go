package config

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	AppVersion string
	Port       string
	SSL        bool
	CSRF       bool
	Debug      bool
}

func LoadConfig(filename string) (*viper.Viper, error) {

}
