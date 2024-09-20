package sharedconfigs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var DEFAULT_HOST = "0.0.0.0"
var DEFAULT_PORT = 8080

var HOST string

func init() {
	godotenv.Load()
	port := DEFAULT_PORT
	if sp, ok := os.LookupEnv("SERVER_PORT"); ok {
		if p, err := strconv.Atoi(sp); err == nil {
			port = p
		}
	}
	HOST = fmt.Sprintf("%s:%d", DEFAULT_HOST, port)
}

type ServerConfiguration struct {
	Host     string
	BasePath string
}

type IServerConfiguration interface {
	GetHost() string
	GetBasePath() string
}

func (c ServerConfiguration) GetHost() string {
	if c.Host != "" && len(c.Host) != 0 {
		return c.Host
	}
	return HOST
}

func (c ServerConfiguration) GetBasePath() string {
	if c.BasePath != "" && len(c.BasePath) != 0 {
		return c.BasePath
	}
	return ""
}
