package arithmetic

import (
	"math"
	"math/rand"

	"github.com/el10savio/Ninety-Nine-Golang-Problems/definitions"
)

func millerTest(d int, value int) bool {
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

func IsPrime(value int) (bool, error) {
	if value <= 0 {
		return false, definitions.ErrNotPositiveNumber
	}

	if value <= 3 {
		return true, nil
	}

	if value == 4 {
		return false, nil
	}

	// n = 2^d * r + 1 for some r >= 1
	d := value - 1
	for d%2 == 0 {
		d = d / 2
	}

	retries := 4
	for i := 0; i < retries; i++ {
		if millerTest(d, value) == false {
			return false, nil
		}
	}

	return true, nil
}
