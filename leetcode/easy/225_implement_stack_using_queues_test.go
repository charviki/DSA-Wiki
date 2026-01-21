package easy

import (
	"testing"
)

func Test_225_MyStack(t *testing.T) {
	obj := _225_Constructor()

	// Test Push
	obj.Push(1)
	obj.Push(2)

	// Test Top
	if got := obj.Top(); got != 2 {
		t.Errorf("Top() = %v, want %v", got, 2)
	}

	// Test Pop
	if got := obj.Pop(); got != 2 {
		t.Errorf("Pop() = %v, want %v", got, 2)
	}

	// Test Empty
	if got := obj.Empty(); got != false {
		t.Errorf("Empty() = %v, want %v", got, false)
	}

	// Test Pop last element
	if got := obj.Pop(); got != 1 {
		t.Errorf("Pop() = %v, want %v", got, 1)
	}

	// Test Empty after clearing
	if got := obj.Empty(); got != true {
		t.Errorf("Empty() = %v, want %v", got, true)
	}
}

func Test_225_MyStack2(t *testing.T) {
	obj := _225_Constructor2()

	// Test Push
	obj.Push(1)
	obj.Push(2)

	// Test Top
	if got := obj.Top(); got != 2 {
		t.Errorf("Top() = %v, want %v", got, 2)
	}

	// Test Pop
	if got := obj.Pop(); got != 2 {
		t.Errorf("Pop() = %v, want %v", got, 2)
	}

	// Test Empty
	if got := obj.Empty(); got != false {
		t.Errorf("Empty() = %v, want %v", got, false)
	}

	// Test Pop last element
	if got := obj.Pop(); got != 1 {
		t.Errorf("Pop() = %v, want %v", got, 1)
	}

	// Test Empty after clearing
	if got := obj.Empty(); got != true {
		t.Errorf("Empty() = %v, want %v", got, true)
	}
}
