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
