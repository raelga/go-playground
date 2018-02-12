package main

import (
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + (r.URL.Path)[1:]))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":1234", nil)
}
