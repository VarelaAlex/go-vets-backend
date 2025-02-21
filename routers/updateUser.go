package routers

import (
	"encoding/json"
	"net/http"
	"project/go-vets-backend/db"
	"project/go-vets-backend/models"
)

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(response, "Incorrect data in *update user*: "+err.Error(),
			http.StatusBadRequest)
		return
	}
	err = db.UpdateUser(user)
	if err != nil {
		http.Error(response, "Error trying to modify the user: "+err.Error(),
			http.StatusBadRequest)
		return
	}
	response.WriteHeader(http.StatusCreated)
}
