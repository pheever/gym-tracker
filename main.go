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
	athl.Sex = false

	tmpl := template.Must(template.ParseFiles("html-templates/athlete.html"))

	tmpl.Execute(resw, athl)
	//fmt.Fprintf(resw, "this is just a message passed at %s", athl.String())
	//test change
}
