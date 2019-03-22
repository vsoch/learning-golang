package main

import (
	"html/template"
	"net/http"
)

type IceCream struct {
	Name  string
	Eaten bool
}

type IceCreamShop struct {
	Name    string
	Flavors []IceCream
}

func main() {

	tmpl := template.Must(template.ParseFiles("layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := IceCreamShop{
			Name: "Dinosaur's Shop",
			Flavors: []IceCream{
				{Name: "Strawberry", Eaten: false},
				{Name: "Mocha Chip", Eaten: true},
				{Name: "Vanilla", Eaten: false},
			},
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
