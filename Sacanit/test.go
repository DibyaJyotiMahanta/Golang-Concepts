package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your name")

	reader := bufio.NewReader(os.Stdin)

	str, _ := reader.ReadString('\n')

	fmt.Printf("hello %s", str)
}
