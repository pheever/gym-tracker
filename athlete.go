package main

import (
	"time"
)

// Athlete type provides details related to a registered athlete
type Athlete struct {
	name, surname string
	sex           bool
	dob           time.Time
	current       AthleteAtributes
	historical    []AthleteAtributes
}

// MakeNewAthlete creates a new athlete
func MakeNewAthlete(name, surname string, dob time.Time) Athlete {
	return Athlete{
		dob:     dob,
		sex:     false,
		name:    name,
		surname: surname,
		current: AthleteAtributes{
			weight: 0,
			height: 0,
			date:   time.Now(),
		},
		historical: make([]AthleteAtributes, 0),
	}
}

// AthleteAtributes provides current athelte attributes
type AthleteAtributes struct {
	weight, height float32
	date           time.Time
}

func (athl *Athlete) String() string {
	return athl.name + " " + athl.surname + " " + athl.dob.UTC().Format(time.RFC3339)
}

// LoadHistoricalData loads athletes historical data from record
func (athl *Athlete) LoadHistoricalData() {
	//athlete.
}
