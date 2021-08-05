package main
//go run st_1_12_slices.go
import "fmt"

func main() {
	var a []int
	var b []int = []int{1, 2, 3}
	c := []int{1, 2, 3}
	d := []int{1: 12}

	fmt.Println(a) // []
	fmt.Println(b) // [1 2 3]
	fmt.Println(c) // [1 2 3]
	fmt.Println(d) // [0 12]

	// make([]T, length, capacity)
	e := make([]int, 10, 10) // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(e)

	// создание среза из исходного
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"} // базовый массив

	users1 := initialUsers[2:6] // с 3-го по 6-й
	users2 := initialUsers[:4] // с 1-го по 4-й
	users3 := initialUsers[3:] // с 4-го до конца

	fmt.Println(users1) // [Kate Sam Tom Paul]
	fmt.Println(users2) // [Bob Alice Kate Sam]
	fmt.Println(users3) // [Sam Tom Paul Mike Robert]

	// добавление элементов в срез
	b = append(b, 4, 5)
	fmt.Println(b)

	// удаление элементов из среза
	f := []int{1, 2, 3, 4, 5, 6, 7}
	f = append(f[0:2], f[3:]...)
	fmt.Println(f) // [1 2 4 5 6 7]

	// копирование элементов среза
	g := []int{1, 2, 3}
	h := make([]int, 3, 3)
	n := copy(h, g)

	fmt.Printf("g = %v\n", g)                  // g = [1 2 3]
	fmt.Printf("h = %v\n", h)                  // h = [1 2 3]
	fmt.Printf("Скопировано %d элемента\n", n) // Скопировано 3 элемента
}
