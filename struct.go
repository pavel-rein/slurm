package main

import "fmt"

func main() {
	type Point struct {
		x    int
		y    int
		z    string
		list []int64
	}
	p := Point{
		x:    1,
		y:    2,
		z:    "pavel",
		list: []int64{1, 2},
	}
	fmt.Println("p:\t", p)

	fmt.Println("p.x:\t", p.x)

	p.x = 3
	fmt.Println("p.x:\t", p.x)
}
