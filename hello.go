package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	var currentYear int
	var birthYear int
	var AgeDifference int

	firstName = "Shalom"
	lastName = "Tata"
	currentYear = 2022
	birthYear = 2000
	AgeDifference = currentYear - birthYear

	currentYear = currentYear + 1

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(AgeDifference)

	AgeDifference = currentYear - birthYear

	fmt.Println(AgeDifference)
}
