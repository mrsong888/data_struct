package main

import (
	"math/rand"
	"time"
	"fmt"
	"./bst"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var b bst.BST
	var i int
	for i < 10 {
		b.Insert(rand.Intn(100))
		i++
	}
	fmt.Println(b)
	b.PostOrder()
	fmt.Println()
	b.PostOrderTwo()
	fmt.Println()
	fmt.Println(bst.MaxDepth(b.Root()))
}