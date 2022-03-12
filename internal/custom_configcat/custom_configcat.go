package custom_configcat

import (
	configcat "github.com/configcat/go-sdk"
)

type Client struct {}

type Connection struct {
	Connection *configcat.Client
}

func (Client) Connect(APIKey string) Connection {
	
	return Connection {
		Connection: configcat.NewClient(APIKey), 
	}
	

}

