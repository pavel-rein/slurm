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

func print(list ...string) {
	fmt.Println("list:", list)
}

func sum(num ...int) (res int) {
	for _, value := range num {
		res += value
	}
	return
}

func adder(x int) (func(int) int, int) {
	sum := 2
	return func(y int) int {
		y = 2
		return x * sum * y
	}, 3
}

func main() {
	func1 := addPrefix("string")
	fmt.Println(func1)

	func2, err := addPrefixWithErr("two")
	fmt.Println("func2:", func2, "err:", err)

	func3, len3 := addPrefixWithLen("three")
	fmt.Println("func3:", func3, "len3:", len3)

	fmt.Println("factorial:", factorial(5))

	list1 := "hi"
	list2 := "world"
	print(list1, list2)

	num1 := 1
	num2 := 2
	num3 := 3
	fmt.Println("sum:", sum(num1, num2, num3))

	numbers := []int{1, 2, 3}
	fmt.Println("sum numbers:", sum(numbers...))

	f1 := func(s string) (res string) {
		return "prefix_" + s
	}
	fmt.Println("f1:", f1("test"))

	var f2, f3 func(s string) int
	f2 = func(s string) int {
		return len(s)
	}

	f3 = func(s string) int {
		return len(s) + 1
	}
	fmt.Println(f2("test"), f3("test"))

	res1, res2 := adder(2)
	fmt.Printf("res1:%v, res2:%v\n", res1, res2)

	defer func() {
		fmt.Println("defer1")
	}()

	num := 5

	defer func(x int) {
		fmt.Println("x:", x)
	}(num)

	num = 20
	fmt.Println("num:", num)
}
