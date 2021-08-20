package main

import (
	"fmt"
	"go-serv/store"
	"net/http"
)

func main() {
	s := store.SetStore("etc/config.yaml")

	s.AddRoute("/", func(writer http.ResponseWriter, request *http.Request, s *store.Store) {
		fmt.Fprintf(writer, "Hello alone")
	})

	store.RunServer(s)
}
