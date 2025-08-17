package main

import "log"

func main() {
	var isTrue bool
	isTrue = true
	if isTrue {
		log.Println("Its true buddy i told you")
	} else {
		log.Println("Its not true, see you never listen")
	}

	batteryCharge := 100
	if isTrue && batteryCharge > 69 {
		log.Println("I am happy Toady")
	} else {
		log.Println("HI hi haa haa")
	}
	chargeDraining:=batteryCharge<80
	log.Println(chargeDraining)


	switch  chargeDraining {
	case true:
		log.Println("gGG my charge is full")
	case false:
		log.Println("hush its draining")
	case false:
		log.Println("Opps do you see a charging station")
	default:
		log.Println("GG mmaa eqwrtg cool")
	}
}