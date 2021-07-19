package main
//go run st_1_10_for.go
import "fmt"

func main() {
	sum := 0
	// for [инициализация счетчика]; [условие]; [изменение счетчика]{
	for i := 1; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Объявление счетчика вне цикла
	var a = 1
	for ; a < 10; a++ {
		fmt.Println(a * a)
	}

	// Изсенение счетчика в теле цикла
	var b = 1
	for ; b < 10; {
		fmt.Println(b * b)
		b++
	}
	// Сокращенная запись
	var c = 1
	for c < 10 {
		fmt.Println(c * c)
		c++
	}
	/*
		Бесконечный цикл
		for {
		}
	*/
}
