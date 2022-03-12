package custom_configcat

import (
	"configcat-homework/internal/custom_viper"
	"errors"
	"fmt"

	configcat "github.com/configcat/go-sdk"
)

type Client struct {}

type Connection struct {
	Connection *configcat.Client
}

func (Client) Connect() Connection {

	appViper := custom_viper.Set{}.Viper().CustomViper

	// Get configcat API key from .env file
	configcatAPIKey, ok := appViper.Get("CONFIGCAT_APIKEY").(string)
	if !ok {
		err := errors.New("CONFIGCAT_APIKEY not found")
		fmt.Println(err)
	}
	
	return Connection {
		Connection: configcat.NewClient(configcatAPIKey), 
	}
	

}

