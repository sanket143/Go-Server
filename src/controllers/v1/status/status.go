package status

import "net/http"

// Index Function for Status
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API Status: active"))
}
