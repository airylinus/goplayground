package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 1; i < 10; i++ {
			fmt.Println("A:", i)
		}
		wg.Done()
	}()
	go func() {
		// defer wg.Done()
		for i := 1; i < 10; i++ {
			fmt.Println("B:", i)
		}
		wg.Done()
	}()
	wg.Wait()
}
