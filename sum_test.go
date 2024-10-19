package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "expected return 3",
			a:        1,
			b:        2,
			expected: 3,
		},
		{
			name:     "expected return 1",
			a:        -1,
			b:        2,
			expected: 1,
		},
		{
			name:     "expected return 0",
			a:        0,
			b:        0,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			result := Sum(test.a, test.b)
			assert.Equal(t, test.expected, result)
		})
	}
}
