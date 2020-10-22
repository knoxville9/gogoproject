package main

import (
	"fmt"
	"sync"
)

func main() {
	d := new(Doublelist)
	d.AddFromTail(0, 1)
	d.AddFromTail(0, 2)
	d.AddFromTail(2, 3)

	fmt.Println(d)
}

type Doublelist struct {
	Head *Listnode
	Tail *Listnode
	Size int
	Lock sync.Mutex
}

type Listnode struct {
	Prev  *Listnode
	Next  *Listnode
	Value int
}

func (l *Listnode) isNil() bool {
	return l == nil
}

func (l *Listnode) name() int {
	return l.Value

}

func (l *Listnode) GetNext() *Listnode {
	return l.Next
}
func (l *Listnode) GetPrev() *Listnode {
	return l.Prev
}
func (l *Listnode) HasPrev() bool {
	return l.Prev != nil
}
func (l *Listnode) HasNext() bool {
	return l.Next != nil
}

func (d *Doublelist) AddFromHead(n int, value int) {
	d.Lock.Lock()
	defer d.Lock.Unlock()
	if n > d.Size {
		panic("number over index ")
	}

	node := d.Head

	for i := 1; i < n; i++ {
		node = node.Next

	}

	newnode := new(Listnode)
	newnode.Value = value

	if node.isNil() {
		d.Head = newnode
		d.Tail = newnode

	} else {
		prev := node.Prev

		if prev.isNil() {
			node.Prev = newnode
			newnode.Next = node
			d.Head = newnode

		} else {
			prev.Next = newnode
			newnode.Prev = prev
			newnode.Next = node
			node.Prev = newnode

		}

	}
	d.Size += 1

}
func (d *Doublelist) AddFromTail(n int, value int) {
	d.Lock.Lock()
	defer d.Lock.Unlock()
	if n > d.Size {
		panic("number over index ")
	}

	node := d.Tail

	for i := 1; i < n; i++ {
		node = node.Prev

	}

	newnode := new(Listnode)
	newnode.Value = value

	if node.isNil() {
		d.Head = newnode
		d.Tail = newnode

	} else {
		next := node.Next

		if next.isNil() {
			node.Next = newnode
			newnode.Prev = node
			d.Tail = newnode

		} else {
			next.Prev = newnode
			newnode.Next = next
			node.Next = newnode
			newnode.Prev = node

		}

	}
	d.Size += 1

}
