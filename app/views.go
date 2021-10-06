package app

import (
	"go-serv/server"
	"net/http"
)

func IndexView(w http.ResponseWriter, r *http.Request, s *server.Server) {
	response := "IndexView"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func UsersView(w http.ResponseWriter, r *http.Request, s *server.Server) {
	//var users []server.User
	//err := s.DB.Model(&users).Select()
	//if err != nil {
	//	panic(err)
	//}
	//
	//var usersMap []map[string]string
	//for _, user := range users {
	//	usersMap = append(usersMap, user.Map())
	//}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("It worked"))
}

func StoriesView(w http.ResponseWriter, r *http.Request, s *server.Server) {
	response := "StoriesView"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func AddStory(w http.ResponseWriter, r *http.Request, s *server.Server) {
	response := "AddStory"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
