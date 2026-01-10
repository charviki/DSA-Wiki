package easy

import "testing"

func Test_26_removeDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		want     int
		wantNums []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 2},
			want:     2,
			wantNums: []int{1, 2},
		},
		{
			name:     "Example 2",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want:     5,
			wantNums: []int{0, 1, 2, 3, 4},
		},
		{
			name:     "No duplicates",
			nums:     []int{1, 2, 3},
			want:     3,
			wantNums: []int{1, 2, 3},
		},
		{
			name:     "All duplicates",
			nums:     []int{1, 1, 1},
			want:     1,
			wantNums: []int{1},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			want:     0,
			wantNums: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := _26_removeDuplicates(tt.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			for i := 0; i < got; i++ {
				if tt.nums[i] != tt.wantNums[i] {
					t.Errorf("nums[%d] = %v, want %v", i, tt.nums[i], tt.wantNums[i])
				}
			}
		})
	}
}
