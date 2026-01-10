package easy

import (
    "reflect"
    "testing"
)

func Test_118_generate(t *testing.T) {
    tests := []struct {
        name    string
        numRows int
        want    [][]int
    }{
        {
            name:    "Example 1",
            numRows: 5,
            want: [][]int{
                {1},
                {1, 1},
                {1, 2, 1},
                {1, 3, 3, 1},
                {1, 4, 6, 4, 1},
            },
        },
        {
            name:    "Example 2",
            numRows: 1,
            want: [][]int{
                {1},
            },
        },
        {
            name:    "Two rows",
            numRows: 2,
            want: [][]int{
                {1},
                {1, 1},
            },
        },
        {
            name:    "Three rows",
            numRows: 3,
            want: [][]int{
                {1},
                {1, 1},
                {1, 2, 1},
            },
        },
        {
            name:    "Four rows",
            numRows: 4,
            want: [][]int{
                {1},
                {1, 1},
                {1, 2, 1},
                {1, 3, 3, 1},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := _118_generate(tt.numRows); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("generate() = %v, want %v", got, tt.want)
            }
        })
    }
}
