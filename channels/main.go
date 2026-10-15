package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Golang Channels")
	ch := make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(1)

	

	go func () {
		fmt.Println("hello from pub function")
		ch <- "hello"
	}()

	go func() {
		fmt.Println("hello from sub function")
		fmt.Println(<-ch)
		wg.Done()
	}()

	wg.Wait()



}