package number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPrimeNumbers(t *testing.T) {
	type result struct {
		hasError bool
		expected []bool
	}
	input := map[string]result{
		"[1,2,3,4,\"nan\"]": {
			hasError: true,
		},
		"[4,6,11,88,101]": {
			hasError: false,
			expected: []bool{false, false, true, false, true},
		},
		"[\"name\": \"Sam\"]": {
			hasError: true,
		},
		"[false, true, 0]": {
			hasError: true,
		},
		"[-1, 10, 5]": {
			hasError: true,
		},
		"[0, 11]": {
			hasError: true,
		},
	}

	for in, res := range input {
		processed, err := FindPrimeNumbers([]byte(in))
		if err != nil && !res.hasError {
			t.Errorf("wrong result for: %s actual: %+v expected: %+v", in, err, res.expected)
		}
		if err == nil && !res.hasError {
			assert.EqualValues(t, res.expected, processed, in)
		}
	}
}

func BenchmarkFindPrimeNumbers(b *testing.B) {
	for c := 0; c < b.N; c++ {
		FindPrimeNumbers([]byte("[4,6,11,88,101,14123,182839,10000220]"))
	}
}
