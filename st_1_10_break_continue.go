package main
//go run st_1_10_break_continue.go
import "fmt"

func main() {
	var sum = 0
	for i := 1; i <= 10; i++{
		if i % 2 == 0 {
			continue        // переходим к следующей итерации
		}
		if i > 4 {
			break       // если число больше 4 выходим из цикла
		}
		sum += i
	}
	fmt.Println("Сумма: ", sum)    // Сумма: 25
}
