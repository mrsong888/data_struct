package main

import (
	"math/rand"
	"time"
	"fmt"
	"./bns"
	//"../sort/recursion"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	l := []int{13, 34, 38, 41, 50, 64, 73, 80, 91, 92}
	//var i int
	//for i < 10 {
	//	l[i] = rand.Intn(100)
	//	i++
	//}
	fmt.Println(l)
	//recursion.QuickSort(l, 0, 9, "aesc")
	fmt.Println(bns.BinarySearch(l, 73))
}