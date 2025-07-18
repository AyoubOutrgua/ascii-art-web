package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"ascii-art-web/functions"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlerMainFunc)
	http.HandleFunc("/asciiart", handlerArtFunc)
	// http.HandleFunc("/error", handleError)
	fmt.Println("runing server : http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error", err)
	}
}

func handlerMainFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.Error(w, "404 Not Found!", http.StatusNotFound)
		handleError(w, "Not Found!", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error!", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func handlerArtFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// http.Error(w, "405 Method Not Allowed!", http.StatusMethodNotAllowed)
		handleError(w, "Method Not Allowed!", http.StatusMethodNotAllowed)
		return
	}
	tmp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error!", http.StatusInternalServerError)
		return
	}
	errorForm := r.ParseForm()
	if errorForm != nil {
		// http.Error(w, "400 Bad Request!", http.StatusBadRequest)
		handleError(w, "Bad Request!", http.StatusBadRequest)
		return
	}

	text, checkText := r.PostForm["text"]
	if !checkText {
		// http.Error(w, "400 Bad Request!", http.StatusBadRequest)
		handleError(w, "Bad Request!", http.StatusBadRequest)
		return
	}

	banner, CheckBanner := r.PostForm["banner"]
	if !CheckBanner {
		// http.Error(w, "400 Bad Request!", http.StatusBadRequest)
		handleError(w, "Bad Request!", http.StatusBadRequest)
		return
	}
	cleanText := strings.ReplaceAll(text[0], "\r", "")
	if len(cleanText) > 1000 || len(cleanText) == 0 {
		// http.Error(w, "400 Bad Request!", http.StatusBadRequest)
		handleError(w, "Bad Request!", http.StatusBadRequest)
		return
	}
	result, checkError := functions.HandelAsciiArt(text[0], banner[0])
	if checkError {
		// http.Error(w, "400 Bad Request!", http.StatusBadRequest)
		handleError(w, "Bad Request!", http.StatusBadRequest)
		return
	}
	tmp.Execute(w, result)
}

func handleError(w http.ResponseWriter, errorText string, statusCode int) {
	myMap := make(map[string]string)
	myMap["errorText"] = errorText
	myMap["statusCode"] = strconv.Itoa(statusCode)

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error!", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, myMap)
}
