package lists

import (
	"reflect"
	"testing"
)

func TestMyLast(t *testing.T) {
	list := []int{1, 2, 3, 4}

	expectedMyLast := 4
	var expectedErr error

	actualMyLast, actualErr := MyLast(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedMyLast, actualMyLast) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedMyLast, actualMyLast)
	}
}

func TestMyLast_EmptyList(t *testing.T) {
	list := []int{}

	expectedMyLast := 0
	expectedErr := ErrEmptyList

	actualMyLast, actualErr := MyLast(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedMyLast, actualMyLast) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedMyLast, actualMyLast)
	}
}

func TestMyButLast(t *testing.T) {
	list := []int{1, 2, 3, 4}

	expectedMyButLast := 3
	var expectedErr error

	actualMyButLast, actualErr := MyButLast(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedMyButLast, actualMyButLast) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedMyButLast, actualMyButLast)
	}
}

func TestMyButLast_EmptyList(t *testing.T) {
	list := []int{}

	expectedMyButLast := 0
	expectedErr := ErrAtLeastTwo

	actualMyButLast, actualErr := MyButLast(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedMyButLast, actualMyButLast) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedMyButLast, actualMyButLast)
	}
}

func TestElementAt(t *testing.T) {
	list := []int{1, 2, 3}
	index := 2

	expectedElementAt := 3
	var expectedErr error

	actualElementAt, actualErr := ElementAt(list, index)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedElementAt, actualElementAt) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedElementAt, actualElementAt)
	}
}

func TestElementAt_EmptyList(t *testing.T) {
	list := []int{}
	index := 2

	expectedElementAt := 0
	expectedErr := ErrEmptyList

	actualElementAt, actualErr := ElementAt(list, index)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedElementAt, actualElementAt) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedElementAt, actualElementAt)
	}
}

func TestElementAt_NegativeIndex(t *testing.T) {
	list := []int{1, 2, 3}
	index := -2

	expectedElementAt := 0
	expectedErr := ErrNegativeIndex

	actualElementAt, actualErr := ElementAt(list, index)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedElementAt, actualElementAt) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedElementAt, actualElementAt)
	}
}

func TestMyLength(t *testing.T) {
	list := []int{1, 2, 3}

	expectedMyLength := 3

	actualMyLength := MyLength(list)

	if !reflect.DeepEqual(expectedMyLength, actualMyLength) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedMyLength, actualMyLength)
	}
}

func TestMyLength_EmptyList(t *testing.T) {
	list := []int{}

	expectedMyLength := 0

	actualMyLength := MyLength(list)

	if !reflect.DeepEqual(expectedMyLength, actualMyLength) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedMyLength, actualMyLength)
	}
}

func TestMyReverse(t *testing.T) {
	list := []int{1, 2, 3}

	expectedMyReverse := []int{3, 2, 1}
	var expectedErr error

	actualMyReverse, actualErr := MyReverse(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedMyReverse, actualMyReverse) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedMyReverse, actualMyReverse)
	}
}

func TestMyReverse_EmptyList(t *testing.T) {
	list := []int{}

	expectedMyReverse := []int{}
	expectedErr := ErrEmptyList

	actualMyReverse, actualErr := MyReverse(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedMyReverse, actualMyReverse) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedMyReverse, actualMyReverse)
	}
}

func TestIsPalindrome(t *testing.T) {
	list := []int{1, 2, 4, 8, 16, 8, 4, 2, 1}

	expectedIsPalindrome := true
	var expectedErr error

	actualIsPalindrome, actualErr := IsPalindrome(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedIsPalindrome, actualIsPalindrome) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedIsPalindrome, actualIsPalindrome)
	}
}

func TestIsPalindrome_NotPalindrome(t *testing.T) {
	list := []int{1, 2, 3}

	expectedIsPalindrome := false
	var expectedErr error

	actualIsPalindrome, actualErr := IsPalindrome(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedIsPalindrome, actualIsPalindrome) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedIsPalindrome, actualIsPalindrome)
	}
}

func TestIsPalindrome_EmptyList(t *testing.T) {
	list := []int{}

	expectedIsPalindrome := false
	expectedErr := ErrEmptyList

	actualIsPalindrome, actualErr := IsPalindrome(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedIsPalindrome, actualIsPalindrome) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedIsPalindrome, actualIsPalindrome)
	}
}
