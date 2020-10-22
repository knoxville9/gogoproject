package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"strconv"
	"sync"
)

func main() {
	a := new(ArrayQueue)

	group := errgroup.Group{}
	for i := 0; i < 1; i++ {
		value := i
		itoa := strconv.Itoa(value)

		group.Go(func() error {
			a.push(itoa)
			return nil
		})

	}

	group.Wait()
	a.pop()
	fmt.Println(a)

}

type Linknode struct {
	left, right *Linknode
	value       int
}

func (l *Linknode) Left() *Linknode {
	panic("implement me")
}

func (l *Linknode) Right() *Linknode {
	panic("implement me")
}

func (l *Linknode) init() {
	if l == nil {
		l.left = l
		l.right = l
	}
}

type circle interface {
	init()
	Left() *Linknode
	Right() *Linknode
	len() int
	Add()
	get(n int) *Linknode
}

type ArrayQueue struct {
	array []string
	size  int
	lock  sync.Mutex
}

func (a *ArrayQueue) push(st string) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.array = append(a.array, st)
	a.size += 1

}

func (a *ArrayQueue) pop() string {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.size == 0 {
		panic("size 0 ")
	}
	first := a.array[0]

	newarray := make([]string, a.size-1, a.size-1)

	for i := 0; i < a.size-1; i++ {
		newarray[i] = a.array[i+1]

	}
	a.array = newarray
	a.size -= 1

	return first

}

type Queue interface {
	pop() string
	push(st string)
}
