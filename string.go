package main

import "fmt"

func main() {
	var s1 string
	fmt.Println("s1", s1) // ""

	var s2 string = "\tString2\n"
	fmt.Println("s2", s2)

	var s3 string = `\tString3\n`
	fmt.Println("s3", s3)

	s4 := "String4"
	fmt.Println("s4", s4)

	var b byte = 'F'
	fmt.Println("b", b)

	var r rune = 'Щ'
	fmt.Println("r", r)

	s5 := "12345"
	fmt.Println(s5[0], s5[1], s5[2], s5[3], s5[4])

	l1 := len(s5)
	fmt.Println("l1", l1)

	s6 := "Привет"
	fmt.Println("s6[?]:\t", s6[0], s6[1], s6[2], s6[3], s6[4], s6[5], s6[6], s6[7], s6[8], s6[9], s6[10], s6[11])

	l2 := len(s6)
	fmt.Println("l2:\t", l2)

	srez1s5 := s5[0:5]
	fmt.Println("srez1s5:\t", srez1s5)

	srez2s5 := s5[1:4]
	fmt.Println("srez2s5:\t", srez2s5)

	srez3s5 := s5[:4]
	fmt.Println("srez3s5:\t", srez3s5)

	srez4s5 := s5[1:]
	fmt.Println("srez4s5:\t", srez4s5)

	z1s6 := s6[1:]
	fmt.Println("z1s6:\t", z1s6)

	hello := "Hello"
	world := "world"
	welcome := hello + world
	fmt.Println(welcome)

	fmt.Println("ac" > "ab")
	fmt.Println("ac" < "ab")
	fmt.Println("ac" == "ab")
	fmt.Println("AC" == "ac")
	fmt.Println("AC" > "ac")
	fmt.Println("AC" < "ac")

	s7 := "A"
	fmt.Println("A - s7[0]:\t", s7[0])

	s8 := "a"
	fmt.Println("a - s8[0]:\t", s8[0])

	hi := "hi"
	var hi2 string
	hi2 += hi
	fmt.Println("hi2:\t", hi2)
}
