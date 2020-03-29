package logicandcodes

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/el10savio/Ninety-Nine-Golang-Problems/definitions"
)

type testCase struct {
	Number   int
	GrayCode []string
	Err      error
}

var testCases = []testCase{
	testCase{-1, []string{}, definitions.ErrNotPositiveNumber},
	testCase{0, []string{}, definitions.ErrNotPositiveNumber},
	testCase{1, []string{"0", "1"}, nil},
	testCase{2, []string{"00", "01", "11", "10"}, nil},
	testCase{3, []string{"000", "001", "011", "010", "110", "111", "101", "100"}, nil},
}

func TestParser(t *testing.T) {
	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			actualGrayCode, actualErr := GrayCodes(testCase.Number)
			assert.Equal(t, testCase.Err, actualErr)
			assert.Equal(t, testCase.GrayCode, actualGrayCode)
		})
	}
}
