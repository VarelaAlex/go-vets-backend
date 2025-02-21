package middlewares

import (
	"net/http"
	"project/go-vets-backend/jwt"
)

func CheckJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		err := jwt.ProcessJWT(request.Header.Get("Authorization"))
		if err != nil {
			http.Error(response, "Error in Token: "+err.Error(),
				http.StatusBadRequest)
			return
		}
		next.ServeHTTP(response, request)
	}
}
