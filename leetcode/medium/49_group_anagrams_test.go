package medium

import (
	"reflect"
	"sort"
	"testing"
)

func _49_sortGroups(groups [][]string) {
	for _, g := range groups {
		sort.Strings(g)
	}
	sort.Slice(groups, func(i, j int) bool {
		n := len(groups[i])
		if n != len(groups[j]) {
			return n < len(groups[j])
		}
		for k := 0; k < n; k++ {
			if groups[i][k] != groups[j][k] {
				return groups[i][k] < groups[j][k]
			}
		}
		return false
	})
}

func Test_49_groupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "Example 1",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			name: "Example 2",
			strs: []string{""},
			want: [][]string{
				{""},
			},
		},
		{
			name: "Example 3",
			strs: []string{"a"},
			want: [][]string{
				{"a"},
			},
		},
		{
			name: "No anagrams",
			strs: []string{"abc", "def", "ghi"},
			want: [][]string{
				{"abc"}, {"def"}, {"ghi"},
			},
		},
		{
			name: "Multiple empty strings",
			strs: []string{"", "", ""},
			want: [][]string{
				{"", "", ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := _49_groupAnagrams(tt.strs)
			_49_sortGroups(got)
			_49_sortGroups(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_49_groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
