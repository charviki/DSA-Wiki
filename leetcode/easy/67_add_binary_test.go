package easy

import "testing"

func Test_67_addBinary(t *testing.T) {
    tests := []struct {
        name string
        a    string
        b    string
        want string
    }{
        {
            name: "Example 1",
            a:    "11",
            b:    "1",
            want: "100",
        },
        {
            name: "Example 2",
            a:    "1010",
            b:    "1011",
            want: "10101",
        },
        {
            name: "Zero",
            a:    "0",
            b:    "0",
            want: "0",
        },
        {
            name: "Different length",
            a:    "1",
            b:    "111",
            want: "1000",
        },
        {
            name: "Carry over",
            a:    "111",
            b:    "1",
            want: "1000",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := _67_addBinary(tt.a, tt.b); got != tt.want {
                t.Errorf("addBinary() = %v, want %v", got, tt.want)
            }
        })
    }
}
