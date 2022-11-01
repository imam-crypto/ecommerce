package utils

import (
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/spf13/viper"
)

type Config struct {
	GinMode                  string          `mapstructure:"GIN_MODE"`
	DbHost                   string          `mapstructure:"DB_HOST"`
	DbPort                   string          `mapstructure:"DB_PORT"`
	DbName                   string          `mapstructure:"DB_NAME"`
	DbUsername               string          `mapstructure:"DB_USERNAME"`
	DbPassword               string          `mapstructure:"DB_PASSWORD"`
	DbTz                     string          `mapstructure:"DB_TZ"`
	GormMigrate              string          `mapstructure:"GORM_MIGRATE"`
	SecretKey                string          `mapstructure:"SECRET_KEY"`
	CloudName                string          `mapstructure:"CLOUDINARY_CLOUD_NAME"`
	CloudApiKey              string          `mapstructure:"CLOUDINARY_API_KEY"`
	CloudApiSecret           string          `mapstructure:"CLOUDINARY_API_SECRET"`
	CloudUploadCategoryImage string          `mapstructure:"CLOUDINARY_UPLOAD_CATEGORY_IMAGE"`
	CloudUploadProductImage  string          `mapstructure:"CLOUDINARY_UPLOAD_PRODUCT_IMAGE"`
	CloudAllowTypeImage      api.CldAPIArray `mapstructure:"CLOUDINARY_ALLOW_FORMAT_IMAGE"`
	CloudFolderCategory      string          `mapstructure:"CLOUDINARY_UPLOAD_CATEGORY_IMAGE"`
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
