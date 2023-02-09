package main

import "fmt"

type Node struct {
	element  interface{}
	next     *Node
	previous *Node
}

type DoublyLinkedLists struct {
	head *Node
}

func (l *DoublyLinkedLists) Find(element interface{}) *Node {
	current := l.head
	for current.element != element {
		current = current.next
	}
	return current
}

func (l *DoublyLinkedLists) FindPrevious(element interface{}) *Node {
	current := l.head
	for current.next != nil && current.next.element != element {
		current = current.next
	}
	return current
}

func (l *DoublyLinkedLists) Print() {
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

func (l *DoublyLinkedLists) Insert(newElement interface{}, element interface{}) {
	newNode := &Node{element: newElement}
	current := l.Find(element)

	newNode.next = current.next
	newNode.previous = current

	if current.next != nil {
		current.next.previous = newNode
	}

	current.next = newNode
}

func NewDoublyLinkedLists() *DoublyLinkedLists {
	return &DoublyLinkedLists{head: &Node{element: "head", next: nil}}
}

func (l *DoublyLinkedLists) FindLast() *Node {
	current := l.head
	for current.next != nil {
		current = current.next
	}
	return current
}

func (l *DoublyLinkedLists) PrintReverse() {
	current := l.FindLast()

	for {
		fmt.Printf("%v", current.element)

		if current.element == "head" {
			fmt.Println()
			break
		}

		current = current.previous

		fmt.Print(" -> ")
	}
}

func (l *DoublyLinkedLists) Remove(element interface{}) {
	current := l.Find(element)
	prevNode := current.previous
	nextNode := current.next

	nextNode.previous = prevNode
    prevNode.next = nextNode
	current = nil
}

func main() {
	l := NewDoublyLinkedLists()
	l.Insert("Hoàng Phúc International", "head")
	l.Insert("Ecko", "Hoàng Phúc International")
	l.Insert("Kappa", "Ecko")
	l.Insert("Staple", "Ecko")
	l.Print()
	l.PrintReverse()

	fmt.Println("\nRemove")
	l.Remove("Ecko")
	l.Print()
}