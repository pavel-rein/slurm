package main

import "fmt"

func main() {
	const earthRadiusInMeter int = 6371000
	fmt.Println(earthRadiusInMeter)

	const (
		secondsInHour = 60
		hoursInDay    = 24
	)
	fmt.Println(secondsInHour, hoursInDay)
}
