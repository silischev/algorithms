package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxValueLinearComplexitySimple(t *testing.T) {
	cases := []struct {
		data   []int
		maxVal int
	}{
		{[]int{1, 2, 1}, 5},
		{[]int{1, 3, 1, 2}, 10},
		{[]int{1, 3, 1, 2, 1}, 11},
		{[]int{1, 3, 1, 2, 1, 1, 2}, 16},
		{[]int{1, 4, 2, 1}, 11},
		{[]int{1, 3, 5, 2}, 17},
		{[]int{1, 3, 2, 5}, 20},
		{[]int{1, 3, 4, 5, 2}, 22},
		{[]int{1, 3, 4, 5, 2, 1}, 23},
		{[]int{1, 3, 3}, 9},
		{[]int{1, 1, 3}, 9},
		{[]int{1, 1, 1}, 3},
	}

	for _, c := range cases {
		act := maxValueLinearComplexitySimple(c.data)
		assert.Equal(t, c.maxVal, act, c.data)
	}
}

func TestMaxValueLinearComplexity(t *testing.T) {
	cases := []struct {
		data   []int
		maxVal int
	}{
		{[]int{1, 2, 1}, 5},
		{[]int{1, 3, 1, 2}, 10},
		{[]int{1, 3, 1, 2, 1}, 11},
		{[]int{1, 3, 1, 2, 1, 1, 2}, 16},
		{[]int{1, 4, 2, 1}, 11},
		{[]int{1, 3, 5, 2}, 17},
		{[]int{1, 3, 2, 5}, 20},
		{[]int{1, 3, 4, 5, 2}, 22},
		{[]int{1, 3, 4, 5, 2, 1}, 23},
		{[]int{1, 3, 3}, 9},
		{[]int{1, 1, 3}, 9},
		{[]int{1, 1, 1}, 3},
	}

	for _, c := range cases {
		act := maxValueLinearComplexity(c.data)
		assert.Equal(t, c.maxVal, act, c.data)
	}
}

func TestMaxValueQuadraticComplexity(t *testing.T) {
	cases := []struct {
		data   []int
		maxVal int
	}{
		{[]int{1, 2, 1}, 5},
		{[]int{1, 3, 1, 2}, 10},
		{[]int{1, 3, 1, 2, 1}, 11},
		{[]int{1, 3, 1, 2, 1, 1, 2}, 16},
		{[]int{1, 4, 2, 1}, 11},
		{[]int{1, 3, 5, 2}, 17},
		{[]int{1, 3, 2, 5}, 20},
		{[]int{1, 3, 4, 5, 2}, 22},
		{[]int{1, 3, 4, 5, 2, 1}, 23},
		{[]int{1, 3, 3}, 9},
		{[]int{1, 1, 3}, 9},
		{[]int{1, 1, 1}, 3},
	}

	for _, c := range cases {
		act := maxValueQuadraticComplexity(c.data)
		assert.Equal(t, c.maxVal, act, c.data)
	}
}
