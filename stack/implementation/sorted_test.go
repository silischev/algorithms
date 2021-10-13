package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushSorted(t *testing.T) {
	cases := []struct {
		in  []string
		out []string
	}{
		{[]string{"33", "3"}, []string{"3", "33"}},
		{[]string{"33"}, []string{"33"}},
		{[]string{"3", "33", "8"}, []string{"3", "8", "33"}},
		{[]string{"3", "8", "33"}, []string{"3", "8", "33"}},
		{[]string{"3", "8", "33", "88"}, []string{"3", "8", "33", "88"}},
		{[]string{"3", "8", "88", "33"}, []string{"3", "8", "33", "88"}},
		{[]string{"3", "88", "8", "33"}, []string{"3", "8", "33", "88"}},
		{[]string{"3", "88", "8", "33", "1"}, []string{"1", "3", "8", "33", "88"}},
		{[]string{"1", "88", "8", "33", "3"}, []string{"1", "3", "8", "33", "88"}},
		{[]string{"1", "8", "33", "3"}, []string{"1", "3", "8", "33"}},
		{[]string{"1", "88", "33", "3"}, []string{"1", "3", "33", "88"}},
		{[]string{"1", "33", "88", "3"}, []string{"1", "3", "33", "88"}},
		{[]string{"1", "2"}, []string{"1", "2"}},
		{[]string{"2", "1"}, []string{"1", "2"}},
		{[]string{"1", "2", "3"}, []string{"1", "2", "3"}},
		{[]string{"3", "2", "1"}, []string{"1", "2", "3"}},
		{[]string{"2", "3", "1"}, []string{"1", "2", "3"}},
		{[]string{"1", "3", "2"}, []string{"1", "2", "3"}},
		{[]string{"3", "1", "2"}, []string{"1", "2", "3"}},
		{[]string{"1", "1"}, []string{"1", "1"}},
		{[]string{"1", "1", "2"}, []string{"1", "1", "2"}},
		{[]string{"1", "2", "1"}, []string{"1", "1", "2"}},
		{[]string{"2", "1", "1"}, []string{"1", "1", "2"}},
	}

	for _, c := range cases {
		s := NewSortedStack()
		for _, in := range c.in {
			s.Push(in)
		}

		assert.Equal(t, len(c.out), s.Len)
		assert.Equal(t, c.out, s.AsSlice(), s.AsSlice())
	}
}
