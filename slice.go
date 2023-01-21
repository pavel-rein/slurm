package main

import "fmt"

func main() {
	var list1 = []int64{1}
	fmt.Println("list1:\t", list1, &list1)
	fmt.Println("len list1:\t", len(list1))
	fmt.Println("cap list1:\t", cap(list1))

	var ptr = &list1
	fmt.Println("ptr:\t", ptr, *ptr)

	list1 = []int64{1, 2, 3, 4, 5}
	fmt.Println("list1:\t", list1, &list1)
	fmt.Println("len list1:\t", len(list1))
	fmt.Println("cap list1:\t", cap(list1))

	fmt.Println("ptr:\t", ptr, *ptr)

	list2 := make([]int64, 0, 5)
	fmt.Println("list2:\t", list2, "\tlen:", len(list2), "\tcap:", cap(list2))

	list2 = make([]int64, 5, 5)
	fmt.Println("list2:\t", list2, "\tlen:", len(list2), "\tcap:", cap(list2))

	list2 = []int64{1, 2, 3, 4, 5}
	fmt.Println("list2:\t", list2, "\tlen:", len(list2), "\tcap:", cap(list2))

	list2 = append(list2, 6)
	fmt.Println("list2:\t", list2, "\tlen:", len(list2), "\tcap:", cap(list2))

	list3 := []int{1, 2, 3}
	fmt.Println("list3:\t", list3, "\tlen:", len(list3), "\tcap:", cap(list3))

	func(list3 []int) {
		list4 := list3[1:]
		list4[0] = 4
		list4[1] = 5
		fmt.Println("list3:\t", list4, "\tlen:", len(list4), "\tcap:", cap(list4))

		list4 = append(list4, 6)
		fmt.Println("list3:\t", list4, "\tlen:", len(list4), "\tcap:", cap(list4))
	}(list3)
	fmt.Println("list3:\t", list3, "\tlen:", len(list3), "\tcap:", cap(list3))
}
