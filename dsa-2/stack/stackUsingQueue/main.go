package main

import (
	"errors"
	"fmt"
)

// Stack represents a stack using a single queue
type Stack struct {
	queue []int
}

// NewStack creates a new stack
func NewStack() *Stack {
	return &Stack{
		queue: make([]int, 0),
	}
}

// Push adds an element to the stack
func (s *Stack) Push(val int) {
	// Add the new element to the end of the queue
	s.queue = append(s.queue, val)

	// Rotate the queue to make the last element the front
	for i := 0; i < len(s.queue)-1; i++ {
		front := s.queue[0]
		s.queue = s.queue[1:]
		s.queue = append(s.queue, front)
	}
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() (int, error) {
	if len(s.queue) == 0 {
		return 0, errors.New("stack is empty")
	}

	// The front of the queue is the top of the stack
	front := s.queue[0]
	s.queue = s.queue[1:]
	return front, nil
}

// Peek returns the top element of the stack without removing it
func (s *Stack) Peek() (int, error) {
	if len(s.queue) == 0 {
		return 0, errors.New("stack is empty")
	}

	// The front of the queue is the top of the stack
	front := s.queue[0]
	return front, nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.queue) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.queue)
}

func main() {
	stack := NewStack()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	topElement, err := stack.Peek()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Top element:", topElement) // Output: Top element: 30
	}
	fmt.Println("Stack size:", stack.Size()) // Output: Stack size: 3

	poppedElement, err := stack.Pop()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Popped element:", poppedElement) // Output: Popped element: 30
	}
	topElement, err = stack.Peek()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Top element:", topElement) // Output: Top element: 20
	}
	fmt.Println("Stack size:", stack.Size()) // Output: Stack size: 2

	fmt.Println("Is stack empty?", stack.IsEmpty()) // Output: Is stack empty? false

	stack.Pop()
	stack.Pop()

	fmt.Println("Is stack empty?", stack.IsEmpty()) // Output: Is stack empty? true
}
