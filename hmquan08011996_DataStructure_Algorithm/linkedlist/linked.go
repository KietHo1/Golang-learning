package main

import "fmt"		

type Node struct {
	element interface{}
	next    *Node
}

type LinkedLists struct {
	head *Node
}

func NewLinkedLists() *LinkedLists {
	return &LinkedLists{head: &Node{element: "head", next: nil}}
}

func (l *LinkedLists) Find(element interface{}) *Node {
	current := l.head
	for current.element != element {
		current = current.next
	}
	return current
}

func (l *LinkedLists) Insert(newElement interface{}, element interface{}) {
	newNode := &Node{element: newElement}
	current := l.Find(element)
	newNode.next = current.next
	current.next = newNode
}

func (l *LinkedLists) Print() {
	current := l.head

	for {
		fmt.Printf("%v", current.element)

		if current.next == nil {
			fmt.Println()
			break
		}
        
        current = current.next

		fmt.Print(" -> ")
	}
}

func (l *LinkedLists) FindPrevious(element interface{}) *Node {
	current := l.head
	for current.next != nil && current.next.element != element {
		current = current.next
	}
	return current
}

func (l *LinkedLists) Remove(element interface{}) {
	prevNode := l.FindPrevious(element)
	if prevNode.next != nil {
		current := prevNode.next
		prevNode.next = current.next
		current = nil
	}
}

func main() {
	l := NewLinkedLists()
	l.Insert("Hoàng Phúc International", "head")
	l.Insert("Ecko", "Hoàng Phúc International")
	l.Insert("Kappa", "Ecko")
	l.Print()	
	l.Remove("Ecko")
	l.Print()
}