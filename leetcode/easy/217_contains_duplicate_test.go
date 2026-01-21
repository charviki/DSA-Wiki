package easy

import "testing"

func Test_217_containsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"Example 1", []int{1, 2, 3, 1}, true},
		{"Example 2", []int{1, 2, 3, 4}, false},
		{"Example 3", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{"Empty array", []int{}, false},
		{"Single element", []int{1}, false},
		{"Two identical elements", []int{1, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _217_containsDuplicate(tt.nums); got != tt.want {
				t.Errorf("_217_containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
