package main

import (
	"go-serv/store"
)

func main() {
	s := store.SetStore("etc/config.yaml")

	store.RunServer(s)
}
