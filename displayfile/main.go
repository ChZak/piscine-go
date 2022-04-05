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
	if length > 1 {
		fmt.Println("Too many arguments")
	} else if length == 0 {
		fmt.Println("File name missing")
	} else if length == 1 {
		content, err := ioutil.ReadFile(args[0]) // content recois le contenu du fichier, et err une erreur si il y en a une
		if err != nil {                          // Un appel reussi a readil renvoi err == nil
			fmt.Println(err.Error())
			return
		}
		fmt.Print(string(content))
	}
}
