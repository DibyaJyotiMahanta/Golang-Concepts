package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// fmt.Printf("%T", res) //it is a http.Response type obj

	data, errors := io.ReadAll(res.Body)
	if errors != nil {
		fmt.Println(errors)
		return
	}
	fmt.Printf("%s", string(data))

}
