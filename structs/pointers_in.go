package main

import "log"

func changeUsingPointer(s *string){
	log.Println("s is set to in memory location",s)
	newValue :="Red"
	*s=newValue
}