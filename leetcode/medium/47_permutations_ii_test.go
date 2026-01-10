package medium

import (
	"reflect"
	"sort"
	"testing"
)

func _47_sortPermutations(perms [][]int) {
	for _, p := range perms {
		// Do not sort inner arrays, they are permutations
		_ = p
	}
	sort.Slice(perms, func(i, j int) bool {
		n := len(perms[i])
		if n != len(perms[j]) {
			return n < len(perms[j])
		}
		for k := 0; k < n; k++ {
			if perms[i][k] != perms[j][k] {
				return perms[i][k] < perms[j][k]
			}
		}
		return false
	})
}

func Test_47_permuteUnique(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example 1",
			nums: []int{1, 1, 2},
			want: [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3}, {1, 3, 2}, {2, 1, 3},
				{2, 3, 1}, {3, 1, 2}, {3, 2, 1},
			},
		},
		{
			name: "All same",
			nums: []int{1, 1, 1},
			want: [][]int{
				{1, 1, 1},
			},
		},
		{
			name: "Two same",
			nums: []int{1, 1},
			want: [][]int{
				{1, 1},
			},
		},
		{
			name: "Single element",
			nums: []int{1},
			want: [][]int{
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := _47_permuteUnique(tt.nums)
			_47_sortPermutations(got)
			_47_sortPermutations(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_47_permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
