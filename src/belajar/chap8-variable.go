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

	//declare multi variable
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	fmt.Println(first, second, third)
	var fourth, fifth, sixth string = "empat", "lima", "enam"
	fmt.Println(fourth, fifth, sixth)
	seventh, eight, ninth := "tujuh", "delapan", " sembilan"
	fmt.Println(seventh, eight, ninth)
	one, isTuesday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Println(one, isTuesday, twoPointTwo, say)

	//underscore variable
	/*
		its not used
	*/
	_ = "learn golang"
}
