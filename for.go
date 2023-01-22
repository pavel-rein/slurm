package main

import "fmt"

func main() {
	a := 1

	if a == 1 {
		fmt.Println("a:", a)
	}

	a2 := 2

	switch a2 {
	case 1:
		fmt.Println("case1:", a2)
	case 2:
		fmt.Println("case2:", a2)
	default:
		fmt.Println("case default", a2)
	}

	a3 := 3

	switch {
	case a3 > 0 && a3 < 1:
		fmt.Println("a3 > 0 or a3 < 1", a3)
	case a3 > 2:
		fmt.Println("a3 > 2", a3)
	case a3 == 3:
		fmt.Println("a3 == 3", a3)
	default:
		fmt.Println("default:", a3)
	}

	a4 := 4

	for a4 == 4 {
		fmt.Println("a4:", a4)
		a4 = 5
		fmt.Println("a4:", a4)
	}

	for i := 0; i < a4; i++ {
		fmt.Println("i:", i)
	}

	sl1 := []int64{4, 5, 6}
	for key, value := range sl1 {
		fmt.Printf("key: %v, val: %v \n", key, value)
	}

	for _, value := range sl1 {
		fmt.Printf("value: %v\n", value)
	}

	ages := map[string]int{
		"pavel":    32,
		"svetlana": 31,
	}

	for key, value := range ages {
		fmt.Printf("key:%v, value:%v\n", key, value)
	}

	word := "Привет"

	for i := 0; i < len(word); i++ {
		fmt.Println(word[i])
		fmt.Printf("%T\n", word[i])
	}

	for key, value := range word {
		fmt.Printf("key: %v, value: %v, value type: %T\n", key, value, value)
	}
}
