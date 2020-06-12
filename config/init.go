package config

import (
	"log"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

var instance *Connection

type Connection struct {
	DBLive driver.Client
}

func init() {
	if instance != nil {
		return
	}

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		log.Printf("Error open connection, cause:%+v\n", err)
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "password"),
	})
	if err != nil {
		log.Printf("Error create a new client connection, cause:%+v\n", err)
	}

	instance = &Connection{
		DBLive: client,
	}
}

// GetInstance return a connection instance
func GetInstance() *Connection {
	return instance
}
