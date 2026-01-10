package easy

import "testing"

func Test_70_climbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1",
			n:    2,
			want: 2,
		},
		{
			name: "Example 2",
			n:    3,
			want: 3,
		},
		{
			name: "1 step",
			n:    1,
			want: 1,
		},
		{
			name: "4 steps",
			n:    4,
			want: 5,
		},
		{
			name: "5 steps",
			n:    5,
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _70_climbStairs(tt.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
