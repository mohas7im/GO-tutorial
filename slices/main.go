package main

import "fmt"

func main() {
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Println("Array:", numbers)

	sliceNumbers := []int{10, 20, 30, 50}
	sliceNumbers = append(sliceNumbers, 60)
	fmt.Println("Slice:", sliceNumbers)

	fmt.Println("\nLooping through slice:")
	oddNumbers := []int{1, 3, 7, 9, 11}
	for index, number := range oddNumbers {
		fmt.Println("index:", index, "value:", number)
	}
}
