package main

import "fmt"

func main() {
	var v1 int
	fmt.Println(v1) // 0

	var v2 int = 2
	fmt.Println(v2) // 2

	v3 := 3
	fmt.Println(v3) // 3

	v3 = 4
	fmt.Println(v3) // 4

	var v4, v5 = 4, 5
	fmt.Println(v4, v5)

	v5, v4 = v4, v5
	fmt.Println(v4, v5)

	v4, v5, v6 := 1, 2, 3
	fmt.Println(v4, v5, v6)

	v6, v5, v4 = v4, v5, v6
	fmt.Println(v4, v5, v6)

	var (
		v7 = 7
		v8 = "string"
	)
	fmt.Println(v7, v8)

	_ = v7
	_ = v8

	var EarthRadiusInMeter = 6371000
	_ = EarthRadiusInMeter
	fmt.Println(EarthRadiusInMeter)
}
