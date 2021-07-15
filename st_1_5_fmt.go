package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("Введите имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	fmt.Println(name, age)

	// чтение нескольких переменных
	var (
		a, b, c int
	)

	fmt.Print("Введите переменные: ")
	fmt.Scan(&a, &b, &c)

	fmt.Println(a, b, c)
}
