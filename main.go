package main

import (
	"fmt"
)

func main() {
	values := []int{1, 2, 3, 4}
	for _, v := range add(mul(values, 2), 1) {
		fmt.Println(v)
	}
}

func mul(arr []int, mulier int) []int {
	m := make([]int, len(arr))
	for i, v := range arr {
		m[i] = v * mulier
	}
	return m
}

func add(arr []int, addlier int) []int {
	r := make([]int, len(arr))
	for k, v := range arr {
		r[k] = v + addlier
	}
	return r
}
