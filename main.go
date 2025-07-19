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
	fmt.Println("runing server : http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error", err)
	}
}

func handlerMainFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handleError(w, "Not Found!", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		handleError(w, "Internal Server Error!", http.StatusInternalServerError)
		return
	}
	execError := tmpl.Execute(w, nil)
	if execError != nil {
		handleError(w, "Internal Server Error!", http.StatusInternalServerError)
		return
	}
}

func handlerArtFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handleError(w, "Method Not Allowed!", http.StatusMethodNotAllowed)
		return
	}
	tmp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		handleError(w, "Internal Server Error!", http.StatusInternalServerError)
		return
	}
	errorForm := r.ParseForm()
	if errorForm != nil {
		handleError(w, "Bad Request!", http.StatusBadRequest)
		return
	}

	text, checkText := r.PostForm["text"]
	if !checkText {
		handleError(w, "Bad Request!", http.StatusBadRequest)
		return
	}

	banner, CheckBanner := r.PostForm["banner"]
	if !CheckBanner {
		handleError(w, "Bad Request!", http.StatusBadRequest)
		return
	}
	cleanText := strings.ReplaceAll(text[0], "\r", "")
	if len(cleanText) > 1000 || len(cleanText) == 0 {
		handleError(w, "Bad Request!", http.StatusBadRequest)
		return
	}
	result, checkError := functions.HandelAsciiArt(text[0], banner[0])
	if checkError {
		handleError(w, "Bad Request!", http.StatusBadRequest)
		return
	}
	execError := tmp.Execute(w, result)
	if execError != nil {
		handleError(w, "Internal Server Error!", http.StatusInternalServerError)
		return
	}
}

func handleError(w http.ResponseWriter, errorText string, statusCode int) {
	myMap := make(map[string]string)
	myMap["errorText"] = errorText
	myMap["statusCode"] = strconv.Itoa(statusCode)

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "500 Internal Server Error!", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	execError := tmpl.Execute(w, myMap)
	if execError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "500 Internal Server Error!", http.StatusInternalServerError)
		return
	}
}
