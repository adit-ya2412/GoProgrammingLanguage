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

	u :=User{
		FirstName: "Aditya",
		LastName:  "Chaudhary",
	}

	log.Println(u.getFirstName())

	var myString = "Green"

	log.Println("myString is set to",myString)
	changeUsingPointer(&myString) // sending reference for the string value stored
	log.Println("after function call the value changed to ",myString)

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
