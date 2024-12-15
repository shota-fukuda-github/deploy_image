package main

import (
	"fmt"
	"net/http"
)

func ok(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "hello")
}

func ng(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "ng")
}

func main() {
	server := http.Server{
		Addr:    ":80",
		Handler: nil,
	}

	http.HandleFunc("/ok", ok)
	http.HandleFunc("/ng", ng)
	server.ListenAndServe()
}
