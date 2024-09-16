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
