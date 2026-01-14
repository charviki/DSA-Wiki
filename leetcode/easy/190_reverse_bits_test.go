package easy

import "testing"

func Test_190_reverseBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1",
			n:    43261596,
			want: 964176192,
		},
		{
			name: "Example 2",
			n:    2147483644,
			want: 1073741822,
		},
		{
			name: "Example 3",
			n:    2,
			want: 1073741824,
		},
		{
			name: "Example 4",
			n:    2147483646,
			want: 2147483646,
		},
		{
			name: "Example 5",
			n:    0,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _190_reverseBits(tt.n); got != tt.want {
				t.Errorf("_190_reverseBits() = %v, want %v", got, tt.want)
			}
			if got := _190_reverseBits2(tt.n); got != tt.want {
				t.Errorf("_190_reverseBits2() = %v, want %v", got, tt.want)
			}
		})
	}
}
