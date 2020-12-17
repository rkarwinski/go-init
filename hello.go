package main 

import (
	"fmt"
	"log"
	//"rsc.io/quote"
	"example.com/greetings"
)
//import "fmt"
//import "rsc.io/quote"

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"RK","Rodrigo","Pinguim"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}