package functions

import (
	"html/template"
	"net/http"
)

func HandlerMainFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		HandleError(w, "Not Found!", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		HandleError(w, "Internal Server Error!", http.StatusInternalServerError)
		return
	}
	execError := tmpl.Execute(w, nil)
	if execError != nil {
		HandleError(w, "Internal Server Error!", http.StatusInternalServerError)
		return
	}
}
