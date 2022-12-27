package main

import "fmt"

func main() {
	// Nambahin komentar
	var firstName string = "john"
	lastName := "wick"
	lastName = "ethan"

	_ = "belajar golang"

	name := new(string)

	var positiveNumber uint8 = 89
	var negativeNumber = -123412345123
	var decimalNumber = 2.62

	var exist bool = true

	var message = `Nama saya "John Pantau".
	 Salam Kenal.
	 Lagi Belajar Golang jancuk!`

	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println(name)
	fmt.Println(*name)
	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
	fmt.Printf("exist? %t \n", exist)
	fmt.Println(message)
}
