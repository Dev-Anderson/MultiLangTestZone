package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from route /index")
}

func DynamicHandler(w http.ResponseWriter, r *http.Request) {
	param := r.PathValue("param")
	fmt.Fprintf(w, "Hello from route /index/%s", param)
}

func main() {
	http.HandleFunc("GET /index", IndexHandler)
	http.HandleFunc("GET /index/{param}", DynamicHandler)

	http.ListenAndServe(":8080", nil)
}
