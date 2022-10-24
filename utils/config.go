package utils

import "github.com/spf13/viper"

type Config struct {
	GinMode     string `mapstructure:"GIN_MODE"`
	DbHost      string `mapstructure:"DB_HOST"`
	DbPort      string `mapstructure:"DB_PORT"`
	DbName      string `mapstructure:"DB_NAME"`
	DbUsername  string `mapstructure:"DB_USERNAME"`
	DbPassword  string `mapstructure:"DB_PASSWORD"`
	DbTz        string `mapstructure:"DB_TZ"`
	GormMigrate string `mapstructure:"GORM_MIGRATE"`
	SecretKey   string `mapstructure:"SECRET_KEY"`
}

func LoadConfig(path string, testing bool) (config Config, err error) {
	if testing {
		viper.AddConfigPath(path)
		viper.SetConfigName("testing")
		viper.SetConfigType("env")
	} else {
		viper.AddConfigPath(path)
		viper.SetConfigName("app")
		viper.SetConfigType("env")
	}

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
