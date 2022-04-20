package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var c int = 0

	c = a + b
	fmt.Printf("a + b = %d\n", c)

	c = a - b
	fmt.Printf("a - b = %d\n", c)

	c = a * b
	fmt.Printf("a * b = %d\n", c)

	c = a / b
	fmt.Printf("a / b = %d\n", c)

	c = a % b
	fmt.Printf("a %% b = %d\n", c)
}
