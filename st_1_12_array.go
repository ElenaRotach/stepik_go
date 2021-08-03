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

	// сравнение массивов
	f := [3]int{1, 2, 3}
	g := [3]int{1, 2, 3}
	i := [3]int{3, 2, 1}

	fmt.Println(f == g) // true
	fmt.Println(f == i) // false

	// обращение по индексу
	fmt.Println(b[0])     // 1
	fmt.Println(b[2])     // 3

	b[0] = 87

	fmt.Println(b)     // 87

	j := [5]int{1, 2, 3, 4, 5}
	fmt.Println(j) // [1 2 3 4 5]

	for k := 0; k < len(j); k++ {
		fmt.Println(j[k])
		// 1
		// 2
		// 3
		// 4
		// 5
	}

	for idx, elem := range j {
		fmt.Printf("Элемент с индексом %d: %d\n", idx, elem)
		// Элемент с индексом 0: 1
		// Элемент с индексом 1: 2
		// Элемент с индексом 2: 3
		// Элемент с индексом 3: 4
		// Элемент с индексом 4: 5
	}
}
