package main

import (
	"fmt"
	"go-serv/store"

	"net/http"
)

func main() {
	s := store.SetStore("etc/config.yaml")

	s.POST("/", func(writer http.ResponseWriter, request *http.Request, s *store.Store) {
		fmt.Fprintf(writer, "Hello alone")
		data, err := store.ParseJSON(request)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s", data)
	})

	store.RunServer(s)
}
