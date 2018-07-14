package main

import (
	"time"
)

// Athlete type provides details related to a registered athlete
type Athlete struct {
	Name, Surname, Nickname string
	Sex                     bool
	dob                     time.Time
	Current                 AthleteAtributes
	Historical              []AthleteAtributes
}

// AthleteAtributes provides current athelte attributes
type AthleteAtributes struct {
	Weight, Height float32
	Date           time.Time
}

// MakeNewAthlete creates a new athlete
func MakeNewAthlete(name, surname string, dob time.Time) Athlete {
	return Athlete{
		dob:     dob,
		Sex:     false,
		Name:    name,
		Surname: surname,
		Current: AthleteAtributes{
			Weight: 0,
			Height: 0,
			Date:   time.Now(),
		},
		Historical: make([]AthleteAtributes, 0),
	}
}

func (athl *Athlete) String() string {
	return athl.Name + " " + athl.Surname + " " + athl.dob.UTC().Format(time.RFC3339)
}

// LoadHistoricalData loads athletes historical data from record
func (athl *Athlete) LoadHistoricalData() {
	athl.Historical = make([]AthleteAtributes, 0)
}

// UpdateCurrentAttributes updates athletes currentattributes
func (athl *Athlete) UpdateCurrentAttributes(weight, height float32) {
	athl.Historical = append(athl.Historical,
		AthleteAtributes{
			Height: athl.Current.Height,
			Weight: athl.Current.Weight,
			Date:   athl.Current.Date,
		})
	athl.Current.Height = height
	athl.Current.Weight = weight
	athl.Current.Date = time.Now()
}
