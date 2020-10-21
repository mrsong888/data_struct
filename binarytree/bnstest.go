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
	l := []int{13, 34, 38, 40, 50, 62, 73, 80, 91, 92}
	//var i int
	//for i < 10 {
	//	l[i] = rand.Intn(100)
	//	i++
	//}
	fmt.Println(l)
	//recursion.QuickSort(l, 0, 9, "aesc")
	fmt.Println(bns.LeftBound(l, 50))
}