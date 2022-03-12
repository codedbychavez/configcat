package custom_viper

import (
	"fmt"
	"github.com/spf13/viper"
)

type Set struct {}

type CustomViper struct {
	CustomViper *viper.Viper
}


func (Set) Viper() CustomViper {
	// Setup
	newViper := viper.New()
	newViper.AutomaticEnv()
	newViper.SetConfigName(".env")
	newViper.SetConfigType("env")
	newViper.AddConfigPath("./")

	// Error checking for .env file
	if err := newViper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Error, could not locate .env file", err)
		} else {
			fmt.Println("Error loading .env file", err)
		}
	}

	return CustomViper{
		CustomViper: newViper,
	}
}