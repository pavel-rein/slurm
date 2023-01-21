package main

import "fmt"

func main() {
	var m1 map[int32]bool
	fmt.Println("m1:\t", m1, "len:", len(m1))

	var m2 map[string]string
	fmt.Println("m2:\t", m2, "len:", len(m2))

	m3 := make(map[int]int)
	fmt.Println("m3:\t", m3, "len:", len(m3))

	ages := map[string]int{
		"Andrey": 30,
		"Pavel":  33,
	}
	fmt.Println("ages:\t", ages, "len:", len(ages))

	ages["Svetlana"] = 32
	fmt.Println("ages:\t", ages, "len:", len(ages))

	fmt.Println(ages["Oleg"])
	fmt.Println("ages:\t", ages, "len:", len(ages))
}
