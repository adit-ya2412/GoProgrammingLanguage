package main

import (
	"log"
	"time"
)

type User struct{
	FirstName string
	LastName string
	PhoneNumber string
	BirthDate time.Time
}

func (u *User) getFirstName() string {
	return u.FirstName
}

func main() {

	u :=User{
		FirstName: "Aditya",
		LastName:  "Chaudhary",
	}

	log.Println(u.getFirstName())

}