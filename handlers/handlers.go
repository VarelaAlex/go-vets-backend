package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"project/go-vets-backend/middlewares"
	"project/go-vets-backend/routers"
)

func Handlers() {
	router := mux.NewRouter()
	router.HandleFunc("/signUp", middlewares.CheckDB(routers.SignUp)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/getUser", middlewares.CheckDB(middlewares.CheckJWT(routers.GetUser))).Methods("GET")
	router.HandleFunc("/updateUser", middlewares.CheckDB(middlewares.CheckJWT(routers.UpdateUser))).Methods("PUT")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
