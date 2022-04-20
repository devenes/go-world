package main

import "fmt"

func main() {

	var (
		name      string = "Enes"
		age       int    = 23
		isMarried bool   = false

		weight float32 = 70.5
		height int     = 170
	)

	// var name, age, isMarried, weight, height = "Enes", 25, true, 70.5, 170

	// name, age, isMarried, weight, height := "Enes", 25, true, 70.5, 170

	// var isMarried bool

	fmt.Println(isMarried) // string --> "", numeric --> 0, bool --> false
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(weight)
	fmt.Println(height)
}

// Convert to comment:
/* alt + shif + a  */
// (ctrl + k) + (ctrl + c)
