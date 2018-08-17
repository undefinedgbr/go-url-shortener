package main

import (
	"github.com/gorilla/mux"
	"./shortener_service"
	"net/http"
	"html/template"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/shorten", shortener_service.ShortenHandler)
	r.HandleFunc("/{url}", shortener_service.FetchHandler)
	http.Handle("/", r)

	err := http.ListenAndServe(":5000", nil)
	fmt.Println(err)
}

func Index(w http.ResponseWriter, r *http.Request) {
	indexPage, _ := template.ParseFiles("index.html")
	indexPage.Execute(w, indexPage)
}