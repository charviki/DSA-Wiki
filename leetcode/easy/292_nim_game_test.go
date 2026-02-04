package easy

import "testing"

func Test_292_canWinNim(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"Example 1", 4, false},
		{"Example 2", 1, true},
		{"Example 3", 2, true},
		{"Can Win 3", 3, true},
		{"Can Win 5", 5, true}, // You take 1, left 4, opponent loses
		{"Can Win 6", 6, true}, // You take 2, left 4, opponent loses
		{"Can Win 7", 7, true}, // You take 3, left 4, opponent loses
		{"Lose 8", 8, false},   // Whatever you take (1,2,3), left (7,6,5), opponent can win
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _292_canWinNim(tt.n); got != tt.want {
				t.Errorf("_292_canWinNim() = %v, want %v", got, tt.want)
			}
		})
	}
}
