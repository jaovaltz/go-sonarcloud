package main

import "fmt"

func main() {
	fmt.Println("sum(1, 2) = ", sum(1, 2))
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}