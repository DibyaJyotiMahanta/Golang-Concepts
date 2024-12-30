package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Animals(num int) {
	defer wg.Done()
	fmt.Println("Dog", num, "Barks")
}

func main() {

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Animals(i)
	}

	wg.Wait()

}
