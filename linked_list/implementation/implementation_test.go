package implementation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddFront(t *testing.T) {
	cases := []struct {
		data   []string
		output []string
		len    int
	}{
		{
			[]string{},
			[]string{},
			0,
		},
		{
			[]string{"1", "2"},
			[]string{"2", "1"},
			2,
		},
		{
			[]string{"1"},
			[]string{"1"},
			1,
		},
		{
			[]string{"1", "2", "3"},
			[]string{"3", "2", "1"},
			2,
		},
	}

	for _, c := range cases {
		l := NewLinkedList()
		for i := 0; i < len(c.data); i++ {
			l.AddFront(c.data[i])
		}

		require.Equal(t, c.output, l.GetAll())
		require.Equal(t, len(c.output), l.len)
	}
}
