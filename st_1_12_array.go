package main
//go run st_1_12_array.go
import "fmt"

func main() {
	var a [3]int
	var b [3]int = [3]int{1, 2, 3}
	c := [3]int{1, 2, 3}
	d := [...]int{1, 2, 3}
	e := [3]int{1: 12}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
