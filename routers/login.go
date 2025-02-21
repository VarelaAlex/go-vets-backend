package routers

import (
	"encoding/json"
	"net/http"
	"project/go-vets-backend/db"
	"project/go-vets-backend/jwt"
	"project/go-vets-backend/models"
)

func Login(response http.ResponseWriter, request *http.Request) {
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(response, "Invalid user and/or password: "+err.Error(),
			http.StatusBadRequest)
		return
	}
	if len(user.Email) == 0 {
		http.Error(response, "Email is a required field", http.StatusBadRequest)
		return
	}
	userWithAllData, exist := db.Login(user.Email, user.Password)
	if !exist {
		http.Error(response, "Invalid user and/or password", http.StatusBadRequest)
		return
	}
	jwtKey, err := jwt.GenerateJWT(userWithAllData)
	if err != nil {
		http.Error(response, "Error while generating JWT Token:"+err.Error(),
			http.StatusBadRequest)
	}
	result := models.LoginResponse{
		Token: jwtKey,
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
