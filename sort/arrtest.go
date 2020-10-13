package main

import (
	"math/rand"
	"time"
	"fmt"
	"./recursion"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	l := make([]int, 10)
	var i int
	for i < 10 {
		l[i] = rand.Intn(100)
		i++
	}
	fmt.Println(l)
	recursion.QuickSort(l, 0, 9, "aesc")
	fmt.Println(l)
}