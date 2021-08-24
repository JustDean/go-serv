package store

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
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
	s.setMongo()

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

func ParseJSON(request *http.Request) (map[string]string, error) {
	var res map[string]string
	b, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
