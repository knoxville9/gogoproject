package main

import (
	"fmt"
	"sync"
)

func main() {
	t := &TreeNode{Data: "A"}
	t.Left = &TreeNode{Data: "B"}
	t.Left.Left = &TreeNode{Data: "C"}

	t.Right = &TreeNode{Data: "F"}
	t.Right.Right = &TreeNode{Data: "z"}

	layer(t)
}

func layer(node *TreeNode) {
	if node == nil {
		return
	}
	q := new(Queue)
	q.add(node)
	for q.Len > 0 {
		remove := q.remove()
		fmt.Println(remove)

		if remove.Left != nil {
			q.add(remove.Left)

		}
		if remove.Right != nil {
			q.add(remove.Right)

		}
	}

}

func (q *Queue) add(node *TreeNode) {
	q.lock.Lock()
	defer q.lock.Unlock()
	l := new(Linknode)
	l.Value = node

	if q.root == nil {
		q.root = l
	} else {
		now := q.root
		for now.Next != nil {
			now = now.Next
		}
		now.Next = l

	}

	q.Len += 1
}

func (q *Queue) remove() *TreeNode {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.Len == 0 {
		panic("no element")
	}
	root := q.root

	q.root = root.Next

	q.Len -= 1
	return root.Value

}

type TreeNode struct {
	Left, Right *TreeNode
	Data        string
}

type Queue struct {
	root *Linknode
	Len  int
	lock sync.Mutex
}

type Linknode struct {
	Next  *Linknode
	Value *TreeNode
}
