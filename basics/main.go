package main

import "fmt"

func main() {
	fmt.Println("Hello World, hi")

	name := "Hashim"
	age := 25

	fmt.Println("name:", name)
	fmt.Println("age:", age)

	if age >= 18 {
		fmt.Println("you are an adult")
	} else {
		fmt.Println("you are kid")
	}
}
