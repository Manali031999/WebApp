package main

import (
	"log"
	"sync"
	"time"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func Odd(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("API call for", i, "started")
}

func Even(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("API call for", i, "started")
}

func main() {
	var wg sync.WaitGroup
	numArray := makeRange(0, 1000)
	start := time.Now()
	for i, _ := range numArray {
		if i%2 == 0 {
			wg.Add(1)
			go Even(i, &wg)
		} else {
			wg.Add(1)
			go Odd(i, &wg)
		}
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Time taken %s", elapsed)
}
