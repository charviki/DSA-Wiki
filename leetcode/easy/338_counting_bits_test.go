package easy

import (
	"reflect"
	"testing"
)

func Test_338_countBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"Example 1", 2, []int{0, 1, 1}},
		{"Example 2", 5, []int{0, 1, 1, 2, 1, 2}},
		{"Zero", 0, []int{0}},
		{"One", 1, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _338_countBits(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_338_countBits() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := _338_countBits2(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_338_countBits() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := _338_countBits3(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_338_countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
