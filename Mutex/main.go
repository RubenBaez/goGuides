package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	const gr = 20
	inc := 0
	fmt.Println("cpu: ", runtime.NumCPU())
	fmt.Println("runtime: ", runtime.NumGoroutine())

	var wg sync.WaitGroup
	var mt sync.Mutex
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			mt.Lock()
			nv := inc
			nv++
			inc = nv
			fmt.Println(nv)
			mt.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("runtime: ", runtime.NumGoroutine())
}
