package main

import (
	_ "embed"
	"net/http"
)

//go:embed index.html
var indexhtml []byte

func main() {
	http.HandleFunc("/", returnStatus)
	http.ListenAndServe(":80", nil)
}

func returnStatus(w http.ResponseWriter, r *http.Request) {
	w.Write(indexhtml)
}
