package problems

import "fmt"

type Node struct {
	data int
	next *Node
}

type MyLinkedList struct {
	head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if this.head == nil {
		return -1
	}
	if index < 0 {
		return -1
	}
	current := this.head

	for u := 0; u < index; u++ {
		if current.next == nil {
			return -1
		}
		current = current.next
	}
	return current.data
}

func (this *MyLinkedList) AddAtHead(val int) {
	data := &Node{data: val}
	if this.head != nil {
		data.next = this.head
	}
	this.head = data
}

func (this *MyLinkedList) AddAtTail(val int) {
	data := &Node{data: val}
	if this.head == nil {
		this.head = data
		return
	}
	current := this.head
	for current.next != nil {
		current = current.next
	}
	current.next = data
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	data := &Node{data: val}
	if this.head == nil {
		if index == 0 {
			this.head = data
		}
		return
	}
	if index < 0 {
		return
	}
	current := this.head

	for u := 1; u < index; u++ {
		if current.next == nil {
			return
		}
		current = current.next
	}
	data.next = current.next
	current.next = data
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.head == nil {
		return
	}
	if index < 0 {
		return
	}
	if index == 0 {
		this.head = this.head.next
		return
	}
	current := this.head

	for u := 1; u < index; u++ {
		if current.next == nil {
			return
		}
		current = current.next
	}
	if current.next == nil {
		return
	}
	current.next = current.next.next
	return
}

func (this *MyLinkedList) Travel() {
	current := this.head
	if this.head != nil {

		fmt.Print(current.data, ",")
		for current.next != nil {
			current = current.next
			fmt.Print(current.data, ",")
		}
	}
	fmt.Println("")
}
