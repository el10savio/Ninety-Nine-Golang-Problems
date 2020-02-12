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

func GCD(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, definitions.ErrNegativeNumber
	}

	if a == 0 {
		return b, nil
	}

	if b == 0 {
		return a, nil
	}

	gcd, err := GCD(b%a, a)
	if err != nil {
		return 0, err
	}

	return gcd, nil
}

func Coprime(a int, b int) (bool, error) {
	gcd, err := GCD(b%a, a)
	if err != nil {
		return false, err
	}

	return gcd == 1, nil
}

func Totient(m int) (int, error) {
	if m <= 0 {
		return 0, definitions.ErrNotPositiveNumber
	}

	count := 1

	for i := 2; i < m; i++ {
		comprime, err := Coprime(m, i)
		if err != nil {
			return 0, err
		}

		if comprime {
			count++
		}
	}

	return count, nil
}

func PrimeFactors(number int) ([]int, error) {
	if number <= 0 {
		return []int{}, definitions.ErrNotPositiveNumber
	}

	if number <= 3 {
		return []int{number}, nil
	}

	sieve := definitions.SieveOfEratosthenes(number)
	primeFactors := make([]int, 0)

	for index := 0; index < len(sieve); index++ {
		if number%sieve[index] == 0 {
			primeFactors = append(primeFactors, sieve[index])
			number = number / sieve[index]
			index = -1
		}
	}

	return primeFactors, nil
}
