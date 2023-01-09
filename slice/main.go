package main

import "fmt"

func main() {
	// var fruits = []string{"apple", "grape", "banana"}
	// var fruitsA = []string{"banana", "melon"} //slice
	// var fruitsB = [2]string{"banana", "melon"} //array
	// var fruitsC = [...]string{"papaya", "grape"} //array
	// fmt.Println(fruits[0])
	// var newFruits = fruits[0:2]
	// fmt.Println(newFruits)

	// var aFruits = fruits[0:3]
	// var bFruits = fruits[1:4]
	// var aaFruits = aFruits[1:2]
	// var baFruits = bFruits[0:1]

	// fmt.Println(fruits)
	// fmt.Println(aFruits)
	// fmt.Println(bFruits)
	// fmt.Println(aaFruits)
	// fmt.Println(baFruits)

	// baFruits[0] = "pinnaple"

	// fmt.Println(fruits)
	// fmt.Println(aFruits)
	// fmt.Println(bFruits)
	// fmt.Println(aaFruits)
	// fmt.Println(baFruits)

	// fmt.Println(len(fruits))
	// fmt.Println(cap(fruits))

	// var aFruits = fruits[0:3]
	// fmt.Println(len(aFruits))
	// fmt.Println(cap(aFruits))

	// var bFruits = fruits[1:4]
	// fmt.Println(len(bFruits))
	// fmt.Println(cap(bFruits))

	// var cFruits = append(fruits, "papaya")
	// fmt.Println(fruits)
	// fmt.Println(cFruits)

	// var bFruits = fruits[0:2]

	// fmt.Println(cap(bFruits))
	// fmt.Println(len(bFruits))

	// fmt.Println(fruits)
	// fmt.Println(bFruits)

	// var cFruits = append(bFruits, "papaya")

	// fmt.Println(fruits)
	// fmt.Println(bFruits)
	// fmt.Println(cFruits)

	// dst := make([]string, 3)
	// src := []string{"watermelon", "pinnaple", "apple", "orange"}
	// n := copy(dst, src)

	// fmt.Println(dst)
	// fmt.Println(src)
	// fmt.Println(n)

	// dst := []string{"potato", "potato", "potato"}
	// src := []string{"watermelon", "pinnaple"}
	// n := copy(dst, src)

	// fmt.Println(dst)
	// fmt.Println(src)
	// fmt.Println(n)

	var fruits = []string{"apple", "grape", "banana"}
	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2]

	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	fmt.Println(aFruits)
	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))

	fmt.Println(bFruits)
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))
}
