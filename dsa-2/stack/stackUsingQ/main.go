package stack

import "errors"

// Stack struct represents a stack data structure implemented using a slice.
// A stack is a Last-In-First-Out (LIFO) data structure, meaning the last element
// added is the first one to be removed - like a stack of plates.
//
// Example of stack operations:
// 1. Empty stack: []
// 2. Push(5): [5]
// 3. Push(3): [3,5]  (3 is on top because it was added last)
// 4. Pop(): returns 3, stack becomes [5]
type Stack struct {
	// queue is a slice of integers that stores our stack elements
	// We call it queue because of how we're implementing the stack internally,
	// but from the outside, it behaves exactly like a stack
	queue []int
}

// NewStack creates and initializes a new empty stack.
// This is a constructor function that sets up the initial state of our stack.
//
// Returns:
//   - *Stack: a pointer to the newly created Stack structure
//
// Example usage:
//   myStack := NewStack()  // Creates a new empty stack
func NewStack() *Stack {
	return &Stack{
		// make([]int, 0) creates a new slice with initial length of 0
		// This means we start with an empty stack that can grow as needed
		queue: make([]int, 0),
	}
}

// Push adds a new element to the top of the stack.
// This method uses an interesting approach where we add the element to the end
// and then rotate the elements to maintain our stack properties.
//
// Parameters:
//   - val: the integer value to add to the stack
//
// Example process for Push(4) on stack [3,5]:
// 1. Append 4: [3,5,4]
// 2. Rotate:
//    - First iteration: [5,4,3]
//    - Second iteration: [4,3,5]
// Final result: [4,3,5] (4 is now at the top of the stack)
func (s *Stack) Push(val int) {
	// First, append the new value to the end of our slice
	s.queue = append(s.queue, val)

	// Now we need to rotate all elements except the last one we just added
	// This rotation ensures the newest element (val) ends up at the front
	for i := 0; i < len(s.queue)-1; i++ {
		// Store the front element before we remove it
		front := s.queue[0]

		// Remove the front element by slicing from index 1 onwards
		// Example: [3,5,4] becomes [5,4]
		s.queue = s.queue[1:]

		// Add the front element to the back
		// Example: [5,4] becomes [5,4,3]
		s.queue = append(s.queue, front)
	}
}

// Pop removes and returns the top element from the stack.
// In a stack, "top" refers to the most recently added element.
//
// Returns:
//   - int: the value at the top of the stack
//   - error: an error if the stack is empty, nil otherwise
//
// Example:
//   stack: [3,5] (3 is at the top)
//   value, err := stack.Pop()
//   // value = 3, err = nil
//   // stack is now [5]
func (s *Stack) Pop() (int, error) {
	// First, check if the stack is empty
	// If it is, we can't pop anything, so return an error
	if len(s.queue) == 0 {
		return 0, errors.New("stack is empty")
	}

	// Get the front element (top of the stack)
	front := s.queue[0]

	// Remove the front element by slicing from index 1 onwards
	// Example: [3,5] becomes [5]
	s.queue = s.queue[1:]

	// Return the element we removed and nil for the error
	return front, nil
}

// Peek returns the top element of the stack without removing it.
// This is useful when you want to see what's on top of the stack
// without actually removing it.
//
// Returns:
//   - int: the value at the top of the stack
//   - error: an error if the stack is empty, nil otherwise
//
// Example:
//   stack: [3,5]
//   value, err := stack.Peek()
//   // value = 3, err = nil
//   // stack remains [3,5]
func (s *Stack) Peek() (int, error) {
	// Check if the stack is empty
	if len(s.queue) == 0 {
		return 0, errors.New("stack is empty")
	}

	// Return the front element without modifying the queue
	// We don't need to slice the queue because we're just looking
	front := s.queue[0]
	return front, nil
}

// IsEmpty checks if the stack has any elements.
// This is a helper method that's often used before performing operations
// to avoid errors.
//
// Returns:
//   - bool: true if the stack is empty, false otherwise
//
// Example:
//   stack: []
//   isEmpty := stack.IsEmpty()  // returns true
//   stack: [3,5]
//   isEmpty := stack.IsEmpty()  // returns false
func (s *Stack) IsEmpty() bool {
	// len(s.queue) returns the number of elements in the slice
	// If it's 0, the stack is empty
	return len(s.queue) == 0
}

// Size returns the current number of elements in the stack.
// This is useful when you need to know how many elements are in the stack
// without modifying it.
//
// Returns:
//   - int: the number of elements currently in the stack
//
// Example:
//   stack: [3,5,7]
//   size := stack.Size()  // returns 3
func (s *Stack) Size() int {
	// len(s.queue) returns the number of elements in the slice
	return len(s.queue)
}
