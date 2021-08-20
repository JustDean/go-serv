package test_app

import (
	"go-serv/store"
)

func AddPaths(s *store.Store) {
	s.AddRoute("/", IndexView)
}
