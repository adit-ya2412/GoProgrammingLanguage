package main

import "log"

func main() {

	// normal for loop in go as follows

	for i := 0; i < 10; i++ {
		log.Println("This is i",i)
	}

	// for iterating over slice, maps we use range
    slice:=[]string{"Aditya","asd","asdf2e","aaw3r1q124e"}

	for _ ,x := range slice{
        log.Println(x)
	}

	myMap:=make(map[string]string)
	myMap["aditya"]="chaudhary"
	myMap["Sakshi"]="Rai"
	myMap["Nutan"]="Chaudhary"

	for i,x:= range myMap{
		log.Println("The value at index ",i," in map is",x)
	}

}
