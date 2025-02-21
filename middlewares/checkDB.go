package middlewares

import (
	"net/http"
	"project/go-vets-backend/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if db.DB == nil {
			http.Error(response, "Connection with the data base lost", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(response, request)
	}
}
