package main

import "fmt"

func main() {
	var fruit [5]string

	fruit[1] = "apple"
	fruit[0] = "pineapple"

	// fmt.Print(fruit)

	// var numbers = [10]int{1, 2, 3, 4}
	fmt.Println(fruit[3])
}
