package main

import "fmt"

func main() {
	var v1 int
	fmt.Println(v1)

	var ptr = &v1
	fmt.Println(*ptr, ptr, &ptr)
	fmt.Println(v1, &v1)

}
