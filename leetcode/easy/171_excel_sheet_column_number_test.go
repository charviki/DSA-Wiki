package easy

import "testing"

func Test_171_titleToNumber(t *testing.T) {
	tests := []struct {
		name        string
		columnTitle string
		want        int
	}{
		{
			name:        "Example 1",
			columnTitle: "A",
			want:        1,
		},
		{
			name:        "Example 2",
			columnTitle: "AB",
			want:        28,
		},
		{
			name:        "Example 3",
			columnTitle: "ZY",
			want:        701,
		},
		{
			name:        "Example 4",
			columnTitle: "FXSHRXW",
			want:        2147483647,
		},
		{
			name:        "Smallest double digit",
			columnTitle: "AA",
			want:        27,
		},
		{
			name:        "Last single digit",
			columnTitle: "Z",
			want:        26,
		},
		{
			name:        "Smallest triple digit",
			columnTitle: "AAA",
			want:        703,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _171_titleToNumber(tt.columnTitle); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
			if got := _171_titleToNumber2(tt.columnTitle); got != tt.want {
				t.Errorf("titleToNumber2() = %v, want %v", got, tt.want)
			}
		})
	}
}
