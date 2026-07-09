package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T comparable] struct {
	next *List[T]
	val  T
}

// Returns index where item was placed
func (l *List[T]) Push(v T) int {
	i := 0

	for l.next != nil {
		l = l.next
		i++
	}

	l.next = &List[T]{val: v}
	i++

	return i
}

func (l *List[T]) Prepend(v T) {
	clone := *l
	l.val = v
	l.next = &clone
}

func (l *List[T]) Print() {
	curr := l
	for curr != nil {
		fmt.Printf("%v", curr.val)

		if curr.next != nil {
			fmt.Printf(" -> ")
		}

		curr = curr.next
	}

	fmt.Print("\n")
}

func (l *List[T]) Len() int {
	i := 0
	curr := l
	for curr != nil {
		i++
		curr = curr.next
	}

	return i
}

// Contains(v T) bool
func (l *List[T]) Contains(v T) bool {
	curr := l
	for curr != nil {
		if curr.val == v {
			return true
		}
		curr = curr.next
	}
	return false
}

// Remove(v T)
func (l *List[T]) Remove(v T) {
	if l.val == v {
		if l.next != nil {
			*l = *l.next
		} else {
			var zero T
			l.val = zero
		}
		return
	}

	curr := l
	for curr.next != nil {
		if curr.next.val == v {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
}

func main() {
	var list List[int]
	list.val = 1
	list.Push(2)
	list.Prepend(0)
	list.Print()
	fmt.Println("Length:", list.Len())
	fmt.Println("Contains 2?", list.Contains(2))
	fmt.Println("Contains 3?", list.Contains(3))
	list.Remove(2)
	fmt.Println("Contains 2?", list.Contains(2))
	list.Push(3)
	list.Print()
}
