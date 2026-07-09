package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// I got the idea for the anonymous function from
	// https://stackoverflow.com/questions/12224042/go-tour-exercise-7-binary-trees-equivalence
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}

		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}

	walker(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v1 := range ch1 {
		v2, ok := <-ch2

		if !ok || v1 != v2 {
			return false
		}
	}

	return true
}

func main() {
	var result bool = Same(tree.New(1), tree.New(1))
	fmt.Println(result)
	result = Same(tree.New(1), tree.New(2))
	fmt.Println(result)
}
