package main

import (
	"fmt"
	"sync"
)

func main() {
	incre := 0
	gr := 1000
	var wg sync.WaitGroup
	wg.Add(gr)
	var mu sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			mu.Lock()
			val := incre
			val++
			incre = val
			fmt.Println(incre)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count:", incre)
}
