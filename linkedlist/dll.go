package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

func (l *LinkedList) InsertAtStart(data int) {
	newNode := &Node{
		data: data,
	}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}

}
func (l *LinkedList) PrintLl() {
	current := l.tail
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.prev
	}
	fmt.Println("nil")
}

func main() {
	//make dll
	l := &LinkedList{}
	l.InsertAtStart(5)
	l.InsertAtStart(6)
	l.PrintLl()
}
