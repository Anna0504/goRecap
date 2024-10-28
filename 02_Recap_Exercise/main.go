package main

import "fmt"

// Constants
const Pi = 3.14159 // Constant value

// Function to add two numbers
func add(x int, y int) int {
	return x + y
}

// Function to calculate the area of a circle
func circleArea(radius float64) float64 {
	return Pi * radius * radius // Using constant Pi
}

func main() {
	// Variables with data types
	var name string = "GoLang" // string data type
	var a int = 10             // int data type
	var b int = 20             // int data type
	var radius float64 = 5.0   // float64 data type

	// Printing a message
	fmt.Println("Welcome to", name)

	// Using the add function
	sum := add(a, b)
	fmt.Println("Sum of", a, "and", b, "is:", sum)

	// Control statement: if-else
	if sum > 25 {
		fmt.Println("Sum is greater than 25")
	} else {
		fmt.Println("Sum is 25 or less")
	}

	// Using the circleArea function
	area := circleArea(radius)
	fmt.Println("Area of the circle with radius", radius, "is:", area)

	// Control statement: for loop (loop from 1 to 5)
	fmt.Println("Counting from 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Control statement: switch
	fmt.Println("Checking the value of sum using switch:")
	switch sum {
	case 30:
		fmt.Println("Sum is 30")
	case 20:
		fmt.Println("Sum is 20")
	default:
		fmt.Println("Sum is neither 20 nor 30")
	}
}
