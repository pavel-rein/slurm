package main

import "fmt"

func main() {
	var a1 [5]int
	fmt.Println(a1)

	var a2 [0]int
	fmt.Println(a2)

	var a3 [5]string
	fmt.Println(a3)

	var a4 [3]int = [3]int{1, 2, 3}
	fmt.Println(a4)

	var a5 = [3]int{1, 2, 3}
	fmt.Println(a5)

	a6 := [3]int{2, 3, 4}
	fmt.Println(a6)

	a7 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a7)

	l := len(a7)
	fmt.Println(l) // 5
}
