package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getJob() func(chan bool) {
	n := rand.Int31n(10)
	if n > 2 {
		return job1
	} else {
		return job2
	}
}

func job1(done chan bool) {
	time.Sleep(100 * time.Millisecond)
	done <- true
}

func job2(done chan bool) {
	time.Sleep(600 * time.Millisecond)
	done <- true
}

func doJob(wg *sync.WaitGroup, f func(chan bool)) error {
	done := make(chan bool, 1)
	defer wg.Done()
	go f(done)
	select {
	case <-done:
		return nil
	case <-time.After(500 * time.Millisecond):
		return fmt.Errorf("timeout")
	}
}
