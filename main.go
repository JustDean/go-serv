package main

import (
	"go-serv/store"
	"go-serv/test-app"
)

func main() {
	s := store.SetStore("etc/config.yaml")

	test_app.AddPaths(s)

	store.RunServer(s)
}
