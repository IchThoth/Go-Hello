package main

import "fmt"

func main() {
	number := 3.14
	var numberPi float32 = float32(number)
	var radius int = 5

	circumference := 2 * numberPi * float32(radius)

	fmt.Println(circumference)

	fmt.Printf("For a Radius of %v the circumference is %.2f.", radius, circumference)
}
