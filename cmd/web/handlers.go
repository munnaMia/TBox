package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	// Checking the URL: if the URL contain anything wrong it will give a 404 error
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("This is Home page"))
}

func tboxView(w http.ResponseWriter, r *http.Request) {

	// Get the id from URL and converting it to int value
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	// Checking the id is less then 0 or does any error ocurr
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// FprintF method can use http.ResponseWrite object
	fmt.Fprintf(w, "Display a specific Tbox with ID %d...", id)
}

func tboxCreate(w http.ResponseWriter, r *http.Request) {
	// Checing the method type : is that POST or not
	if r.Method != http.MethodPost {

		// Set a header : Allow POST to give idea whats are allow
		w.Header().Set("Allow", http.MethodPost)

		// Give 405 Error using http.Error Shortcut
		http.Error(w, "Methods Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new text box..."))
}