package main

import "fmt"

func main() {
	var num = "hello, sexy"
	for index, value := range num {
		fmt.Printf("index = %d, value = %c \n", index, value)

	}
}
