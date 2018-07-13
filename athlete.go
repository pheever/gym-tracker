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

// AthleteAtributes provides current athelte attributes
type AthleteAtributes struct {
	Weight, Height float32
	Date           time.Time
}

func (athl *Athlete) String() string {
	return athl.Name + " " + athl.Surname + " " + athl.dob.UTC().Format(time.RFC3339)
}

// LoadHistoricalData loads athletes historical data from record
func (athl *Athlete) LoadHistoricalData() {
	athl.Historical = make([]AthleteAtributes, 0)
}
