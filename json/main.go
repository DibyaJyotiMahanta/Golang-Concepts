package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Person struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		IsAdult bool   `json:"is_adult"`
	}

	var prson = new(Person)
	prson.Name = "Dibya"
	prson.Age = 23
	prson.IsAdult = true

	// fmt.Println(prson)

	data, err := json.Marshal(prson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T \n", data)
	// fmt.Printf("%T", string(data))

	var prsnn Person

	errorrr := json.Unmarshal(data, &prsnn)
	if errorrr != nil {
		fmt.Println(errorrr)
	}
	fmt.Println(prsnn)
}
