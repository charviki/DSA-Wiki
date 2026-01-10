package easy

import "testing"

func Test_168_convertToTitle(t *testing.T) {
	tests := []struct {
		name         string
		columnNumber int
		want         string
	}{
		{
			name:         "Example 1",
			columnNumber: 1,
			want:         "A",
		},
		{
			name:         "Example 2",
			columnNumber: 28,
			want:         "AB",
		},
		{
			name:         "Example 3",
			columnNumber: 701,
			want:         "ZY",
		},
		{
			name:         "Example 4",
			columnNumber: 2147483647,
			want:         "FXSHRXW",
		},
		{
			name:         "Smallest double digit",
			columnNumber: 27,
			want:         "AA",
		},
		{
			name:         "Last single digit",
			columnNumber: 26,
			want:         "Z",
		},
		{
			name:         "Smallest triple digit",
			columnNumber: 703,
			want:         "AAA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _168_convertToTitle(tt.columnNumber); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
