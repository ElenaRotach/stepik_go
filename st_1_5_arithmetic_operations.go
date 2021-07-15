package main

import "fmt"

func main() {
	a := 100
	b := 10

	c := a + b // с = 110
	fmt.Println(c)

	c = a * b  // с = 1000
	fmt.Println(c)

	c = a - b// с = 90
	fmt.Println(c)

	c = a / b  // с = 10
	fmt.Println(c)

	c = a % b  // с = 0
	fmt.Println(c)

	c++  // с = 1
	fmt.Println(c)

	c--  // с = 0
	fmt.Println(c)
}
