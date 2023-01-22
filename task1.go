package main

import "fmt"

const PassStatus string = "pass"
const FailStatus string = "fail"

type HealthCheck struct {
	ServiceID int
	status    string
}

func GenerateCheck() (res []HealthCheck) {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			res = append(res, HealthCheck{i, PassStatus})
		} else {
			res = append(res, HealthCheck{i, FailStatus})
		}
	}
	return
}

func main() {
	sl1 := GenerateCheck()
	for key := range sl1 {
		if sl1[key].status == PassStatus {
			fmt.Println(sl1[key].ServiceID)
		}
	}
	fmt.Println("Тут будет выведен идентификатор")
}
