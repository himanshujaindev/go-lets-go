package main

import (
	"fmt"
	// "math/rand"
	"sync"
	"time"
)

// var m = sync.Mutex{}
var m = sync.RWMutex{}
var wg = sync.WaitGroup{} // Counters
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func goroutine() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)    // increment counter by 1
		go dbCall(i) // Runs the fn concurrently
	}
	wg.Wait() // counter back to 0 ; end of goroutine
	fmt.Printf("Results = %v\n", results)
	fmt.Printf("Execution time = %v\n", time.Since(t0))
}

func dbCall(i int) {
	// Simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	// fmt.Println("Result from DB = ", dbData[i])

	// m.Lock()
	// results = append(results, dbData[i]) // muliple processes writing data to same location, result in unexpected output -> use mutex
	// m.Unlock()

	save(dbData[i])
	log()

	wg.Done() // decrement counter
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("Current results = %v\n", results)
	m.RUnlock()
}
