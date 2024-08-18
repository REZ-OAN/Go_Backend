package utils

import "github.com/spf13/viper"

type Config struct {
	DB_DRIVER      string `mapstructure:"DB_DRIVER"`
	DB_SOURCE      string `mapstructure:"DB_SOURCE"`
	DB_NAME        string `mapstructure:"DB_NAME"`
	DB_USER        string `mapstructure:"USER"`
	DB_PASSWORD    string `mapstructure:"PASS"`
	DB_PORT        string `mapstructure:"PORT"`
	SERVER_ADDRESS string `mapstructure:"SERVER_ADDRESS"`
	SERVER_PORT    string `mapstructure:"SERVER_PORT"`
}

func LoadConfig(path string, name string, ext string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(ext)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}
