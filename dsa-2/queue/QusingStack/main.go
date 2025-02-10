package main

import "errors"

// Package provides a Queue implementation using two stacks
// This is an efficient way to implement a queue using two stacks
// where one stack is used for enqueueing (stackInput) and another
// for dequeueing (stackOutput)

// Queue struct defines our queue data structure
// It uses two slices (which act as stacks) to implement queue behavior
type Queue struct {
	// stackInput is used to store new elements as they are added
	// Example: If we enqueue 1,2,3 they will be stored as [1,2,3]
	stackInput []int

	// stackOutput is used when we need to remove elements
	// When dequeuing, if this is empty, we flip stackInput into this
	// Example: [1,2,3] becomes [3,2,1] when moved to stackOutput
	stackOutput []int
}

// NewQeueue creates and initializes a new Queue
// Returns a pointer to the new Queue structure
func NewQeueue() *Queue {
	return &Queue{
		// make([]int, 0) creates empty slices with initial capacity of 0
		// These slices will grow automatically as we add elements
		stackInput:  make([]int, 0),
		stackOutput: make([]int, 0),
	}
}

// Enqueue adds a new value to the back of the queue
// Example: Enqueue(1), Enqueue(2), Enqueue(3) results in [1,2,3]
func (q *Queue) Enqueue(value int) {
	// Simply append the new value to stackInput
	// This is an O(1) operation most of the time
	q.stackInput = append(q.stackInput, value)
}

// Dequeue removes and returns the front element of the queue
// Returns an error if the queue is empty
func (q *Queue) Dequeue() (int, error) {
	// If stackOutput is empty, we need to transfer elements from stackInput
	if len(q.stackOutput) == 0 {
		// Transfer all elements from stackInput to stackOutput
		// This reverses their order, which is exactly what we want
		// Example: stackInput [1,2,3] becomes stackOutput [3,2,1]
		for len(q.stackInput) > 0 {
			// Get the last element from stackInput
			top := q.stackInput[len(q.stackInput)-1]
			// Remove it from stackInput (slice up to last element)
			q.stackInput = q.stackInput[:len(q.stackInput)-1]
			// Add it to stackOutput
			q.stackOutput = append(q.stackOutput, top)
		}
	}

	// If after transfer stackOutput is still empty, queue is empty
	if len(q.stackOutput) == 0 {
		return 0, errors.New("queue is empty")
	}

	// Get the last element from stackOutput (which was the first one in)
	front := q.stackOutput[len(q.stackOutput)-1]
	// Remove it from stackOutput
	q.stackOutput = q.stackOutput[:len(q.stackOutput)-1]
	// Return the element
	return front, nil
}

// Peek returns the front element of the queue without removing it
// Returns an error if the queue is empty
func (q *Queue) Peek() (int, error) {
	// Similar to Dequeue, if stackOutput is empty, transfer elements
	if len(q.stackOutput) == 0 {
		// Transfer process is identical to Dequeue
		for len(q.stackInput) > 0 {
			top := q.stackInput[len(q.stackInput)-1]
			q.stackInput = q.stackInput[:len(q.stackInput)-1]
			q.stackOutput = append(q.stackOutput, top)
		}
	}

	// If queue is empty, return error
	if len(q.stackOutput) == 0 {
		return 0, errors.New("queue is empty")
	}

	// Return the front element without removing it
	front := q.stackOutput[len(q.stackOutput)-1]
	return front, nil
}

// IsEmpty checks if the queue has any elements
// Returns true if the queue is empty, false otherwise
func (q *Queue) IsEmpty() bool {
	// Queue is empty if both stacks are empty
	return len(q.stackInput) == 0 && len(q.stackOutput) == 0
}
