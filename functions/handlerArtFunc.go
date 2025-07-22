package functions

import (
	"html/template"
	"net/http"
	"strings"
)

func HandlerArtFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		HandleError(w, "Method Not Allowed!", http.StatusMethodNotAllowed)
		return
	}
	tmp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		HandleError(w, "Internal Server Error!", http.StatusInternalServerError)
		return
	}
	errorForm := r.ParseForm()
	if errorForm != nil {
		HandleError(w, "Bad Request!", http.StatusBadRequest)
		return
	}

	text, checkText := r.PostForm["text"]
	if !checkText {
		HandleError(w, "Bad Request!", http.StatusBadRequest)
		return
	}

	banner, CheckBanner := r.PostForm["banner"]
	if !CheckBanner {
		HandleError(w, "Bad Request!", http.StatusBadRequest)
		return
	}
	cleanText := strings.ReplaceAll(text[0], "\r", "")
	if len(cleanText) > 1000 {
		HandleError(w, "Bad Request!", http.StatusBadRequest)
		return
	}
	result, checkError := HandelAsciiArt(text[0], banner[0])
	if checkError {
		HandleError(w, "Bad Request!", http.StatusBadRequest)
		return
	}
	execError := tmp.Execute(w, result)
	if execError != nil {
		HandleError(w, "Internal Server Error!", http.StatusInternalServerError)
		return
	}
}
