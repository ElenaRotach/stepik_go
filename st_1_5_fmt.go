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
	fmt.Println("=================================================")

	// чтение нескольких переменных
	var (
		a, b, c int
	)

	fmt.Print("Введите переменные: ")
	fmt.Scan(&a, &b, &c)

	fmt.Println(a, b, c)
	fmt.Println("=================================================")

	// вывод в одну/несколько строк (если не строки, то разделяется пробелом)
	fmt.Print("test", 1) // test1
	fmt.Println("")
	fmt.Println("test", 2) // test 2
	fmt.Print(3, 4) // 3 4
	fmt.Println("")
	fmt.Println("=================================================")

	myName := "Helen"
	myAge := 34

	fmt.Println("My name is", myName, "and I am", myAge, "years old.")
}
