//go run st_1_7_const.go
package main

import "fmt"

const(
	A int = 45
	B
	C float32 = 3.3
	D
)

func main() {
	const pi float64 = 3.1415
	const (
		a int = 45
		b float32 = 3.3
	)

	fmt.Println(pi, a, b, A, B, C, D)
}