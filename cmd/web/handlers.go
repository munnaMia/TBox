package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	// Checking the URL: if the URL contain anything wrong it will give a 404 error
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set. Notice that we can pass the slice of file
	// paths as a variadic parameter?
	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// We then use the Execute() method on the template set to write the
	// template content as the response body. The last parameter to Execute()
	// represents any dynamic data that we want to pass in, which for now we'll
	// leave as nil.
	// ------> err = ts.Execute(w, nil)

	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body.

	/*================>> So now, instead of containing HTML directly, our template set contains 3 named templates —
	base , title and main . We use the ExecuteTemplate() method to tell Go that we specifically
	want to respond using the content of the base template (which in turn invokes our title and
	main templates)*/
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
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
