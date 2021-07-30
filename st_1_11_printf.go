package main
//go run st_1_11_printf.go
import "fmt"

func main() {
	var a bool = false
	fmt.Printf("%t", a) // для вывода значений типа boolean

	var b int = 63
	fmt.Printf("%b", b) // для вывода целых чисел в двоичной системе
	fmt.Println("")

	fmt.Printf("%c", b) // для вывода символов, представленных числовым кодом
	fmt.Println("")

	fmt.Printf("%d", b) // для вывода целых чисел в десятичной системе
	fmt.Println("")

	fmt.Printf("%o", b) // для вывода целых чисел в восьмеричной системе
	fmt.Println("")

	fmt.Printf("%q", b) // для вывода символов в одинарных кавычках
	fmt.Println("")

	fmt.Printf("%x", b) // для вывода целых чисел в шестнадцатеричной системе, буквенные символы числа имеют нижний регистр a-f
	fmt.Println("")

	fmt.Printf("%X", b) // для вывода целых чисел в шестнадцатеричной системе, буквенные символы числа имеют верхний регистр A-F
	fmt.Println("")

	fmt.Printf("%U", b) // для вывода символов в формате кодов Unicode, например, U+1234
	fmt.Println("")

	var c float64 = 123.50000000789
	fmt.Printf("%e", c) // для вывода чисел с плавающей точкой в экспоненциальном представлении, например, -1.234456e+78
	fmt.Println("")

	fmt.Printf("%E", c) // тоже самое что %e но в верхнем регистре, например, -1.234456E+78
	fmt.Println("")

	fmt.Printf("%f", c) // для вывода чисел с плавающей точкой, например, 123.456
	fmt.Println("")

	fmt.Printf("%F", c) // то же самое, что и %f
	fmt.Println("")

	fmt.Printf("%g", c) // %e для огромных экспонент, %f в противном случае
	fmt.Println("")

	fmt.Printf("%G", c) // %E для огромных экспонент, %F в противном случае
	fmt.Println("")

	var str string = "jhgjhgj"
	fmt.Printf("%s", str) // для вывода строки
	fmt.Println("")

	fmt.Printf("%p", b) // для вывода значения указателя - адреса в шестнадцатеричном представлении
	fmt.Println("")

	fmt.Printf("%T", c) // для вывода типа переменной
	fmt.Println("")

	// На вход подается число типа float64. Нужно вывести преобразованное число по правилу:
	// число возводится в квадрат, затем обрезается дробная часть так что остается 4 знака после запятой.
	// Но если число равно или меньше нуля - выводить: "число R не подходит", где R - исходное число с 2 цифрами после запятой и с 2 по ширине.
	// А если число больше чем 10 000 - выводить исходное число в экспоненциальном представлении (знак экспоненты в нижнем регистре).
	var (
		d, e float64
	)
	fmt.Scan(&d)
	if d <= 0 {
		fmt.Printf("число %.2f не подходит", d)
	} else if d > 10000 {
		fmt.Printf("%e", d)
	} else {
		e = d*d
		fmt.Printf("%.4f", e)
	}
	fmt.Println("")

	// Sprintf() работает как и Printf(), за исключением того что она ничего не печатает, а возвращает результат форматирования
	var f float64 = 100.123456789
	result := fmt.Sprintf("%.2f", f)
	fmt.Printf("%q", result) // вывод: "100.12"
	fmt.Println("")

	fmt.Printf("\\") // Обратная косая черта
	fmt.Println("")

	fmt.Printf("\"") // Двойные кавычки
	fmt.Println("")

	fmt.Printf("123\f456\f789") // Табуляция
	fmt.Println("")

	fmt.Printf("\vf\v") // вертикальный таб
	fmt.Println("")

	fmt.Printf("\r Input ")
	fmt.Scan(&a)
	fmt.Println("")

	fmt.Printf("123\b") // возврат
	fmt.Scan(&a)
	fmt.Println("")

	fmt.Printf("\ttest") // Табуляция
	fmt.Println("")

	fmt.Printf("test\ntest") // перевод строки
	fmt.Println("")
	fmt.Printf(`
		This string is on
		multiple lines
		within three single
		quotes on either side.
	`) // многострочная печать
	fmt.Println(`Sammy says,\"The balloon\'s color is red.\"`) // не преобразовывать управляющие символы
}
