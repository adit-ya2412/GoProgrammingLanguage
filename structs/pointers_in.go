package main

import "log"


func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to",myString)
	changeUsingPointer(&myString) // sending reference for the string value stored
	log.Println("after function call the value changed to ",myString)
}
func changeUsingPointer(s *string){
	log.Println("s is set to in memory location",s)
	newValue :="Red"
	*s=newValue
}