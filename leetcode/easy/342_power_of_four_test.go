package easy

import "testing"

func Test_342_isPowerOfFour(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{n: 16},
			want: true,
		},
		{
			name: "Example 2",
			args: args{n: 5},
			want: false,
		},
		{
			name: "Example 3",
			args: args{n: 1},
			want: true,
		},
		{
			name: "Zero",
			args: args{n: 0},
			want: false,
		},
		{
			name: "Power of 4 (4)",
			args: args{n: 4},
			want: true,
		},
		{
			name: "Power of 4 (64)",
			args: args{n: 64},
			want: true,
		},
		{
			name: "Negative number",
			args: args{n: -16},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _342_isPowerOfFour(tt.args.n); got != tt.want {
				t.Errorf("_342_isPowerOfFour() = %v, want %v", got, tt.want)
			}
			if got := _342_isPowerOfFour2(tt.args.n); got != tt.want {
				t.Errorf("_342_isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
