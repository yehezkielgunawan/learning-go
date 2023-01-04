package main

import (
	"fmt"
)

func main() {
	// Nambahin komentar
	// const firstName string = "john"
	// const lastName = "wick"
	// // lastName = "ethan"

	// _ = "belajar golang"

	// name := new(string)

	// var positiveNumber uint8 = 89
	// var negativeNumber = -123412345123
	// var decimalNumber = 2.62

	// var exist bool = true

	// var message = `Nama saya "John Pantau".
	//  Salam Kenal.
	//  Lagi Belajar Golang jancuk!`

	// const (
	// 	square         = "kotak"
	// 	isToday  bool  = true
	// 	numeric  uint8 = 1
	// 	floatNum       = 2.2
	// )

	// // Const with same data and value type
	// const (
	// 	a = "konstanta"
	// 	b
	// )

	// today pakai string dengan value senin
	// sekarang tipe datanya sama dengan today dan nilainya senin
	// isToday2 dideklarasikan dengan metode type inference dengan tipe data bool dan nilainya true
	// const (
	// 	today string = "senin"
	// 	sekarang
	// 	isToday2 = true
	// )

	// // Cara declare multiple const
	// const satu, dua = 1, 2
	// const three, four string = "tiga", "empat"

	// fmt.Printf("halo %s %s!\n", firstName, lastName)
	// fmt.Print("Hallo ", firstName)
	// fmt.Print("Nice to meet you ", lastName, "!\n")
	// fmt.Println(name)
	// fmt.Println(*name)
	// fmt.Printf("bilangan positif: %d\n", positiveNumber)
	// fmt.Printf("bilangan negatif: %d\n", negativeNumber)
	// fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	// fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
	// fmt.Printf("exist? %t \n", exist)
	// fmt.Println(message)

	// var value = (((2+6)%3)*4 - 2) / 3
	// var isEqual = (value == 2)

	// var left = false
	// var right = true

	// var leftAndRight = left && right
	// fmt.Printf("left && right \t (%t) \n", leftAndRight)

	// var leftOrRight = left || right
	// fmt.Printf("left || right \t (%t)\n", leftOrRight)

	// var leftReverse = !left
	// fmt.Printf("!left \t\t(%t)\n", leftReverse)

	// fmt.Printf("nilai %d (%t)", value, isEqual)

	// var point = 4

	// if point == 10 {
	// 	fmt.Println("Lulus sempurna")
	// } else if point > 5 {
	// 	fmt.Println("Lulus")
	// } else if point == 4 {
	// 	fmt.Println("Hampir lulus")
	// } else {
	// 	fmt.Printf("Tidak lulus, nilai anda %d\n", point)
	// }

	// var point = 8840.0

	// if percent := point / 100; percent >= 100 {
	// 	fmt.Printf("%.1f%s perfect", percent, "%")
	// } else if percent >= 70 {
	// 	fmt.Printf("%.1f%s good\n", percent, "%")
	// } else {
	// 	fmt.Printf("%.1f%s bad\n", percent, "%")
	// }

	// var point = 6

	// switch {
	// case point == 8:
	// 	fmt.Println("Perfect")
	// case point < 8 && point > 3:
	// 	fmt.Println("Awesome")
	// 	fallthrough
	// default:
	// 	{
	// 		fmt.Println("Bad")
	// 		fmt.Println("You can be better!")
	// 	}
	// }

	var point = 10

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("Nice")
		}
	} else {
		if point == 5 {
			fmt.Println("Not bad")
		} else if point == 3 {
			fmt.Println("Keep trying")
		} else {
			fmt.Println("You can do it")
			if point == 0 {
				fmt.Println("Gk ketolong")
			}
		}
	}
}
