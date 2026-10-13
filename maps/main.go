package main

import "fmt"

func main() {
	patients := map[int]string{
		1: "hashim",
		2: "john",
	}

	fmt.Println("Map output:", patients)

	fmt.Println("\nIterating map:")
	for id, name := range patients {
		fmt.Println("ID:", id, "Name:", name)
	}
}
