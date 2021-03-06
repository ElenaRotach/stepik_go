package main
//go run st_1_13_task.go
import "fmt"

func main() {
	// Дано трехзначное число. Найдите сумму его цифр.
	//
	//Формат входных данных
	//На вход дается трехзначное число.
	//
	//Формат выходных данных
	//Выведите одно целое число - сумму цифр введенного числа.
	var a, b int
	fmt.Scan(&a)
	b = 0
	b += a % 10
	b += a / 100
	b += (a / 10) % 10
	fmt.Printf("%d", b)
	
	// Дано трехзначное число. Переверните его, а затем выведите.
	//
	//Формат входных данных
	//На вход дается трехзначное число, не оканчивающееся на ноль.
	//
	//Формат выходных данных
	//Выведите перевернутое число.
	var c, d int
	fmt.Scan(&c)
	d = 0
	d += c % 10 * 100
    d += c / 100
    d += ((c / 10) % 10) * 10
	fmt.Printf("%d", d)

}
