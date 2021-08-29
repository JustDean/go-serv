package store

//
//import (
//	"fmt"
//	"net/http"
//	"regexp"
//)
//
//func initRoad() *http.ServeMux{
//	return http.NewServeMux()
//}
//
//func (r *Router) GET(path string, f func(http.ResponseWriter, *http.Request, *Store)) *Router{
//	thePath := fmt.Sprintf("^%s$", path)
//	compiledPath := regexp.MustCompile(thePath)
//	http.HandleFunc(compiledPath, f(http.ResponseWriter, *http.Request))
//
//	r.AvailableRoutes = append(r.AvailableRoutes, path)
//}
