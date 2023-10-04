package main

import (
	"log"
	"net/http"
)

func main() {
	port := "8080"
	if port == "" {
		log.Fatal("PORT env must be set")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("Hello World!"))
	})

	server := new(http.Server)
	server.Handler = mux
	server.Addr = "0.0.0.0:" + port

	log.Println("Starting server on port", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
