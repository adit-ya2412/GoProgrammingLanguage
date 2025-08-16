package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct{
	FirstName string
	LastName string
	PhoneNumber string
	BirthDate time.Time
}

func main() {

	user :=User{
		FirstName: "Aditya",
		LastName: "Chaudhary",
		BirthDate:   time.Date(2000, time.December, 15, 0, 0, 0, 0, time.UTC),

	}
	log.Println("The user",user," current age is ",findAge(user.BirthDate))
	log.Println(s)
}

func findAge(bday time.Time) int {
    now := time.Now()
    age := now.Year() - bday.Year()
    // If birthday hasn't occurred yet this year, subtract 1
    if now.Month() < bday.Month() || (now.Month() == bday.Month() && now.Day() < bday.Day()) {
        age--
    }
    return age
}
