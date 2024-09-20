package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"helpers/internal/api"
	server "helpers/pkg"
	logutil "helpers/pkg/logutils"
	"helpers/sharedconfigs"
	"os"
	"runtime/debug"
)

func main() {
	logutil.InitLogger()

	defer func() {
		if r := recover(); r != nil {
			logutil.Logger().Errorf("Unexpected panic: %v,\n%s", r, debug.Stack())
			os.Exit(1)
		}
	}()

	godotenv.Load()

	FileServerConfig := sharedconfigs.ServerConfiguration{}

	s := api.FileCollabServer{}
	s.ServerConfig = FileServerConfig
	fmt.Printf("", s.GetBasePath(), s.GetHost())
	server.RunServer(s)
}
