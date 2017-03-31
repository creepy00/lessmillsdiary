package main

import (
	//"fmt"
	"net/http"
	"log"
)

func responseHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "front/index.html")
}

func main() {
	http.HandleFunc("/", responseHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}