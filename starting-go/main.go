package main

import "fmt"

const (
	a = iota
	b
	c
	d
)

func main() {
	fmt.Println("starting go")

	arraySample1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arraySample1)
	fmt.Println(plus(1, 2))

	f := func(x, y int) int { return x + y }
	fmt.Println(f(4, 4))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	// iota
	fmt.Println(a, b, c, d)
}

func plus(x, y int) int {
	return x + y
}

// クロージャーによるジェネレータの実装
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}
