package medium

import (
	"reflect"
	"sort"
	"testing"
)

func Test_47_permuteUnique(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{}, [][]int{{}}},
		{[]int{1}, [][]int{{1}}},
		{[]int{1, 2}, [][]int{{1, 2}, {2, 1}}},
		{[]int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		{[]int{1, 1, 1}, [][]int{{1, 1, 1}}},
	}

	for _, test := range tests {
		result := _47_permuteUnique(test.nums)
		if !reflect.DeepEqual(_47_sortAndRemoveDuplicates(result), _47_sortAndRemoveDuplicates(test.expected)) {
			t.Errorf("permuteUnique(%v) = %v, want %v", test.nums, result, test.expected)
		}
	}
}

// sortAndRemoveDuplicates sorts the 2D slice and removes duplicate entries.
func _47_sortAndRemoveDuplicates(perms [][]int) [][]int {
	sort.Slice(perms, func(i, j int) bool {
		for k := range perms[i] {
			if perms[i][k] != perms[j][k] {
				return perms[i][k] < perms[j][k]
			}
		}
		return false
	})
	uniquePerms := [][]int{perms[0]}
	for i := 1; i < len(perms); i++ {
		if !reflect.DeepEqual(perms[i], perms[i-1]) {
			uniquePerms = append(uniquePerms, perms[i])
		}
	}
	return uniquePerms
}
