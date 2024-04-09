package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	mux.HandleFunc("GET /books/{id}", func(w http.ResponseWriter, r *http.Request) {
		bookID := r.PathValue("id")
		w.Write([]byte("books - " + bookID))
	})

	http.ListenAndServe(":8080", mux)
}
