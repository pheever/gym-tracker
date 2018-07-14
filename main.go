package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("assigning handlers")
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(resw http.ResponseWriter, req *http.Request) {
	athl := MakeNewAthlete("phivos", "phivou", time.Now())
	athl.Nickname = "vazanos"
	athl.Current.Height = 1.74
	athl.Current.Weight = 115
	athl.Sex = true

	fmt.Println("Current attributes:", athl.Current.Weight, athl.Current.Height)

	tmpl := template.Must(template.ParseFiles("html-templates/athlete.html"))

	tmpl.Execute(resw, athl)
}
