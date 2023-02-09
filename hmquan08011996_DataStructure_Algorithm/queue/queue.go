package main

import "fmt"

type Queue struct {
	data []interface{}
}

func (q *Queue) Enqueue(element interface{}) {
	q.data = append([]interface{}{element}, q.data...)
}

func (q *Queue) Dequeue() interface{} {
	length := len(q.data)

	if length > 0 {
		front := q.data[length-1]
		q.data = q.data[:length-1]

		return front
	}

	return nil
}

func (q *Queue) Front() interface{} {
	length := len(q.data)

	if length > 0 {
		return q.data[length-1]
	}

	return nil
}

func (q *Queue) Back() interface{} {
	length := len(q.data)

	if length > 0 {
		return q.data[0]
	}

	return nil
}

func (q *Queue) Length() int {
	return len(q.data)
}

func (q *Queue) Clear() {
	q.data = []interface{}{}
}

func (q *Queue) Print() {
	fmt.Printf("Queue %+v\n", q.data)
}

func main() {
	queue := Queue{}
	queue.Enqueue("1: Hoàng Phúc International")
	queue.Enqueue("2: Kappa")
	queue.Enqueue("3: Ecko Unltd")
	queue.Enqueue("4: Superga")
	queue.Enqueue("5: Staple")

	queue.Print()
	fmt.Println()

	fmt.Println(queue.Dequeue())
	queue.Print()
	fmt.Println()

	fmt.Printf("Front %v\n", queue.Front())
	fmt.Printf("Back %v\n", queue.Back())
	queue.Print()
	fmt.Println()

	queue.Clear()
	fmt.Println(queue.Length())
	queue.Print()
}
