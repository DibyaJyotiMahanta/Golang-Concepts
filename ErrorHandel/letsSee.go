package main

import "fmt"

func divide(a, b int) (error, int) {
	if b == 0 {
		return fmt.Errorf("0 cannot be used to divide"), 0
	}
	return nil, a / b
}

func main() {
	_, result := divide(10, 0)

	fmt.Println(result)
}
