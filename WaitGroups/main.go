package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	fmt.Println("num cpu", runtime.NumCPU())
	fmt.Println("runtime", runtime.())
	fmt.Println("---------------")
	counter := 0
	const c = 2
	wg.Add(c)
	for i := 0; i < c; i++ {
		go func() {
			counter++
			fmt.Println("runtime", runtime.NumGoroutine())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("---------------")
	fmt.Println("runtime", runtime.NumGoroutine())
}

