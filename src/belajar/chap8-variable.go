package main

import "fmt"

func main() {
	//declare variable with data type
	var firstName string = "John"
	var lastName string
	lastName = "Wick"
	fmt.Printf("halo %s %s !\n", firstName, lastName)

	//declare variable without data type
	var secondName string = "John"
	thirdName := "Wick"
	fmt.Println("halo", secondName, thirdName, "!")
}
