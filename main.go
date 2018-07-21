package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	/*
		athletes := make([]Athlete, 0)
		for i := 0; i < 10; i++ {
			athletes = append(athletes, MakeNewAthlete(fmt.Sprintf("name-%d", i), fmt.Sprintf("surname-%d", i), time.Now(), false))
		}

		for i := 0; i < 4; i++ {
			athletes[4].UpdateCurrentAttributes(112.4+float32(i), 1.74+float32(i))
		}
		for i := 0; i < len(athletes[4].Historical); i++ {
			fmt.Println(i, athletes[4].Historical[i])
		}
		fmt.Println(athletes[4].Current)
	*/
	//*
	fmt.Println("assigning handlers")
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//*/
}

//*
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

	fmt.Println("Current attributes:", athl.Current.Weight, athl.Current.Height, athl.Age(), athl.IsAdult())

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

//*/
