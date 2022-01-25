package main

import "fmt"

func main() {
	a := 10
	b := 2
	result := TestMult(a, b)
	fmt.Println("result", result)
}

func TestMult(a, b int) int {
	return a << b
}
