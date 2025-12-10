package main

import "sync"

type Counter struct {
	sync.Mutex
	value int
}

func main() {
	counter := Counter{}

	for i := 0; i < 2; i++ {
		go func(i int) {
			counter.value++
		}(i)
	}
}