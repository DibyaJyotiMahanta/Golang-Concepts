package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num string = "12.1"

	num1, error := strconv.ParseInt(num, 10, 64)

	fmt.Printf("%d", num1)
	fmt.Println(error)
}
