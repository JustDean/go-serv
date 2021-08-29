package store

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"time"
)

type Store struct {
	Config
	router   *mux.Router
	Database *mongo.Client
}

func (s *Store) setConfig(ConfigPath string) (*Store, error) {
	_, err := s.Config.readConf(ConfigPath)
	if err != nil {
		return nil, fmt.Errorf("error reading %q: %v", ConfigPath, err)
	}

	return s, nil
}

func (s *Store) setMongo() *Store {
	uri := fmt.Sprintf("mongodb://%s:%s", s.Config.Database.Host, s.Config.Database.Port)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	s.Database = client

	defer func() {
		err = s.Database.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
	}()

	if err := s.Database.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Database: successfully connected and pinged.")

	return s
}

func (s *Store) GetAddress() (string, error) {
	res := fmt.Sprintf("%s:%s", s.Config.Server.Addr, s.Config.Server.Port)
	if res == ":" {
		return "", fmt.Errorf("addres was not specified")
	}

	return res, nil
}

func (s *Store) GET(path string, f func(http.ResponseWriter, *http.Request, *Store)) *Store {
	s.router.Path(path).HandlerFunc(func(w http.ResponseWriter, r *http.Request) { f(w, r, s) }).Methods("GET")

	return s
}

func (s *Store) POST(path string, f func(http.ResponseWriter, *http.Request, *Store)) *Store {
	s.router.Path(path).HandlerFunc(func(w http.ResponseWriter, r *http.Request) { f(w, r, s) }).Methods("POST")

	return s
}

//func (s *Store) setDatabase() *Store {
//	// Postgres Support
//	databaseOptions := &pg.Options{
//		Addr:     fmt.Sprintf("%s:%s", s.Config.Database.Host),
//		User:     s.Config.Database.User,
//		Password: s.Config.Database.Password,
//		Database: s.Config.Database.Database,
//	}
//
//	s.Database = pg.Connect(databaseOptions)
//
//	return s
//}
