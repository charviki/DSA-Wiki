package easy

import (
	"reflect"
	"testing"
)

func Test_119_getRow(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		want     []int
	}{
		{
			name:     "Example 1",
			rowIndex: 3,
			want:     []int{1, 3, 3, 1},
		},
		{
			name:     "Example 2",
			rowIndex: 0,
			want:     []int{1},
		},
		{
			name:     "Example 3",
			rowIndex: 1,
			want:     []int{1, 1},
		},
		{
			name:     "Row 2",
			rowIndex: 2,
			want:     []int{1, 2, 1},
		},
		{
			name:     "Row 4",
			rowIndex: 4,
			want:     []int{1, 4, 6, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _119_getRow(tt.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
			if got := _119_getRow2(tt.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
