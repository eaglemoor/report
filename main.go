package main

import "fmt"

func main() {
	fmt.Println("Hellow World")
	fmt.Println(Plus(10, 15))
}

func Plus(a, b int) int {
	return a + b
}
