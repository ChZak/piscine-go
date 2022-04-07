package main

import "fmt"

func main() {
	var donnie struct {
		Name     string
		Life     int
		Age      int
		Aircraft int
	}

	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = 1

	fmt.Println(donnie)
}
