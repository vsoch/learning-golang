package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/flowers/{flower}/color/{color}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprint(w, vars["flower"], " ", vars["color"])
	})
	http.ListenAndServe(":8080", r)
}
