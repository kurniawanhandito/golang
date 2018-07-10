package main

import "fmt"

func main() {
	//numerical non-decimal data type
	var positiveNumber uint8 = 89
	var negativeNumber = -124234123
	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	//numericl decimal data type
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	//boolean data type
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	//string data type
	var message string = "halo"
	fmt.Printf("message: %s \n", message)

	var messageTwo = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Println(messageTwo)
}
