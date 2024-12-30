package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func putReq() {
	type Todo struct {
		Userid    int    `json:"userId"`
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	var todo = Todo{Userid: 123, Id: 123, Title: "Sharmila loves me", Completed: false}

	jsonData, _ := json.Marshal(todo)
	stringJsonData := string(jsonData)
	bufferData := strings.NewReader(stringJsonData)

	urls := "https://jsonplaceholder.typicode.com/todos/1"

	req, _ := http.NewRequest(http.MethodPut, urls, bufferData)
	req.Header.Set("content-type", "application/json")

	client := http.Client{}
	res, _ := client.Do(req)
	defer res.Body.Close()

	response, _ := io.ReadAll(res.Body)

	fmt.Println(string(response))
	fmt.Println(res.StatusCode)

}

func main() {
	putReq()
}
