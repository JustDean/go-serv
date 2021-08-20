package test_app

import (
	"fmt"
	"net/http"
	"server-test/store"
)

func IndexView(w http.ResponseWriter, r *http.Request, s *store.Store) {
	_, err := fmt.Fprintf(w, "Hello, alone")
	if err != nil {
		panic("Error happened while handling IndexView")
	}
}

//func (s *Store) HealthView(w ResponseWriter, r *Request) {
//	// an example API handler
//	w.Header().Set("Content-Type", "application/json")
//	err := json.NewEncoder(w).Encode(map[string]bool{"ok": true})
//	if err != nil {
//		panic("Error happened while encoding")
//	}
//}
