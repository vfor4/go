package main

import (
	"fmt"
)

func main() {
	done := make(chan interface{})
	defer close(done)
	inChan := generator(done, 1, 2, 3, 4)
	for v := range multiply(done, addv(done, inChan, 1), 2) {
		fmt.Println(v)
	}
}

func generator(done <-chan interface{}, values ...int) <-chan int {
	c := make(chan int, len(values))
	go func() {
		defer close(c)
		for _, v := range values {
			select {
			case c <- v:
			case <-done:
				return
			}
		}
	}()
	return c
}

func multiply(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for v := range intStream {
			select {
			case c <- v * multiplier:
			case <-done:
				return
			}
		}
	}()
	return c
}

func addv(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for v := range intStream {
			select {
			case c <- v + additive:
			case <-done:
				return
			}
		}
	}()
	return c
}
