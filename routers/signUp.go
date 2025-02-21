package routers

import (
	"encoding/json"
	"net/http"
	"project/go-vets-backend/db"
	"project/go-vets-backend/models"
)

func SignUp(response http.ResponseWriter, request *http.Request) {
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(response, "Error in the received data: "+err.Error(), http.StatusBadRequest)
		return
	}
	if len(user.Email) == 0 {
		http.Error(response, "Email is a required field", http.StatusBadRequest)
		return
	}
	if len(user.Password) < 6 {
		http.Error(response, "Password should have at least 6 characters", http.StatusBadRequest)
		return
	}
	_, found := db.CheckUserExists(user.Email)
	if found {
		http.Error(response, "The user already exists with the same email", http.StatusBadRequest)
		return
	}
	_, err = db.SignUp(user)
	if err != nil {
		http.Error(response, "An error occurred while trying to register a user: "+err.Error(),
			http.StatusBadRequest)
		return
	}
	response.WriteHeader(http.StatusCreated)
}
