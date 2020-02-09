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
