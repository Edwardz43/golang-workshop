package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sum(t *testing.T) {
	s := make(chan int, 1)
	sum(5, s)
	assert.Equal(t, 55, <-s)
}
