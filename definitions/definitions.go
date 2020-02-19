package definitions

import (
	"errors"
	"math"
	"math/rand"
)

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

	// ErrLessThanTwo - Error for factor cannot be less than 3
	ErrLessThanTwo = errors.New("factor cannot be less than 3")

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

func MillerTest(d int, value int) bool {
	// Random number between [2, n-2]
	a := 2 + rand.Intn(value-4) + 1

	// a^d % n
	x := int(math.Pow(float64(a), float64(d))) % value

	if x == 1 || x == value-1 {
		return true
	}

	for d != value-1 {
		x := (x * x) % value
		d *= 2

		if x == 1 {
			return false
		}

		if x == value-1 {
			return true
		}
	}

	return false
}

func SieveOfEratosthenes(maxN int) []int {
	if maxN <= 1 {
		return []int{}
	}

	maxN++
	sieve := make([]bool, maxN)

	for i := 1; i < maxN; i++ {
		sieve[i] = true
	}

	for i := 2; i <= int(math.Sqrt(float64(maxN))); i++ {
		if sieve[i] {
			for j := i * 2; j < maxN; j += i {
				sieve[j] = false
			}
		}
	}

	primeList := make([]int, 0)
	for i := 2; i < maxN; i++ {
		if sieve[i] {
			primeList = append(primeList, i)
		}
	}

	return primeList
}
