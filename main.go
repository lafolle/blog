package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	SERVE_ADDR = "localhost:8080"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		// serve home page
		http.ServeFile(w, r, "home.html")
	} else {
		// serve blog if corresponding blog exists.
		file_to_serve := fmt.Sprintf("blogs/%s.html", r.URL.Path[1:])
		log.Println("Serving file", file_to_serve)
		http.ServeFile(w, r, file_to_serve)

		// other wise go redirect the user to home page
	}
}

func main() {

	http.HandleFunc("/", mainHandler)
	log.Println("Listening on", SERVE_ADDR)
	log.Fatal(http.ListenAndServe(SERVE_ADDR, nil))

}
