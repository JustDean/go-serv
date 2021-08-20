package store

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Store struct {
	Config
	router *mux.Router
	//DB *pg.DB
}

func (s *Store) setConfig(ConfigPath string) (*Store, error) {
	_, err := s.Config.readConf(ConfigPath)
	if err != nil {
		return nil, fmt.Errorf("error reading %q: %v", ConfigPath, err)
	}

	return s, nil
}

func (s *Store) GetAddress() (string, error) {
	res := fmt.Sprintf("%s:%s", s.Config.Server.Addr, s.Config.Server.Port)
	if res == ":" {
		return "", fmt.Errorf("addres was not specified")
	}

	return res, nil
}

func (s *Store) AddRoute(path string, f func(http.ResponseWriter, *http.Request, *Store)) *Store {
	s.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		f(w, r, s)
	})

	return s
}

//func  (s *Store) setDatabase(DatabaseConfigPath string) *Store {
//	//dbConfig :=
//	s.DB = pg.Connect(&pg.Options{
//		User: "gopher",
//	})
//
//	return s
//}
