package main

import (
	"fmt"
	"go.uber.org/goleak"
	"os"
	"runtime"
	"sync"
	"testing"
)

func TestWorker(t *testing.T) {
	defer goleak.VerifyNone(t)
	fmt.Fprintf(os.Stderr, "%d\n", runtime.NumGoroutine())
	count := 300
	wg := &sync.WaitGroup{}
	errorCount := 0
	for i := 0; i < count; i++ {
		wg.Add(1)
		j := getJob()
		go func(wg *sync.WaitGroup, j func(chan bool)) {
			e := doJob(wg, j)
			if e != nil {
				errorCount++
			}
		}(wg, j)
	}
	wg.Wait()
	fmt.Printf("errorCount: %d\n", errorCount)
	fmt.Fprintf(os.Stderr, "%d\n", runtime.NumGoroutine())
}
