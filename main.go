package main

import (
	"fmt"
	"html/template"
	"net/http"

	"ascii-art-web/functions"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlerMainFunc)
	http.HandleFunc("/asciiart", handlerArtFunc)
	fmt.Println("runing server : http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func handlerMainFunc(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		return
	}
	tmpl.Execute(w, nil)
}

func handlerArtFunc(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		return
	}

	data := map[string]string{
		"banner": r.FormValue("banner"),
		"text":   r.FormValue("text"),
	}
	
	data["result"] = functions.HandelAsciiArt(data["text"], data["banner"])
	tmp.Execute(w, data)
}
