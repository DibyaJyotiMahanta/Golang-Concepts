package main

import (
	"fmt"
	"net/http"
)

func deleteReq() {
	url := "https://jsonplaceholder.typicode.com/todos/5"

	req, _ := http.NewRequest(http.MethodDelete, url, nil)

	client := http.Client{}
	res, _ := client.Do(req)
	defer res.Body.Close()

	fmt.Println(res.StatusCode)

}

func main() {
	deleteReq()
}
