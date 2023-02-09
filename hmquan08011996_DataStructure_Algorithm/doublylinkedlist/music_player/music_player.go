package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	element  interface{}
	next     *Node
	previous *Node
}

type DoublyLinkedLists struct {
	head    *Node
	current *Node
}

func (l *DoublyLinkedLists) Find(element interface{}) *Node {
	current := l.head
	for current.element != element {
		current = current.next
	}
	return current
}

func (l *DoublyLinkedLists) FindLast() *Node {
	current := l.head
	for current.next != nil {
		current = current.next
	}
	return current
}

func (l *DoublyLinkedLists) Insert(newElement interface{}, element interface{}) {
	newNode := &Node{element: newElement}
	current := l.Find(element)
	newNode.next = current.next
	newNode.previous = current
	current.next = newNode
}

func (l *DoublyLinkedLists) Remove(element interface{}) {
	current := l.Find(element)
	prevNode := current.previous
	nextNode := current.next

	prevNode.next = nextNode
	nextNode.previous = prevNode
	current = nil
}

func (l *DoublyLinkedLists) Print() {
	current := l.head

	for {
		if current.element == "head" {
			current = current.next
			continue
		}

		fmt.Printf("%v", current.element)

		if current.next == nil {
			fmt.Println()
			break
		}

		current = current.next

		fmt.Print(" -> ")
	}
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

func (l *DoublyLinkedLists) Advance(step int) {
	for i := 0; i < step; i++ {
		if l.current.next != nil {
			l.current = l.current.next
		}
	}
}

func (l *DoublyLinkedLists) Back(step int) {
	for i := 0; i < step; i++ {
		if l.current.previous != nil && l.current.previous.element != "head" {
			l.current = l.current.previous

		}
	}
}

func (l *DoublyLinkedLists) Play() {
	if l.current.element == "head" {
		l.Advance(1)
	}

	fmt.Printf("Play music %+v\n", l.current.element)
}

var music = NewDoublyLinkedLists()

func init() {
	music.Insert("Mùa xuân về trên Thành Phố Hồ Chí Minh", "head")
	music.Insert("Thương quá Việt Nam", "Mùa xuân về trên Thành Phố Hồ Chí Minh")
	music.Insert("Năm anh em trên chiếc xe tăng", "Thương quá Việt Nam")
	music.Insert("Tiến Quân Ca", "Năm anh em trên chiếc xe tăng")
}

func main() {
	music.Print()
	fmt.Println()
	music.Play()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Printf("Enter prev or next to play prev or next music: ")

		scanner.Scan()
		input := scanner.Text()

		if input == "prev" {
			music.Back(1)
			music.Play()
			continue
		}

		if input == "next" {
			music.Advance(1)
			music.Play()
			continue
		}
	}

}

func NewDoublyLinkedLists() *DoublyLinkedLists {
	head := &Node{element: "head", next: nil, previous: nil}
	return &DoublyLinkedLists{head: head, current: head}
}
