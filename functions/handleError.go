package functions

import (
	"html/template"
	"net/http"
	"strconv"
)

func HandleError(w http.ResponseWriter, errorText string, statusCode int) {
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
