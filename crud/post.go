package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Todo struct {
	Userid    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func postReq() {
	var todo Todo
	todo.Userid = 69
	todo.Id = 6969
	todo.Title = "jab web met"
	todo.Completed = true

	var url string = "https://jsonplaceholder.typicode.com/todos"

	res, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
		return
	}

	stringData := string(res)
	readerString := strings.NewReader(stringData)
	res1, err1 := http.Post(url, "application/json", readerString)
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	fmt.Println(res1.StatusCode)
}

func main() {
	postReq()
}
