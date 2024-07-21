package main

import (
	"log"
	"net/http"
)

func main(){
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/tbox/view", tboxView)
	mux.HandleFunc("/tbox/create", tboxCreate)

	// Create a server
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}