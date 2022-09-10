package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}

var wg sync.WaitGroup
var mut sync.Mutex /* Mutex -> Mutual Exclusion */

func main() {
	/* go greeter("Mayank") */ /* This will start a goroutine */
	/* greeter("Harsh") */

	websiteList := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.github.com",
	}

	for _, website := range websiteList {
		getStatus(website)
	} /* This will take time as they are running one after other. We can make it fast by using goroutines that will achieve parallelism */

	fmt.Printf("\n\n\n")

	for _, website := range websiteList {
		wg.Add(1)
		go getStatusGoRoutine(website) /* This will start a goroutine for each website */
	}
	wg.Wait()

	fmt.Printf("\n\n\n")

	fmt.Println(signals)
}

func greeter(name string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		println("Hello, ", name)
	}
}

func getStatus(endpoint string) {

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	fmt.Println(endpoint, " -> ", res.Status)
}

func getStatusGoRoutine(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	}

	mut.Lock() /* This will lock the mutex. This will prevent other goroutines from accessing the signals slice at the same time */
	signals = append(signals, endpoint)
	mut.Unlock() /* This will unlock the mutex after the append operation is done. */

	defer res.Body.Close()

	fmt.Println(endpoint, " -> ", res.Status)
}
