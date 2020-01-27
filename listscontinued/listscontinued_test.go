package listscontinued

import (
	"reflect"
	"testing"

	"github.com/el10savio/Ninety-Nine-Golang-Problems/lists"
)

func TestEncodeModified(t *testing.T) {
	list := []string{"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "e", "e", "e", "e"}

	expectedEncodeModified := []lists.RLE{{4, "a"}, {0, "b"}, {2, "c"}, {2, "a"}, {0, "d"}, {4, "e"}}
	var expectedErr error

	actualEncodeModified, actualErr := EncodeModified(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedEncodeModified, actualEncodeModified) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedEncodeModified, actualEncodeModified)
	}
}

func TestEncodeModified_EmptyList(t *testing.T) {
	list := []string{}

	expectedEncodeModified := []lists.RLE{}
	expectedErr := lists.ErrEmptyList

	actualEncodeModified, actualErr := EncodeModified(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedEncodeModified, actualEncodeModified) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedEncodeModified, actualEncodeModified)
	}
}

func TestDecodeModified(t *testing.T) {
	list := []lists.RLE{{4, "a"}, {0, "b"}, {2, "c"}, {2, "a"}, {0, "d"}, {4, "e"}}

	expectedDecodeModified := []string{"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "e", "e", "e", "e"}
	var expectedErr error

	actualDecodeModified, actualErr := DecodeModified(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedDecodeModified, actualDecodeModified) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedDecodeModified, actualDecodeModified)
	}
}

func TestDecodeModified_EmptyList(t *testing.T) {
	list := []lists.RLE{}

	expectedDecodeModified := []string{}
	expectedErr := lists.ErrEmptyList

	actualDecodeModified, actualErr := DecodeModified(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedDecodeModified, actualDecodeModified) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedDecodeModified, actualDecodeModified)
	}
}

func TestDupli(t *testing.T) {
	list := []int{1, 2, 3}

	expectedDupli := []int{1, 1, 2, 2, 3, 3}
	var expectedErr error

	actualDupli, actualErr := Dupli(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedDupli, actualDupli) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedDupli, actualDupli)
	}
}

func TestDupli_EmptyList(t *testing.T) {
	list := []int{}

	expectedDupli := []int{}
	expectedErr := lists.ErrEmptyList

	actualDupli, actualErr := Dupli(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedDupli, actualDupli) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedDupli, actualDupli)
	}
}

func TestRepli(t *testing.T) {
	list := []int{1, 2, 3}
	factor := 3

	expectedRepli := []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	var expectedErr error

	actualRepli, actualErr := Repli(list, factor)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRepli, actualRepli) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRepli, actualRepli)
	}
}

func TestRepli_EmptyList(t *testing.T) {
	list := []int{}
	factor := 3

	expectedRepli := []int{}
	expectedErr := lists.ErrEmptyList

	actualRepli, actualErr := Repli(list, factor)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRepli, actualRepli) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRepli, actualRepli)
	}
}

func TestRepli_ZeroFactor(t *testing.T) {
	list := []int{1, 2, 3}
	factor := 0

	expectedRepli := []int{}
	expectedErr := ErrNotPositiveNumber

	actualRepli, actualErr := Repli(list, factor)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRepli, actualRepli) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRepli, actualRepli)
	}
}

func TestRepli_NegativeFactor(t *testing.T) {
	list := []int{1, 2, 3}
	factor := -3

	expectedRepli := []int{}
	expectedErr := ErrNotPositiveNumber

	actualRepli, actualErr := Repli(list, factor)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRepli, actualRepli) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRepli, actualRepli)
	}
}

func TestDropEvery(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	factor := 3

	expectedDropEvery := []int{1, 2, 4, 5, 7, 8, 10}
	var expectedErr error

	actualDropEvery, actualErr := DropEvery(list, factor)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedDropEvery, actualDropEvery) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedDropEvery, actualDropEvery)
	}
}

func TestDropEvery_EmptyList(t *testing.T) {
	list := []int{}
	factor := 3

	expectedDropEvery := []int{}
	expectedErr := lists.ErrEmptyList

	actualDropEvery, actualErr := DropEvery(list, factor)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedDropEvery, actualDropEvery) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedDropEvery, actualDropEvery)
	}
}

func TestDropEvery_ZeroFactor(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	factor := 0

	expectedDropEvery := []int{}
	expectedErr := ErrNotPositiveNumber

	actualDropEvery, actualErr := DropEvery(list, factor)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedDropEvery, actualDropEvery) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedDropEvery, actualDropEvery)
	}
}

func TestDropEvery_NegativeFactor(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	factor := -3

	expectedDropEvery := []int{}
	expectedErr := ErrNotPositiveNumber

	actualDropEvery, actualErr := DropEvery(list, factor)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedDropEvery, actualDropEvery) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedDropEvery, actualDropEvery)
	}
}
