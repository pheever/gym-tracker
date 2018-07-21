package main

import (
	"time"
)

// Athlete type provides details related to a registered athlete
type Athlete struct {
	Name, Surname, Nickname string
	Sex                     bool
	DOB, DOE                time.Time
	Current                 AthleteAtributes
	Historical              []AthleteAtributes
}

// AthleteAtributes provides current athelte attributes
type AthleteAtributes struct {
	Weight, Height float32
	Date           time.Time
}

func (athl *Athlete) String() string {
	return athl.Name + " " + athl.Surname + " " + string(athl.Age())
}

// Age of the athlete
func (athl *Athlete) Age() int {
	age := time.Now().Year() - athl.DOB.Year()
	if athl.DOB.YearDay() > time.Now().YearDay() {
		return age - 1
	}
	return age
}

// IsAdult checkswether the athlete is an adult
func (athl *Athlete) IsAdult() bool {
	return athl.Age() >= 18
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
