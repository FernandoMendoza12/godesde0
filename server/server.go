package server

import "net/http"

func WebServer() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
