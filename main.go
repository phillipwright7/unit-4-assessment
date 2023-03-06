package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(calculate(8, 9))
}

func calculate(x int, y int) int {
	return x * y
}
