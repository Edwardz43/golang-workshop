package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(30)
	for i := 0; i < 30; i++ {
		go func(i int) {
			defer wg.Done()
			log.Printf("%d start job\n", i)
			d := rand.Int31n(3000)
			time.Sleep(time.Duration(d))
			log.Printf("%d start finish\n", i)
		}(i)
	}
	wg.Wait()
}
