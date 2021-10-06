package main

import (
	"go-serv/app"
	"go-serv/server"
	"os"
)

func main() {
	commandLineArguments := os.Args
	if len(commandLineArguments) > 1 {
		if commandLineArguments[1] == "create_test_data" {
			config := new(server.Config)
			config.ReadConfig("etc/config.yaml")
			server.ExampleDB_Model(config)
			os.Exit(1)
		}
	}

	s := server.SetServer("etc/config.yaml")

	s.AddMiddleware(app.LoggingMiddleware)

	app.RegisterUrls(s)

	server.RunServer(s)
}
