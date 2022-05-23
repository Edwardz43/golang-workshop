package main

func sum(i int, ch chan<- int) {
	sum := 0
	for j := 1; j <= i; j++ {
		sum += j * j
	}
	ch <- sum
}
