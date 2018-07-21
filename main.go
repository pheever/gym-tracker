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
	athl := Athlete{
		Name:     "phivos",
		Surname:  "Phivou",
		Nickname: "vazanos",
		DOB:      time.Date(2000, 4, 13, 0, 0, 0, 0, time.UTC),
		DOE:      time.Now(),
		Sex:      true,
		Current: AthleteAtributes{
			Date:   time.Now(),
			Height: 1.74,
			Weight: 116,
		},
	}

	tmpl, err := template.New("Athlete").Funcs(template.FuncMap{
		"Age": athl.Age,
	}).ParseFiles("html-templates/athlete.html")

	if err != nil {
		fmt.Println(err)
		return
	}
	err = tmpl.Execute(resw, athl)
	if err != nil {
		http.Error(resw, err.Error(), http.StatusInternalServerError)
	}

}
