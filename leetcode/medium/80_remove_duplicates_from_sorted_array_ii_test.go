package medium

import (
	"reflect"
	"testing"
)

func Test_80_removeDuplicates(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		want         int
		expectedNums []int
	}{
		{
			name:         "Example 1",
			nums:         []int{1, 1, 1, 2, 2, 3},
			want:         5,
			expectedNums: []int{1, 1, 2, 2, 3},
		},
		{
			name:         "Example 2",
			nums:         []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			want:         7,
			expectedNums: []int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			name:         "No duplicates",
			nums:         []int{1, 2, 3},
			want:         3,
			expectedNums: []int{1, 2, 3},
		},
		{
			name:         "All same",
			nums:         []int{1, 1, 1, 1},
			want:         2,
			expectedNums: []int{1, 1},
		},
		{
			name:         "Empty array",
			nums:         []int{},
			want:         0,
			expectedNums: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			got := _80_removeDuplicates(numsCopy)
			if got != tt.want {
				t.Errorf("_80_removeDuplicates() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(numsCopy[:got], tt.expectedNums) {
				t.Errorf("Array content = %v, want %v", numsCopy[:got], tt.expectedNums)
			}
		})
	}
}

func Test_80_removeDuplicates2(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		want         int
		expectedNums []int
	}{
		{
			name:         "Example 1",
			nums:         []int{1, 1, 1, 2, 2, 3},
			want:         5,
			expectedNums: []int{1, 1, 2, 2, 3},
		},
		{
			name:         "Example 2",
			nums:         []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			want:         7,
			expectedNums: []int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			name:         "No duplicates",
			nums:         []int{1, 2, 3},
			want:         3,
			expectedNums: []int{1, 2, 3},
		},
		{
			name:         "All same",
			nums:         []int{1, 1, 1, 1},
			want:         2,
			expectedNums: []int{1, 1},
		},
		{
			name:         "Empty array",
			nums:         []int{},
			want:         0,
			expectedNums: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			got := _80_removeDuplicates2(numsCopy)
			if got != tt.want {
				t.Errorf("_80_removeDuplicates2() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(numsCopy[:got], tt.expectedNums) {
				t.Errorf("Array content = %v, want %v", numsCopy[:got], tt.expectedNums)
			}
		})
	}
}
