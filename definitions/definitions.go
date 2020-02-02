package definitions

import "errors"

var (
	// ErrEmptyList - Error for empty list
	ErrEmptyList = errors.New("empty list")

	// ErrAtLeastTwo - Error for list needs at least two elements
	ErrAtLeastTwo = errors.New("list needs at least two elements")

	// ErrNegativeIndex - Error for list index cannot be negative
	ErrNegativeIndex = errors.New("list index cannot be negative")

	// ErrNotPositiveNumber - Error for factor cannot be 0 or negative
	ErrNotPositiveNumber = errors.New("factor cannot be 0 or negative")

	// ErrNegativeNumber - Error for factor cannot be negative
	ErrNegativeNumber = errors.New("factor cannot be negative")

	// ErrZero - Error for factor cannot be 0
	ErrZero = errors.New("factor cannot be 0")

	// ErrIndexOutOfRange - Error for index out of range
	ErrIndexOutOfRange = errors.New("index out of range")
)

// RLE is the type in
// Run Length Encoding
type RLE struct {
	Count int
	Value string
}
