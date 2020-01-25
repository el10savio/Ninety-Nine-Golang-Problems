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
