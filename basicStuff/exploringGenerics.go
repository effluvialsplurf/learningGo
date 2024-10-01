package main

import (
	"fmt"
	"strconv"
)

type Number interface {
	~int | ~float64
}

func double[T Number](num *T) {
	*num = *num + *num
}

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type PrintInt int

func (pi PrintInt) String() string {
	return strconv.Itoa(int(pi))
}

type PrintFloat float64

func (pf PrintFloat) String() string {
	return fmt.Sprintf("%f", pf)
}

func TryPrint[P Printable](p P) {
	fmt.Println(p)
}

// generic singly linked list
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type List[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *List[T]) PrintList() {
	n := l.Head
	count := 0
	for n != nil {
		fmt.Printf("Value: %v, Position: %v \n", n.Value, count)
		count++
		n = n.Next
	}
}

func (l *List[T]) Add(t T) {
	n := &Node[T]{
		Value: t,
	}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = l.Tail.Next
}

func (l *List[T]) Insert(t T, pos int) {
	newNode := &Node[T]{
		Value: t,
	}

	var prev *Node[T]
	var n *Node[T]

	for i := range pos + 1 {

		if i == 0 {
			n = l.Head
		}

		if i == pos-1 {
			prev = n
		}

		if i == pos {
			fmt.Printf("new: %v, prev: %v, curr: %v \n", newNode.Value, prev.Value, n.Value)
			newNode.Next = n
			prev.Next = newNode
		}

		n = n.Next
	}
}

func (l *List[T]) Index(t T) int {
	n := l.Head

	idx := 0
	for n != nil {
		if n.Value == t {
			return idx
		}
		n = n.Next
		idx++
	}

	return -1
}
