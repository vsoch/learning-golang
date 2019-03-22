package main

import (
	"fmt"
	"net/http"
)

func main() {

        // Handler will print a string at the web root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Where have all the flowers gone?")
	})

        // Handler will print a string at the web root
	http.HandleFunc("/flowers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Oh, there they are.")
	})

	// Define the static files folder, strip away folder name
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

        // Serve on Port 8080
	http.ListenAndServe(":8080", nil)
}
