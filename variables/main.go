// Package Clause
package main

// Import Statement
import "fmt"

// Main Function
func main() {

	// var name1 string = "Enes"
	// var name2 = "Enes"
	name := "Enes"

	job := "Software Engineer"
	age := 23
	age += 1
	var deparment string = "Software Engineering"

	var isMarried bool = false
	if isMarried {
		fmt.Println("I am married")
	} else {
		fmt.Println("I am not married")
	}

	fmt.Printf("Hello, %s!\n", name)
	fmt.Printf("You are %d years old.\n", age)
	fmt.Println("You are working as a", job)
	fmt.Printf("You work in %s.\n", deparment)
	fmt.Printf("My name is %s, I am %d years old, I am %s and I am %t\n", name, age, job, isMarried)

	fmt.Printf("%d + %d = %d", age, age, age+age)
}
