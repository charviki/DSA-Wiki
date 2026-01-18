package easy

import "testing"

func Test_202_isHappy(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"Example 1", 19, true},
		{"Example 2", 2, false},
		{"Case 7", 7, true},
		{"Case 1", 1, true},
		{"Case 4", 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _202_isHappy(tt.n); got != tt.want {
				t.Errorf("_202_isHappy() = %v, want %v", got, tt.want)
			}
			if got := _202_isHappy2(tt.n); got != tt.want {
				t.Errorf("_202_isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
