package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesToSize(t *testing.T) {
	tests := []struct {
		bytes    int64
		expected string
	}{
		{0, "0B"},
		{1023, "1023B"},
		{1024, "1KB"},
		{1048576, "1MB"},
		{1073741824, "1GB"},
		{1099511627776, "1TB"},
		{1125899906842624, "1024TB"},
	}

	for _, test := range tests {
		result := BytesToSize(test.bytes)
		assert.Equal(t, test.expected, result)
	}
}
