package easy

import "testing"

func Test_219_containsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{"Example 1", []int{1, 2, 3, 1}, 3, true},
		{"Example 2", []int{1, 0, 1, 1}, 1, true},
		{"Example 3", []int{1, 2, 3, 1, 2, 3}, 2, false},
		{"No duplicate", []int{1, 2, 3, 4}, 3, false},
		{"Duplicate within k", []int{1, 2, 1}, 2, true},
		{"Duplicate exactly k apart", []int{1, 2, 3, 1}, 3, true},
		{"Duplicate more than k apart", []int{1, 2, 3, 4, 1}, 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _219_containsNearbyDuplicate(tt.nums, tt.k); got != tt.want {
				t.Errorf("_219_containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
