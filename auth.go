package main

import (
	"log"
	"net/http"
)

func Auth(db *BoltDB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth == "" {
			http.Error(w, "no auth in header", http.StatusUnauthorized)
			return
		}

		email, err := db.GetUser(auth)

		if err != nil {
			http.Error(w, "no user found", http.StatusUnauthorized)
			return
		}

		if email == "" {
			http.Error(w, "no user found", http.StatusUnauthorized)
			return
		}

		log.Println(email, auth)

		next.ServeHTTP(w, r)
	})
}
