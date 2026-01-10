package medium

import "testing"

func Test_11_maxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Example 1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "Example 2",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "Increasing height",
			height: []int{1, 2, 3, 4, 5},
			want:   6, // 2 and 5 (index 1 and 4, dist 3, min height 2 -> 6) vs 1 and 5 (index 0 and 4, dist 4, min height 1 -> 4). 3 and 5 (idx 2,4 dist 2 min 3 -> 6). Wait, 2 and 5: idx 1,4 width 3, height 2 -> 6. 3 and 5: idx 2,4 width 2, height 3 -> 6. 4 and 5: width 1, height 4 -> 4.
			// Actually, best is index 1 (val 2) and index 4 (val 5) -> width 3 * 2 = 6. Or index 2 (val 3) and 4 (val 5) -> width 2 * 3 = 6.
		},
		{
			name:   "Decreasing height",
			height: []int{5, 4, 3, 2, 1},
			want:   6,
		},
		{
			name:   "Same height",
			height: []int{5, 5, 5, 5},
			want:   15, // index 0 and 3 -> width 3 * 5 = 15
		},
		{
			name:   "Two bars",
			height: []int{10, 20},
			want:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _11_maxArea(tt.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
			if got := _11_maxArea2(tt.height); got != tt.want {
				t.Errorf("maxArea2() = %v, want %v", got, tt.want)
			}
		})
	}
}
