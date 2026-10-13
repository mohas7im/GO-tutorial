package main

import "fmt"

// type Number int

// func (num Number) is_even() bool{
// 	return num % 2 == 0

// }
// func main() {
// 	fmt.Println("Oops Concept")
// 	var num Number	 
// 	num =10
// 	println(num)
// 	println(num.is_even())
// }

type Car struct{
	Name string
	Model string
	Year int
}

func (c Car) print()string{
	return c.Name + "-" + c.Model
}

func main(){
	fmt.Println("oops")

	vw:=Car{
		Name  : "Vw",
		Model : "vento",
		Year : 2002,
	}
	fmt.Println(vw)
	fmt.Println(vw.print())
}