package main

import (
	"fmt"
	"sync"
	"testing"
)

func Test_Dispatcher(t *testing.T) {
	bufferSize := 100
	var dispatcher Dispatcher = NewDispatcher(bufferSize)
	workers := 3
	for i := 0; i < workers; i++ {
		var w WorkerLauncher = &PreffixSuffexWorker{
			prefixS: fmt.Sprintf("WorkerID: %d -> ", i),
			suffixS: " World",
			id: i,
		}
		dispatcher.LaunchWorker(w)
	}

	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)
}