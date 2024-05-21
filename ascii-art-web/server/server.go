package server

import (
	asciiart "ascii-art/ascii-art-web/ascii-art"
	"html/template"
	"net/http"
	"strconv"
)

type PageData struct {
	OutputText   string
	ErrorMessage string
	StatusCode   string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//  GET requests only allowed
	if r.Method != "GET" {
		renderErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Serve HTML file for root URL
	if r.URL.Path == "/" {
		//http.ServeFile(w, r, "./statics/index.html")
		tmpl, err := template.ParseFiles("./statics/index.html")
		if err != nil {
			renderErrorPage(w, "internal server error", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, map[string]string{"OutputText": ""})
		if err != nil {
			renderErrorPage(w, "internal server error", http.StatusInternalServerError)
			return
		}
		return
	}
	if r.URL.Path == "/style.css" {
		// Set the content type to CSS
		w.Header().Set("Content-Type", "text/css")
		//serve css file
		http.ServeFile(w, r, "./statics/style.css")
		return
	}
	renderErrorPage(w, "Not found", http.StatusNotFound)

}

func Submit(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.NotFound(w, r)
		return
	}
	if r.Method != "POST" {
		//http.Error(w, "", http.StatusMethodNotAllowed)

		renderErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	//parsing the input form and check possible errors
	err := r.ParseForm()
	if err != nil {
		renderErrorPage(w, "internal server error", http.StatusInternalServerError)
		return
	}
	//get the text form <textarea> with name "text"
	text := r.FormValue("text")

	format := r.FormValue("format")
	//pass the text to ascci() function
	output, err := asciiart.Ascii(text, format)
	//checking errors
	if err != nil {
		renderErrorPage(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Create a PageData struct to pass data to the template
	data := PageData{
		OutputText: output,
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("./statics/index.html")
	if err != nil {
		renderErrorPage(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// Execute the template with the PageData struct
	err = tmpl.Execute(w, data)
	if err != nil {
		renderErrorPage(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func renderErrorPage(w http.ResponseWriter, errMsg string, statusCode int) {
	// Create a PageData struct with the error message
	w.WriteHeader(statusCode)
	data := PageData{
		ErrorMessage: errMsg,
		StatusCode:   strconv.Itoa(statusCode),
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("./statics/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template with the PageData struct
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
