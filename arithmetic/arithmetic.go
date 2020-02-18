package arithmetic

import (
	"math"

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

	lenPrimeFactors := len(primeFactors)
	if lenPrimeFactors == 1 {
		primeFactorsMult = append(primeFactorsMult, []int{previousPrimeFactor, mult})
		return primeFactorsMult, nil
	}

	for index := 1; index < lenPrimeFactors; index++ {
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

func TotientImproved(m int) (int, error) {
	if m <= 0 {
		return 0, definitions.ErrNotPositiveNumber
	}

	primeFactorsMult, err := PrimeFactorsMult(m)
	if err != nil {
		return 0, err
	}

	totientImproved := 1

	for index := 0; index < len(primeFactorsMult); index++ {
		primeFactor := float64(primeFactorsMult[index][0])
		mult := float64(primeFactorsMult[index][1])
		totientImproved *= int(math.Pow(primeFactor*(primeFactor-1), mult-1))
	}

	return totientImproved, nil
}

// goos: linux
// goarch: amd64
// BenchmarkTotient10090-4           	     650	   1832848 ns/op
// BenchmarkTotientImproved10090-4   	   18412	     68079 ns/op

func PrimesR(m int, n int) ([]int, error) {
	if m <= 0 || n <= 0 {
		return []int{}, definitions.ErrNotPositiveNumber
	}

	pivot := m
	end := n

	if m > n {
		pivot = n
		end = m
	}

	primesList := definitions.SieveOfEratosthenes(end)
	primesListCut := make([]int, 0)

	for _, prime := range primesList {
		if prime >= pivot {
			primesListCut = append(primesListCut, prime)
		}
	}

	return primesListCut, nil
}
