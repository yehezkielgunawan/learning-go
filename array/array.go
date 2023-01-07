package main

import "fmt"

func main() {
	// 	var names [4]string
	// 	names[0] = "Trafagal"
	// 	names[1] = "d"
	// 	names[2] = "water"
	// 	names[3] = "lawlaw"

	// 	fmt.Println(names[0], names[1], names[2], names[3])

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// fmt.Println("Jumlah element \t\t", len(fruits))
	// fmt.Println("Isi semua element \t", fruits)

	// var fruits [4]string

	// // cara horizontal
	// fruits = [4]string{"apple", "grape", "banana", "melon"}

	// // cara vertical
	// fruits = [4]string{
	// 	"apple",
	// 	"banana",
	// 	"grape",
	// 	"melon",
	// }

	// var numbers = [...]int{2, 3, 4, 5, 6, 7}

	// fmt.Println("data array \t:", numbers)
	// fmt.Println("jumlah element \t:", len(numbers))

	// var numbers1 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	// var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	// fmt.Println("numbers1", numbers1)
	// fmt.Println("numbers2", numbers2)

	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("element %d: %s\n", i, fruits[i])
	// }

	// for i, fruit := range fruits {
	// 	fmt.Printf("elemen %d: %s\n", i, fruit)
	// }

	// for _, fruits := range fruits {
	// 	fmt.Printf("nama buah: %s\n", fruits)
	// }

	// for i, _ := range fruits {
	// 	fmt.Printf("cek %d\n", i)
	// }

	for i := range fruits {
		fmt.Printf("Cek %d\n", i)
	}
}
