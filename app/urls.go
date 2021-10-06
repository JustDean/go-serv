package app

import (
	"go-serv/server"
)

func RegisterUrls(s *server.Server) {
	s.HandleRoad("GET", "/", IndexView)
	s.HandleRoad("GET", "/users", UsersView)
	s.HandleRoad("GET", "/stories", StoriesView)
	s.HandleRoad("POST", "/stories/add", AddStory)
}
