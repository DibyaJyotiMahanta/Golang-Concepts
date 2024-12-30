package main

import "fmt"

func main() {
	num := -1

	switch {
	case num < 0:
		fmt.Println("freezing")
	case num > 0 && num <= 10:
		fmt.Println("cold")
	default:
		fmt.Println("warm")
	}
}
