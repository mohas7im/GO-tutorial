// package main

// import "fmt"

// func main(){
// 	fmt.Println("Hello World,hi")
// }

// package main

// import "fmt"

// func main(){
// 	name:= "Hashim"
// 	age := 25

// 	fmt.Println("name",name)
// 	fmt.Println("age",age)

// 	if age >= 18{
// 		fmt.Println("you are an adult")

// 	} else{
// 		fmt.Println("you are kid")
// 	}

// }

// package main

// import "fmt"

// func greet(){
// 	fmt.Println("im hashim")
// }

// func main(){
// 	greet()
// }


// package main

// import "fmt"

// func greet(name string){
// 	fmt.Println("Hello",name)
// }

// func main() {
// 	greet("hashim")
// 	greet("go")
// }

// package main
// import "fmt"

// func add (a int, b int) int {
// 	return a + b
// }

// func main() {
// 	sum := add(10,20)
// 	fmt.Println(sum)

// }

// package main
// import "fmt"

// func calculate(a int, b int )(int, int){
// 	sum := a + b
// 	difference := a-b

// 	return sum, difference
// }

// func main(){

// 	sum, difference := calculate(20,5)
	
// 	fmt.Println("sum",sum)
// 	fmt.Println("difference",difference)
// }

package main 
import "fmt"

func calculate(a int , b int) (int, int){
	product := a * b
	division := a/b

	return product, division
}

func main(){
	product, division := calculate(10,5)

	fmt.Println("product",product)
	fmt.Println("division",division)
}