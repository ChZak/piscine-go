package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	length := 0

	for range args {
		length++
	}

	for i := 0; i < length; i++ {
		content, err := ioutil.ReadFile(args[i])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(content))
	}
}
