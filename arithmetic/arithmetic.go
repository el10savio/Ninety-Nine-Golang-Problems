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

	sieve := definitions.SieveOfEratosthenes(value)

	for _, prime := range sieve {
		if value == prime {
			return true, nil
		}
	}

	return false, nil
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

func PrimeFactorsMult(number int) ([][]int, error) {
	if number <= 0 {
		return [][]int{}, definitions.ErrNotPositiveNumber
	}

	primeFactors, err := PrimeFactors(number)
	if err != nil {
		return [][]int{}, err
	}

	primeFactorsMult := make([][]int, 0)

	previousPrimeFactor := primeFactors[0]
	mult := 1

	if len(primeFactors) == 1 {
		primeFactorsMult = append(primeFactorsMult, []int{previousPrimeFactor, mult})
		return primeFactorsMult, nil
	}

	for index := 1; index < len(primeFactors); index++ {
		if previousPrimeFactor != primeFactors[index] {
			primeFactorsMult = append(primeFactorsMult, []int{previousPrimeFactor, mult})
			previousPrimeFactor = primeFactors[index]
			mult = 1
		} else {
			mult++
		}
	}
	primeFactorsMult = append(primeFactorsMult, []int{previousPrimeFactor, mult})

	return primeFactorsMult, nil
}
