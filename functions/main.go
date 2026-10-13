package main

import "fmt"

func greet() {
	fmt.Println("im hashim")
}

func greetUser(name string) {
	fmt.Println("Hello", name)
}

func add(a int, b int) int {
	return a + b
}

func calculate(a int, b int) (int, int) {
	sum := a + b
	difference := a - b
	return sum, difference
}

func calculateMulDiv(a int, b int) (int, int) {
	product := a * b
	division := a / b
	return product, division
}

func main() {
	greet()

	greetUser("hashim")
	greetUser("go")

	sum := add(10, 20)
	fmt.Println("sum:", sum)

	s, difference := calculate(20, 5)
	fmt.Println("sum:", s)
	fmt.Println("difference:", difference)

	product, division := calculateMulDiv(10, 5)
	fmt.Println("product:", product)
	fmt.Println("division:", division)
}
