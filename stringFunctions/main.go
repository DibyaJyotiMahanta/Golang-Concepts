package main

import (
	"fmt"
	"strings"
)

func main() {
	// var data = "apple, banana, guava, pineapple"
	// var fruits []string = strings.Split(data, ",")
	// fmt.Println(fruits)

	// var str = "one two one one three one four"

	// var count int = strings.Count(str, "one")

	// var str = []string{"I", "am", "Dibya", "Jyoti", "Mahanta"}
	// var result string = strings.Join(str, " ")

	var caps string = "I am a hero"
	var capss string = strings.ToUpper(caps)

	fmt.Println(capss)

}
