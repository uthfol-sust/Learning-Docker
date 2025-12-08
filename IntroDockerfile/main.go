package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Intro to docker file")

    mux := http.NewServeMux()

   
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Docker!"))
	})

    fmt.Println("Server Running on Port : 8080..")
	http.ListenAndServe(":8080",mux)
}