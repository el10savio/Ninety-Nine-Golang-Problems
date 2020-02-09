package arithmetic

import (
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
