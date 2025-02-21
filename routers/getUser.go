package routers

import (
	"encoding/json"
	"net/http"
	"project/go-vets-backend/db"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(response, "You should sent the ID parameter",
			http.StatusBadRequest)
		return
	}
	user, err := db.GetUser(ID)
	if err != nil {
		http.Error(response, "Error trying to find the record: "+err.Error(),
			http.StatusBadRequest)
		return
	}
	response.Header().Set("context-type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(user)
}
