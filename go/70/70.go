package main

import (
	"fmt"
	"tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) (same bool) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	same = true

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for i := range ch1 {
		if i != <-ch2 {
			same = false
		}
	}

	return
}

func main() {
	ch := make(chan int)
	go func() {
		Walk(tree.New(1), ch)
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(1)))
}
