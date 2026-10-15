package main

// import("fmt"
// 	   "time"
// )

// func TaskA(){
// 	fmt.Println("Task A")
// 	time.Sleep(time.Second)
// 	fmt.Println("Task A Done")

// }

// func main() {
// 	fmt.Println("concurrency")
// 	go TaskA()
// 	time.Sleep(time.Second * 2)

// 	fmt.Println("Main exit")

// }

import (
	"fmt"
	"sync"
	"time"
) 
 func taskA(wg *sync.WaitGroup){
	fmt.Println("Task A")
	time.Sleep(time.Second)
	fmt.Println("Task A done")
	wg.Done()	
	wg.Done()	

 }

 func main(){
	fmt.Println("Concurrency")

	var wg sync.WaitGroup
	
	wg.Add(2)

	go taskA(&wg)

	wg.Wait()

	fmt.Println("Main exit")
 }