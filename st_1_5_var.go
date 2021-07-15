package main
//go run st_1_5_var.go
import "fmt"

func main() {

	var hello string
	var a int = 2019
	var symbol int32 = 'c'
	var (
		name string = "Helen"
		age int = 34
	)
	hello = "Hello Go!"

	fmt.Println(hello)
	fmt.Println(a)
	fmt.Println(string(symbol))
	fmt.Println(name)
	fmt.Println(age)
}
