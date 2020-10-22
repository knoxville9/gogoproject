package main

import (
	"fmt"
	"sync"
)

func main() {
	tree := Treenode{value: "A"}
	tree.left = &Treenode{value: "b"}

	layer(&tree)

}

func layer(tree *Treenode) {
	if tree == nil {
		panic("nil")
	}
	q := new(Queue)
	q.add(tree)
	for q.len > 0 {
		fmt.Println(tree)
		q.remove()

		if tree.left != nil {
			q.add(tree.left)
		} else if tree.rifgt != nil {
			q.add(tree.rifgt)

		}

	}

}
func (q *Queue) add(node *Treenode) {
	newnode := new(Linknode)

	if q.root == nil {
		q.root = newnode
		q.root.value = node

	} else {

		root := q.root
		if root.next != nil {
			root = root.next
		}
		root.next = newnode
		newnode.value = node
	}

	q.len += 1
}
func (q *Queue) remove() {

	q.len -= 1
}

type Treenode struct {
	left, rifgt *Treenode
	value       interface{}
}

type Queue struct {
	root *Linknode
	len  int
	lock sync.Mutex
}

type Linknode struct {
	next  *Linknode
	value *Treenode
}
