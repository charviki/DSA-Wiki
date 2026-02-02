package easy

import (
	"reflect"
	"testing"
)

func Test_228_summaryRanges(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []string
	}{
		{"Example 1", []int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{"Example 2", []int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		{"Empty", []int{}, nil},
		{"Single", []int{1}, []string{"1"}},
		{"Negative", []int{-1}, []string{"-1"}},
		{"Negative Range", []int{-3, -2, -1}, []string{"-3->-1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _228_summaryRanges(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_228_summaryRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
