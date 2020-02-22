package arithmetic

import (
	"math"

	"github.com/el10savio/Ninety-Nine-Golang-Problems/definitions"
)

// IsPrime Problem 31
// Determine whether a given integer number is prime.
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

// GCD Problem 32
// Determine the greatest common divisor of two
// positive integer numbers.
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

// Coprime Problem 33
// Determine whether two positive integer
// numbers are coprime.
func Coprime(a int, b int) (bool, error) {
	gcd, err := GCD(b%a, a)
	if err != nil {
		return false, err
	}

	return gcd == 1, nil
}

// Totient Problem 34
// Calculate Euler's totient function phi(m).
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

// PrimeFactors Problem 35
// Determine the prime factors of a
// given positive integer.
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

// PrimeFactorsMult Problem 36
// Determine the prime factors of a
// given positive integer.
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

// TotientImproved Problem 37
// Calculate Euler's totient
// function phi(m) (improved).
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

// Problem 38
// Compare the two methods of calculating
// Euler's totient function.
// BenchmarkTotient10090
// BenchmarkTotientImproved10090
// goos: linux
// goarch: amd64
// BenchmarkTotient10090-4           	     650	   1832848 ns/op
// BenchmarkTotientImproved10090-4   	   18412	     68079 ns/op

// PrimesR Problem 39
// Given a range of integers by its lower and
// upper limit, construct a list of all
// prime numbers in that range.
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

// Goldbach Problem 40
// Goldbach's conjecture.
func Goldbach(number int) ([]int, error) {
	if number <= 2 {
		return []int{}, definitions.ErrAtLeastTwo
	}

	if number%2 != 0 {
		return []int{}, definitions.ErrOddNumber
	}

	primesToNumber := definitions.SieveOfEratosthenes(number)

	for index := len(primesToNumber) - 1; index >= 0; index-- {
		prime1 := primesToNumber[index]
		prime2 := number - prime1

		if prime1 <= 1 || prime2 <= 1 {
			continue
		}

		isPrime, err := IsPrime(prime2)
		if err != nil {
			return []int{}, err
		}

		if isPrime == false {
			continue
		}

		if prime1+prime2 == number {
			return []int{prime2, prime1}, nil
		}
	}

	return []int{}, nil
}

// GoldbachList Problem 41
// Given a range of integers by its lower
// and upper limit, print a list of all
// even numbers and their
// Goldbach composition.
func GoldbachList(lower int, upper int) ([][]int, error) {
	if lower <= 2 || upper <= 2 {
		return [][]int{}, definitions.ErrAtLeastTwo
	}

	if upper < lower {
		upper, lower = lower, upper
	}

	if lower%2 != 0 {
		lower++
	}

	goldbachListLarge := make([][]int, 0)

	for index := lower; index <= upper; index += 2 {
		goldbach, err := Goldbach(index)
		if err != nil {
			return [][]int{}, err
		}

		goldbachListLarge = append(goldbachListLarge, goldbach)
	}

	return goldbachListLarge, nil
}

// GoldbachListLarge Problem 41 II
// Find out how many such cases there are in a
// range that are both greater than a value.
func GoldbachListLarge(lower int, upper int, value int) ([][]int, error) {
	if value <= 0 {
		return [][]int{}, definitions.ErrNotPositiveNumber
	}

	if lower <= 2 || upper <= 2 {
		return [][]int{}, definitions.ErrAtLeastTwo
	}

	if upper < lower {
		upper, lower = lower, upper
	}

	if lower%2 != 0 {
		lower++
	}

	goldbachListLarge := make([][]int, 0)

	for index := lower; index <= upper; index += 2 {
		goldbach, err := Goldbach(index)
		if err != nil {
			return [][]int{}, err
		}

		if goldbach[0] >= value && goldbach[1] >= value {
			goldbachListLarge = append(goldbachListLarge, goldbach)
		}
	}

	return goldbachListLarge, nil
}
