package arithmetic

import (
	"log"
	"reflect"
	"testing"

	"github.com/el10savio/Ninety-Nine-Golang-Problems/definitions"
)

func TestIsPrime(t *testing.T) {
	number := 7

	expectedIsPrime := true
	var expectedErr error

	actualIsPrime, actualErr := IsPrime(number)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedIsPrime, actualIsPrime) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedIsPrime, actualIsPrime)
	}
}

func TestIsPrime_NegativeNumber(t *testing.T) {
	number := -7

	expectedIsPrime := false
	expectedErr := definitions.ErrNotPositiveNumber

	actualIsPrime, actualErr := IsPrime(number)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedIsPrime, actualIsPrime) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedIsPrime, actualIsPrime)
	}
}

func TestIsPrime_Zero(t *testing.T) {
	number := 0

	expectedIsPrime := false
	expectedErr := definitions.ErrNotPositiveNumber

	actualIsPrime, actualErr := IsPrime(number)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedIsPrime, actualIsPrime) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedIsPrime, actualIsPrime)
	}
}

func TestIsPrime_OneTwoThree(t *testing.T) {
	numbers := []int{1, 2, 3}

	expectedIsPrime := true
	var expectedErr error

	for _, number := range numbers {
		actualIsPrime, actualErr := IsPrime(number)

		if !reflect.DeepEqual(expectedErr, actualErr) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
		}

		if !reflect.DeepEqual(expectedIsPrime, actualIsPrime) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedIsPrime, actualIsPrime)
		}
	}
}

func TestIsPrime_Four(t *testing.T) {
	number := 4

	expectedIsPrime := false
	var expectedErr error

	actualIsPrime, actualErr := IsPrime(number)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedIsPrime, actualIsPrime) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedIsPrime, actualIsPrime)
	}
}

func TestGCD(t *testing.T) {
	a, b := 36, 63
	// 98, 56 => 14
	// 48, 18 => 6

	expectedGCD := 9
	var expectedErr error

	actualGCD, actualErr := GCD(a, b)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedGCD, actualGCD) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedGCD, actualGCD)
	}
}

func TestGCD_NegativeValue(t *testing.T) {
	a, b := -36, 63

	expectedGCD := 0
	expectedErr := definitions.ErrNegativeNumber

	actualGCD, actualErr := GCD(a, b)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedGCD, actualGCD) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedGCD, actualGCD)
	}
}

func TestGCD_ZeroValue(t *testing.T) {
	a, b := 0, 63

	expectedGCD := 63
	var expectedErr error

	actualGCD, actualErr := GCD(a, b)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedGCD, actualGCD) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedGCD, actualGCD)
	}
}

func TestCoprime(t *testing.T) {
	a, b := 35, 64

	expectedCoprime := true
	var expectedErr error

	actualCoprime, actualErr := Coprime(a, b)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedCoprime, actualCoprime) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedCoprime, actualCoprime)
	}
}

func TestCoprime_NegativeValue(t *testing.T) {
	a, b := -35, 64

	expectedCoprime := false
	expectedErr := definitions.ErrNegativeNumber

	actualCoprime, actualErr := Coprime(a, b)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedCoprime, actualCoprime) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedCoprime, actualCoprime)
	}
}

func TestTotient(t *testing.T) {
	Ms := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedTotients := []int{1, 1, 2, 2, 4, 2, 6, 4, 6, 4}

	var expectedErr error

	for index := 0; index < len(Ms); index++ {
		m := Ms[index]
		expectedTotient := expectedTotients[index]

		actualTotient, actualErr := Totient(m)

		if !reflect.DeepEqual(expectedErr, actualErr) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
		}

		if !reflect.DeepEqual(expectedTotient, actualTotient) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedTotient, actualTotient)
		}
	}
}

func TestTotient_NegativeM(t *testing.T) {
	m := -10

	expectedTotient := 0
	expectedErr := definitions.ErrNotPositiveNumber

	actualTotient, actualErr := Totient(m)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedTotient, actualTotient) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedTotient, actualTotient)
	}
}

func TestTotient_ZeroM(t *testing.T) {
	m := 0

	expectedTotient := 0
	expectedErr := definitions.ErrNotPositiveNumber

	actualTotient, actualErr := Totient(m)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedTotient, actualTotient) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedTotient, actualTotient)
	}
}

func TestPrimeFactors(t *testing.T) {
	Numbers := []int{2, 33, 68, 97, 177}
	expectedPrimeFactorsList := [][]int{{2}, {3, 11}, {2, 2, 17}, {97}, {3, 59}}

	var expectedErr error

	for index := 0; index < len(Numbers); index++ {
		Number := Numbers[index]
		expectedPrimeFactors := expectedPrimeFactorsList[index]

		log.Println(Number, expectedPrimeFactors)

		actualPrimeFactors, actualErr := PrimeFactors(Number)

		if !reflect.DeepEqual(expectedErr, actualErr) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
		}

		if !reflect.DeepEqual(expectedPrimeFactors, actualPrimeFactors) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedPrimeFactors, actualPrimeFactors)
		}
	}
}
