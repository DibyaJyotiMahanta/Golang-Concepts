package main

import "fmt"

func main() {
	type person struct {
		firstName string
		lastName  string
		age       int
	}

	type contact struct {
		email  string
		number string
	}
	type address struct {
		house int
		city  string
		state string
	}

	type employee struct {
		employee_details person
		employee_contact contact
		employee_address address
	}
	var employee1 employee

	employee1.employee_details = person{
		firstName: "Dibya",
		lastName:  "Mahanta",
		age:       23,
	}

	employee1.employee_contact.email = "mahanta@gmail.com"
	employee1.employee_contact.number = "7978742327"

	employee1.employee_address = address{
		house: 1,
		city:  "Bengaluru",
		state: "karnataka",
	}

	fmt.Println(employee1)
}
