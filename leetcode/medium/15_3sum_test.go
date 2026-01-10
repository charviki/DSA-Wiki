package medium

import (
	"reflect"
	"sort"
	"testing"
)

func Test_15_threeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "Example 2",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "Example 3",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "No solution",
			nums: []int{1, 2, 3},
			want: [][]int{},
		},
		{
			name: "Multiple zeros",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := _15_threeSum(tt.nums)
			if !checkThreeSumResult(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// checkThreeSumResult compares two lists of triplets ignoring order
func checkThreeSumResult(got, want [][]int) bool {
	if len(got) != len(want) {
		return false
	}

	normalize := func(arr [][]int) [][]int {
		// Sort each triplet
		for i := range arr {
			sort.Ints(arr[i])
		}
		// Sort the list of triplets
		sort.Slice(arr, func(i, j int) bool {
			for k := 0; k < 3; k++ {
				if arr[i][k] != arr[j][k] {
					return arr[i][k] < arr[j][k]
				}
			}
			return false
		})
		return arr
	}

	// Deep copy to avoid modifying original tests if we were reusing them (though we aren't really)
	// Actually, normalize modifies in place. It's safer to copy if 'want' is shared, but here it's fine or we can copy.
	// Let's copy 'got' and 'want' just to be safe.
	gotCopy := make([][]int, len(got))
	for i, v := range got {
		gotCopy[i] = make([]int, len(v))
		copy(gotCopy[i], v)
	}
	wantCopy := make([][]int, len(want))
	for i, v := range want {
		wantCopy[i] = make([]int, len(v))
		copy(wantCopy[i], v)
	}

	normGot := normalize(gotCopy)
	normWant := normalize(wantCopy)

	return reflect.DeepEqual(normGot, normWant)
}
