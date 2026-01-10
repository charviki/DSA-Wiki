package easy

import "testing"

func Test_125_isPalindrome(t *testing.T) {
    tests := []struct {
        name string
        s    string
        want bool
    }{
        {
            name: "Example 1",
            s:    "A man, a plan, a canal: Panama",
            want: true,
        },
        {
            name: "Example 2",
            s:    "race a car",
            want: false,
        },
        {
            name: "Example 3",
            s:    " ",
            want: true,
        },
        {
            name: "Empty string",
            s:    "",
            want: true,
        },
        {
            name: "Single char",
            s:    "a",
            want: true,
        },
        {
            name: "Simple palindrome",
            s:    "aba",
            want: true,
        },
        {
            name: "Numbers",
            s:    "121",
            want: true,
        },
        {
            name: "Mixed alphanumeric",
            s:    "0P",
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := _125_isPalindrome(tt.s); got != tt.want {
                t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
            }
        })
    }
}
