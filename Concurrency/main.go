package main

import (
	"fmt"
	"log"
	"time"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func Odd(ch chan bool, i int) {
	fmt.Println("API call for", i, "started")
	ch <- true
	time.Sleep(100 * time.Millisecond)
}

func Even(ch chan bool, i int) {
	fmt.Println("API call for", i, "started")
	ch <- true
	time.Sleep(100 * time.Millisecond)
}

func main() {
	numArray := makeRange(0, 1000)
	ch := make(chan bool, len(numArray))

	start := time.Now()

	for i, _ := range numArray {
		if i%2 == 0 {
			go Even(ch, i)
			<-ch
		} else {
			go Odd(ch, i)
			<-ch
		}
	}

	elapsed := time.Since(start)
	log.Printf("Time taken %s", elapsed)
}
