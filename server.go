package main

import (
	"fmt"
	"net/http"
	"log"
)

var port string

func handleRoot(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println("Enpoint called: " + path)
	http.ServeFile(w, r, "front/" + path)
}

func handleJs(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[4:]
	log.Println("Enpoint called to JS: " + path)
	http.ServeFile(w, r, "node_modules/jquery/dist/" + path)
	http.ServeFile(w, r, "node_modules/bootstrap/dist/js/" + path)
}

func handleCss(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[5:]
	log.Println("Enpoint called to CSS: " + path)
	http.ServeFile(w, r, "node_modules/bootstrap/dist/css/" + path)
}

func main() {
	port = "8080"

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/js/", handleJs)
	http.HandleFunc("/css/", handleCss)

	fmt.Println("Server started on port " + port)

	log.Fatal(http.ListenAndServe(":" + port, nil))

}