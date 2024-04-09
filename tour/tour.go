package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}
func Walking(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	defer close(ch) // close channel after Walk() finishes
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	x := make(chan int) // define separate channels for both trees
	y := make(chan int)

	go Walking(t1, x) // call the Walking func for both trees
	go Walking(t2, y)

	for {
		v1, ok1 := <-x // receive from channel x and y
		v2, ok2 := <-y // the ok param tells us if channel is closed

		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("Yes!")
	} else {
		fmt.Println("No!")
	}

	if Same(tree.New(1), tree.New(2)) {
		fmt.Println("Yes!")
	} else {
		fmt.Println("No!")
	}

}
