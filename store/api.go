package store

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetStore(ConfigPath string) *Store {
	s := new(Store)
	_, err := s.setConfig(ConfigPath)
	if err != nil {
		panic(err)
	}

	s.router = mux.NewRouter()
	s.setDatabase()

	return s
}

func RunServer(s *Store) {
	address, err := s.GetAddress()
	if err != nil {
		panic(err)
	}

	fmt.Printf("server is starting on %s\n", address)

	srv := &http.Server{
		Handler: s.router,
		Addr:    address,
	}

	log.Fatal(srv.ListenAndServe())
}
