package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	data []interface{}
	top  int
}

func (s *Stack) Push(element interface{}) {
	s.top++
	s.data = append(s.data, element)
}

func (s *Stack) Pop() interface{} {
	if len(s.data) > 0 {
		s.top--
		last := s.data[s.top]
		s.data = s.data[:s.top]

		return last
	}

	return nil
}

func (s *Stack) Peek() interface{} {
	if len(s.data) > 0 {
		return s.data[s.top-1]
	}

	return nil
}

func (s *Stack) Clear() {
	s.top = 0
	s.data = []interface{}{}
}

func (s *Stack) Length() int {
	return s.top
}

//chuyển một số từ dạng thập phân sang dạng nhị phân
func MulBase(number int, base int) string {
	s := Stack{}
	for {
		s.Push(number % base)
		number = number / base

		if number <= 0 {
			break
		}
	}

	converted := ""

	for s.Length() > 0 {
		converted = converted + strconv.Itoa(s.Pop().(int))
	}

	return converted
}

func main() {
	/* stack := Stack{}
	stack.Push("Hoàng Phúc International")
	stack.Push("Kappa")
	stack.Push("Ecko Unltd")
	stack.Push("Superga")
	stack.Push("Staple")

	fmt.Printf("length %d\n", stack.Length())
	fmt.Println()

	fmt.Printf("length %s\n", stack.Pop())
	fmt.Printf("length after pop %d\n", stack.Length())
	fmt.Println()

	fmt.Printf("length %s\n", stack.Peek())
	fmt.Printf("length after peek %d\n", stack.Length())
	fmt.Println()

	stack.Clear()
	fmt.Printf("length %d\n", stack.Length()) */

	var number = 100

	fmt.Printf("Number %d converted to base %d is %s\n", number, 2, MulBase(number, 2))
	fmt.Printf("Number %d converted to base %d is %s\n", number, 8, MulBase(number, 8))
}
