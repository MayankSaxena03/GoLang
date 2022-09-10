package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}

	wg.Add(3) /* We can add 1 in every goroutine or add one time for all goroutines. We are creating 3 goroutines so we Add(3) */
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Add 1 to score")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Add 2 to score")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Add 3 to score")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)

	// We can use sync.RWMutex{} to lock read operation too. mut.RLock() and mut.RUnlock() are the functions
	/* Run -> "go run --race ." with and without mutex to check for race condition */
}
