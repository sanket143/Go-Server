package home

import (
	"net/http"
)

// Index function
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
