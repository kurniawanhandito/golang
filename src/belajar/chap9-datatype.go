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
}
