package easy

import (
	"testing"
)

func Test_232_MyQueue(t *testing.T) {
	obj := _232_Constructor()

	// Test Push
	obj.Push(1)
	obj.Push(2)

	// Test Peek
	if got := obj.Peek(); got != 1 {
		t.Errorf("Peek() = %v, want %v", got, 1)
	}

	// Test Pop
	if got := obj.Pop(); got != 1 {
		t.Errorf("Pop() = %v, want %v", got, 1)
	}

	// Test Empty
	if got := obj.Empty(); got != false {
		t.Errorf("Empty() = %v, want %v", got, false)
	}

	obj.Push(3)

	// Test Pop last element
	if got := obj.Pop(); got != 2 {
		t.Errorf("Pop() = %v, want %v", got, 2)
	}

	if got := obj.Pop(); got != 3 {
		t.Errorf("Pop() = %v, want %v", got, 2)
	}

	// Test Empty after clearing
	if got := obj.Empty(); got != true {
		t.Errorf("Empty() = %v, want %v", got, true)
	}
}
