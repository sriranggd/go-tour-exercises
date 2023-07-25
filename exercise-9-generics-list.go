package main

import "fmt"

type List[T comparable] struct {
	next *List[T]
	val  T
}

func (l *List[T]) printAll() {
	current := l
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
}

func (l *List[T]) insertAfter(v T) *List[T] {
	newNode := List[T]{next: l.next, val: v}
	l.next = &newNode
	return l.next
}

func (l *List[T]) delete(v T) {
	current, previous := l.next, l
	found := false
	for ; current != nil; current = current.next {
		if current.val == v {
			previous.next = current.next
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Did not find %v in the list\n", v)
	}
}
