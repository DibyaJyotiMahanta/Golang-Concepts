package main

import "fmt"

func main() {
	var marks = map[string]int{
		"Dibya": 100,
		"Kusha": 87,
		"Jute":  90,
	}

	for name, mark := range marks {
		fmt.Println("for", name, "the mark is", mark)
	}
}
