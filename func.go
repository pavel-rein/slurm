package main

import "fmt"

func addPrefix(origin string) string {
	return "prefix_" + origin
}

func addPrefixWithErr(origin string) (string, error) {
	return "prefix_" + origin, nil
}

func addPrefixWithLen(origin string) (res string, length int) {
	res = "prefix_" + origin
	length = len(res)
	return
}

func factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return factorial(n-1) * n
}

func main() {
	func1 := addPrefix("string")
	fmt.Println(func1)

	func2, err := addPrefixWithErr("two")
	fmt.Println("func2:", func2, "err:", err)

	func3, len3 := addPrefixWithLen("three")
	fmt.Println("func3:", func3, "len3:", len3)

	fmt.Print("factorial:", factorial(5))
}
