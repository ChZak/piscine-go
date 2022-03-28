package main

import "fmt"

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	} else {
		res := 1
		for i := 1; i < nb+1; i++ {
			res *= i
			if res > 2147483647 {
				res = 0
				break
			}
		}
		return res
	}
}

func main() {
	fmt.Println(IterativeFactorial(13))
}
