package main

import (
	"go-serv/store"
	"net/http"
)

func main() {
	s := store.SetStore("etc/config.yaml")

	s.GET("/", func(writer http.ResponseWriter, request *http.Request, s *store.Store) {
		_, err := store.ParseJSON(request)
		if err != nil {
			panic(err)
		}

		err = store.ReturnJSON(writer, map[string]string{"hello": "alone"})
		if err != nil {
			panic(err)
		}
	})

	store.RunServer(s)
}
