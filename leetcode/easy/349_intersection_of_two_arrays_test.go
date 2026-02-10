package easy

import (
	"reflect"
	"sort"
	"testing"
)

func Test_349_intersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}},
			want: []int{2},
		},
		{
			name: "Example 2",
			args: args{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}},
			want: []int{9, 4},
		},
		{
			name: "No Intersection",
			args: args{nums1: []int{1, 2, 3}, nums2: []int{4, 5, 6}},
			want: []int{},
		},
		{
			name: "Identical Arrays",
			args: args{nums1: []int{1, 2, 3}, nums2: []int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
		{
			name: "Empty Array 1",
			args: args{nums1: []int{}, nums2: []int{1, 2, 3}},
			want: []int{},
		},
		{
			name: "Empty Array 2",
			args: args{nums1: []int{1, 2, 3}, nums2: []int{}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := _349_intersection(tt.args.nums1, tt.args.nums2)
			// Sort both slices to ignore order
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_349_intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
