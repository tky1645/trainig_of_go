package main

import (
	"fmt"
)

var (
	example = []string{"The", "Hello", "world"}
	data    []string
)

func main() {

	data = example
	for _, v := range data {
		fmt.Println(v)

		fmt.Printf("input:")

		var answer string
		fmt.Scanln(&answer)

		if v == answer {
			fmt.Println("o")
		} else {
			fmt.Println("x")
		}

	}

}
