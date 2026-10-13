package main  

// import "fmt"

// func main(){
// 	fmt.Println("Branching")

// 	i:= 6

// 	if i<6 {
// 		fmt.Println("i<6")
// 	}else if i>6{
// 		fmt.Println("i>6")

// 	}else{
// 		fmt.Println("equal")
// 	}
// }

// import "fmt"

// func main(){
// 	fmt.Println("Branching")
// 	i:=7

// 	switch i {
// 	case 3:
// 		fmt.Println("value is 3")		
// 	case 5:
// 		fmt.Println("value is 3")
// 	default:
// 		fmt.Println("default")
// 	}
		
// }

import "fmt"
func main(){
	c:= 8

	if c>7 {
		fmt.Println("c is > 7")
		goto mylabel
	}

mylabel:
	if c>7 {
		println("after label : c>7")
	}
}