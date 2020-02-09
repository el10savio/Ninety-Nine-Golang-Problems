package arithmetic

import (
	"github.com/el10savio/Ninety-Nine-Golang-Problems/definitions"
)

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
		if definitions.MillerTest(d, value) == false {
			return false, nil
		}
	}

	return true, nil
}
