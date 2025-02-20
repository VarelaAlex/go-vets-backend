package main

import "net/http"

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)
}
func home(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "index.html")
}
