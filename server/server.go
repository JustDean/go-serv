package server

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

///////////////////////////// STRUCT ////////////////////////////
type Server struct {
	Config
	router      *mux.Router
	Middlewares []func(next http.Handler) http.Handler
	DB          *pg.DB
}

///////////////////////////// SETUP ////////////////////////////
func (s *Server) setConfig(ConfigPath string) (*Server, error) {
	_, err := s.Config.ReadConfig(ConfigPath)
	if err != nil {
		return nil, fmt.Errorf("error reading %q: %v", ConfigPath, err)
	}

	return s, nil
}

//func (s *Server) setDatabase() (*Server, error){
//	// TODO
//	databaseConfig := s.Config.Database
//}

///////////////////////////// UTILS ////////////////////////////
func (s *Server) GetAddress() (string, error) {
	res := fmt.Sprintf("%s:%s", s.Config.Server.Addr, s.Config.Server.Port)

	// in case the address was not specified
	if res == ":" {
		return "", fmt.Errorf("addres was not specified")
	}
	return res, nil
}
func SetServer(ConfigPath string) *Server {
	s := new(Server)
	// get server config
	_, err := s.setConfig(ConfigPath)
	if err != nil {
		panic(err)
	}
	// create connection to database
	databaseConfig := s.Config.Database
	s.DB = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", databaseConfig.Addr, databaseConfig.Port),
		User:     databaseConfig.User,
		Password: databaseConfig.Password,
		Database: databaseConfig.Database,
	})

	// create route handler
	s.router = mux.NewRouter()

	return s
}
func RunServer(s *Server) {
	// making synchronization channels to gracefully shutdown the server
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// running server
	address, err := s.GetAddress()
	if err != nil {
		panic(err)
	}
	fmt.Printf("server is started on %s\n", address)
	srv := &http.Server{
		Handler: s.router,
		Addr:    address,
	}

	// handle shutdown
	go func() {
		<-sigs
		// close databaseConnection
		err := s.DB.Close()
		if err != nil {
			panic(err)
		}
		// shutdown server
		err = srv.Shutdown(context.TODO())
		if err != nil {
			panic(err)
		}
	}()

	// starting and logging server
	log.Fatal(srv.ListenAndServe())
}

///////////////////////////// HANDLERS ////////////////////////////
func (s *Server) HandleRoad(method string, path string, handleFunction func(http.ResponseWriter, *http.Request, *Server)) {
	s.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handleFunction(w, r, s)
	}).Methods(method)

	for _, middleware := range s.Middlewares {
		s.router.Use(middleware)
	}
}

///////////////////////////// MIDDLEWARES ////////////////////////////
func (s *Server) AddMiddleware(f func(http.Handler) http.Handler) {
	s.Middlewares = append(s.Middlewares, f)
}
