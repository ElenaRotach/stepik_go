//go run st_1_8_if_else.go
package main

import (
	"fmt"
	"math"
)

func main() {
	a := 6
	b := 7
	if a < b {
		fmt.Println("a меньше, чем b")
	}

	// краткая запись
	lim := 10.0
	x := 2
	n := 3
	if v := math.Pow(float64(x), float64(n)); v < lim {
		// v видима только в if
		fmt.Println(v)
	}

	// if else
	if a < b {
		fmt.Println("a меньше b")
	} else if a > b {
		fmt.Println("a больше b")
	} else {
		fmt.Println("a равно b")
	}
}
