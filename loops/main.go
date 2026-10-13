// package main
// import "fmt"

// func main(){
// 	for{
// 		fmt.Println("hello hashim")
// 		break
// 	}
// }

// package main
// import "fmt"

// func main(){
// 	i:= 1

// 	for i <= 5 {
// 		fmt.Println(i)
// 		i++
// 	}
// }

// package main
// import "fmt"

// func main(){
// 	for i:=0;i<5;i++{
// 		fmt.Println(i)
// 	}

// }

// package main

// import "fmt"
// func main (){
// 	numbers := [4]int{1,1,2,3}

// 	for i,v:= range numbers {
// 		fmt.Println(i,v)
// 	}
	
// }

package main
import "fmt"

func main(){
	letters := map[string]string{
	"a": "A",
	"b": "B",
	"c": "C",
	}

	for _,v := range letters{
		 fmt.Println(v)
	}
	
}
