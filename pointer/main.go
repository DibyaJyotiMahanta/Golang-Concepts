package main

import "fmt"

func changeByAddress(address *int) int {
	return *address + 50
}

func main() {
	var num int = 10
	ans := changeByAddress(&num)

	fmt.Println(ans)
}
