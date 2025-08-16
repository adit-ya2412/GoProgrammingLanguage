package main

import (
	"fmt"
	"log"
)

func main() {

   var whatToSay string
   var saySomeThingElse string
//    var i int
   
//    for( i=0, i<=4;i++){
// 	log.println(saySomeThing(i))
//    }

   whatToSay,_=saySomeThing("Aditya Chaudhary")
   saySomeThingElse,_=saySomeThing("FFAASSAWRE")
   fmt.Println(whatToSay)
   log.Println(saySomeThingElse)

}

func saySomeThing(s string) (string,string) {

	return "Hello , i am returning the string passed to me" + s, "World"

}