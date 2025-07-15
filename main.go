package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handlerFunction(w http.ResponseWriter, r *http.Request) {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := map[string]string{
		"Body": "",
	}
	tmpl.Execute(w, data)

	if r.Method == "POST" {
		r.ParseForm()

		banner := r.FormValue("banner")
		text := r.FormValue("text")

		fmt.Fprintf(w, "Smitk: %s<br>Lugha: %s<br>Risala: %s", banner, text)
	} else {
		http.ServeFile(w, r, "templates/index.html") // fih dak form li lfo9
	}
}

func main() {
	http.HandleFunc("/", handlerFunction)
	http.ListenAndServe(":8080", nil)
}
