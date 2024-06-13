package day20

import (
	"fmt"
	"strconv"
)

type node struct {
	value int
	moved int
	id    int
	next  *node
	prev  *node
}

type linkedlist struct {
	head *node
	tail *node
	size int
}

func NewLinkedList() *linkedlist {
	return &linkedlist{}
}

func (l *linkedlist) Add(n int) {
	nd := node{
		value: n,
		prev:  l.tail,
		id:    l.size,
	}

	if l.head == nil {
		l.head = &nd
		l.tail = &nd
	} else {
		l.tail.next = &nd
		l.tail = &nd
	}

	l.size++
}

func (l *linkedlist) Get(index int) *node {
	result := l.head
	for i := 0; i < index; i++ {
		result = result.next
		if result == nil {
			result = l.head
		}
	}
	return result
}

func (l *linkedlist) Mix(value, moved int) {
	i, n := l.ScanFor(value, moved)
	var tailflag bool
	if n == nil {
		fmt.Println("failed to find node in scan for:", value, moved)
	}
	if (i + 1) == l.size {
		tailflag = true
	}
	n.moved += 1
	if n.value == 0 {
		return
	}
	if i == 0 {
		l.head = n.next
	}

	before := n.prev
	after := n.next
	if before != nil {
		before.next = after
	}
	if after != nil {
		after.prev = before
	}

	pointer := n
	if n.value < 0 { // move backwards
		interations := n.value * -1
		for i := 0; i < interations; i++ {
			pointer = pointer.prev
			if pointer == nil {
				pointer = l.tail
			}
			if i%811589153 == 0 {
				fmt.Println(i / 811589153)
			}
		}

		newBefore := pointer.prev
		newAfter := pointer
		n.prev = newBefore
		n.next = newAfter
		if newBefore != nil {
			newBefore.next = n
		} else {
			n.prev = l.tail
			l.tail.next = n
			l.tail = n
			n.next = nil
			if tailflag {
				l.ResetTail()
			}
			return
		}
		if newAfter != nil {
			newAfter.prev = n
		} else {
			l.tail = n
		}
	} else { // move forward
		interations := n.value
		for i := 0; i < interations; i++ {
			pointer = pointer.next
			if pointer == nil {
				pointer = l.head
			}
		}
		newBefore := pointer
		newAfter := pointer.next
		n.prev = newBefore
		n.next = newAfter

		if newAfter != nil {
			newBefore.next = n
			newAfter.prev = n
		} else {
			n.prev = l.tail
			l.tail.next = n
			l.tail = n
			if tailflag {
				l.ResetTail()
			}
			return
		}
	}
	if tailflag {
		l.ResetTail()
	}
}

func (l *linkedlist) ModMix(id int) {
	//i, n := l.ScanFor(value, moved)
	i, n := l.ScanForId(id)
	if n == nil {
		// fmt.Println("failed to find node in scan for:", value, moved)
		fmt.Println("failed to find node in scan for ID:", id)
	}
	// if n.value != 0 && n.value%(l.size-1) == 0 {
	// 	l.Mix(value, moved)
	// 	return
	// }
	n.moved += 1
	if n.value == 0 || n.value%(l.size-1) == 0 {
		return
	}
	if i == 0 {
		l.head = n.next
	}

	before := n.prev
	after := n.next
	if before != nil {
		before.next = after
	}
	if after != nil {
		after.prev = before
	}

	if (i + 1) == l.size {
		l.ResetTail()
	}
	pointer := n
	if n.value < 0 { // move backwards
		interations := (n.value * -1) % (l.size - 1)
		if interations == 0 {
			interations = l.size - 1
		}
		for i := 0; i < interations; i++ {
			pointer = pointer.prev
			if pointer == nil {
				pointer = l.tail
			}
		}
		// if pointer == n {
		// 	n.prev.next = n
		// 	n.next.prev = n
		// 	return
		// }
		newBefore := pointer.prev
		newAfter := pointer
		n.prev = newBefore
		n.next = newAfter
		if newBefore != nil {
			newBefore.next = n
		} else {
			n.prev = l.tail
			l.tail.next = n
			l.tail = n
			n.next = nil
			return
		}
		if newAfter != nil {
			newAfter.prev = n
		} else {
			l.tail = n
		}
	} else { // move forward
		interations := n.value % (l.size - 1)
		if interations == 0 {
			interations = l.size - 1
		}
		for i := 0; i < interations; i++ {
			pointer = pointer.next
			if pointer == nil {
				pointer = l.head
			}
		}
		// if pointer == n {
		// 	n.prev.next = n
		// 	n.next.prev = n
		// 	return
		// }
		newBefore := pointer
		newAfter := pointer.next
		n.prev = newBefore
		n.next = newAfter

		if newAfter != nil {
			newBefore.next = n
			newAfter.prev = n
		} else {
			n.prev = l.tail
			l.tail.next = n
			l.tail = n
			return
		}
	}
}

func (l *linkedlist) ScanFor(val int, moved int) (int, *node) {
	result := l.head
	index := 0
	for ; ; index++ {
		if result.value == val && result.moved == moved {
			return index, result
		} else {
			result = result.next
			if result == nil {
				return -1, nil
			}
		}
		if index > l.size {
			panic("cycle!")
		}
	}
}

func (l *linkedlist) ScanForId(id int) (int, *node) {
	result := l.head
	index := 0
	for ; ; index++ {
		if result.id == id {
			return index, result
		} else {
			result = result.next
			if result == nil {
				return -1, nil
			}
		}
		if index > l.size {
			panic("cycle!")
		}
	}
}

func (l *linkedlist) Print() {
	nd := l.head
	for index := 0; nd != nil; index++ {
		s := strconv.Itoa(nd.value) + "(" + strconv.Itoa(nd.moved) + ")"
		fmt.Print(s)
		fmt.Print(" ")
		nd = nd.next
		if index > l.size+1 {
			break
		}
	}
	fmt.Println()
}

func (l *linkedlist) ResetTail() {
	nd := l.head
	for index := 0; ; index++ {
		if nd.next == nil {
			l.tail = nd
			break
		}
		nd = nd.next
		if index > l.size+1 {
			panic("cycle!")
		}
	}
}
