package main

import (
	"fmt"
	"net/http"
	"text/template"

	chi "github.com/go-chi/chi"
)


func main() {
	r := chi.NewRouter()
	r.Handle("/static/*",http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		templ , err := template.ParseFiles("templates/index.html")
		if err != nil {
			fmt.Printf("Error is %v", err)			
		}
		err = templ.ExecuteTemplate(w,"Base",nil) 
		if err != nil {
			fmt.Printf("Error is  : %v", err)
		}
	})

	http.ListenAndServe(":8080", r)
}

