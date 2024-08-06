package feed1x

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewServer() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/users/{userID}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("userID")
		if !isValidUserID(userID) {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		feed, err := GetFeed(r.Context(), userID)

		if err != nil {
			log.Printf("err: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, feed)
	}))
	return r
}

func isValidUserID(userID string) bool {
	if len(userID) == 0 {
		return false
	}
	for _, char := range userID {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
