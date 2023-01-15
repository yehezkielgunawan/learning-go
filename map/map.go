package main

import "fmt"

func main() {
	// var chicken map[string]int
	// chicken = map[string]int{}

	// chicken["januari"] = 50
	// chicken["februari"] = 40

	// fmt.Println("januari", chicken["januari"])
	// fmt.Println("mei", chicken["mei"])

	// // This will be error
	// var data map[string]int
	// data["one"] = 1

	// // This will be normal
	// data = map[string]int{}
	// data["one"] = 1

	// // horizontal
	// var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// // vertical
	// var chicken2 = map[string]int{
	// 	"januari":  50,
	// 	"februari": 40,
	// }

	// // inisialisasi map tanpa nilai awal
	// var chicken3 = map[string]int{}
	// var chicken4 = make(map[string]int)
	// var chicken5 = *new(map[string]int)

	// var chicken = map[string]int{
	// 	"januari":  50,
	// 	"februari": 40,
	// 	"maret":    34,
	// 	"april":    67,
	// }

	// for key, val := range chicken {
	// 	fmt.Println(key, "	\t:", val)
	// }

	// var chicken = map[string]int{"januari": 50, "februari": 40}

	// fmt.Println(len(chicken))
	// fmt.Println(chicken)

	// delete(chicken, "januari")

	// fmt.Println(len(chicken))
	// fmt.Print(chicken)

	// // Cek keberadaan item
	// var chicken = map[string]int{"januari": 50, "februari": 40}
	// var value, isExist = chicken["mei"]

	// if isExist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("item is not exists")
	// }

	// Kombinasi SLice & Map (new syntax)
	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "male"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

}
