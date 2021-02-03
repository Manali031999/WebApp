package controller

import (
	"net/http"
)

func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "new.html", nil)
}