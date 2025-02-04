package main

// Import the errors package for error handling
import "errors"

// Stack struct represents a stack data structure implemented using a slice (queue)
// Even though we name it queue internally, this implementation behaves like a stack
type Stack struct {
	queue []int // Using a slice to store integer values
}

// NewStack creates and initializes a new empty stack
// Returns a pointer to the newly created Stack
func NewStack() *Stack {
	return &Stack{
		queue: make([]int, 0), // Initialize an empty slice with 0 capacity
	}
}

// Push adds a new element to the stack
// Example: If we Push(5) to empty stack, queue becomes [5]
// If we then Push(3), it first becomes [5,3] and then rotates to [3,5]
func (s *Stack) Push(val int) {
	// First append the new value to the end
	s.queue = append(s.queue, val)

	// Rotate all elements except the last one
	// This ensures the newest element ends up at the front (top of stack)
	for i := 0; i < len(s.queue)-1; i++ {
		front := s.queue[0]              // Store the front element
		s.queue = s.queue[1:]            // Remove the front element
		s.queue = append(s.queue, front) // Add it to the back
	}
}

// Pop removes and returns the top element from the stack
// Returns an error if the stack is empty
// Example: If stack is [3,5], Pop() returns 3 and stack becomes [5]
func (s *Stack) Pop() (int, error) {
	if len(s.queue) == 0 {
		return 0, errors.New("stack is empty")
	}
	front := s.queue[0]   // Get the front element
	s.queue = s.queue[1:] // Remove the front element
	return front, nil
}

// Peek returns the top element of the stack without removing it
// Returns an error if the stack is empty
// Example: If stack is [3,5], Peek() returns 3 but stack remains [3,5]
func (s *Stack) Peek() (int, error) {
	if len(s.queue) == 0 {
		return 0, errors.New("stack is empty")
	}
	front := s.queue[0] // Return the front element without modifying the queue
	return front, nil
}

// IsEmpty checks if the stack is empty
// Returns true if stack has no elements, false otherwise
func (s *Stack) IsEmpty() bool {
	return len(s.queue) == 0
}

// Size returns the current number of elements in the stack
// Example: If stack is [3,5], Size() returns 2
func (s *Stack) Size() int {
	return len(s.queue)
}
