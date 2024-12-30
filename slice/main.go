package main

import "fmt"

func main() {

	var name []int
	name = append(name, 100, 202, 2)

	fmt.Println(name)
	fmt.Println(len(name))
	fmt.Println(cap(name))
}
