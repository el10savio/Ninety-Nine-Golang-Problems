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

func TestSplit(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	splitPoint := 3

	expectedSplit := [][]int{{1, 2, 3}, {4, 5, 6, 7, 8, 9, 10}}
	var expectedErr error

	actualSplit, actualErr := Split(list, splitPoint)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedSplit, actualSplit) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedSplit, actualSplit)
	}
}

func TestSplit_EmptyList(t *testing.T) {
	list := []int{}
	splitPoint := 3

	expectedSplit := [][]int{}
	expectedErr := lists.ErrEmptyList

	actualSplit, actualErr := Split(list, splitPoint)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedSplit, actualSplit) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedSplit, actualSplit)
	}
}

func TestSplit_NegativeSplitPoint(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	splitPoint := -3

	expectedSplit := [][]int{}
	expectedErr := ErrIndexOutOfRange

	actualSplit, actualErr := Split(list, splitPoint)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedSplit, actualSplit) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedSplit, actualSplit)
	}
}

func TestSplit_OutOfRangeSplitPoint(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	splitPoint := 25

	expectedSplit := [][]int{}
	expectedErr := ErrIndexOutOfRange

	actualSplit, actualErr := Split(list, splitPoint)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedSplit, actualSplit) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedSplit, actualSplit)
	}
}

func TestSplit_ZeroSplitPoint(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	splitPoint := 0

	expectedSplit := [][]int{}
	expectedErr := ErrIndexOutOfRange

	actualSplit, actualErr := Split(list, splitPoint)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedSplit, actualSplit) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedSplit, actualSplit)
	}
}

func TestSlice(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	startIndex, endIndex := 3, 7

	expectedSlice := []int{3, 4, 5, 6, 7}
	var expectedErr error

	actualSlice, actualErr := Slice(list, startIndex, endIndex)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedSlice, actualSlice) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedSlice, actualSlice)
	}
}

func TestSlice_EmptyList(t *testing.T) {
	list := []int{}
	startIndex, endIndex := 3, 7

	expectedSlice := []int{}
	expectedErr := lists.ErrEmptyList

	actualSlice, actualErr := Slice(list, startIndex, endIndex)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedSlice, actualSlice) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedSlice, actualSlice)
	}
}

func TestSlice_OutOfBoundsStartIndex(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	startIndex, endIndex := -3, 7

	expectedSlice := []int{}
	expectedErr := ErrIndexOutOfRange

	actualSlice, actualErr := Slice(list, startIndex, endIndex)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedSlice, actualSlice) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedSlice, actualSlice)
	}
}

func TestSlice_OutOfBoundsEndIndex(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	startIndex, endIndex := 3, 17

	expectedSlice := []int{}
	expectedErr := ErrIndexOutOfRange

	actualSlice, actualErr := Slice(list, startIndex, endIndex)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedSlice, actualSlice) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedSlice, actualSlice)
	}
}

func TestRotate(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotationIndex := 3

	expectedRotate := []int{4, 5, 6, 7, 8, 1, 2, 3}
	var expectedErr error

	actualRotate, actualErr := Rotate(list, rotationIndex)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRotate, actualRotate) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRotate, actualRotate)
	}
}

func TestRotate_NegativeRotationIndex(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotationIndex := -2

	expectedRotate := []int{7, 8, 1, 2, 3, 4, 5, 6}
	var expectedErr error

	actualRotate, actualErr := Rotate(list, rotationIndex)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRotate, actualRotate) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRotate, actualRotate)
	}
}

func TestRotate_EmptyList(t *testing.T) {
	list := []int{}
	rotationIndex := 3

	expectedRotate := []int{}
	expectedErr := lists.ErrEmptyList

	actualRotate, actualErr := Rotate(list, rotationIndex)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRotate, actualRotate) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRotate, actualRotate)
	}
}

func TestRotate_ZeroRotationIndex(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotationIndex := 0

	expectedRotate := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var expectedErr error

	actualRotate, actualErr := Rotate(list, rotationIndex)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRotate, actualRotate) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRotate, actualRotate)
	}
}

func TestRotate_RotationIndexMultipleOfLength(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotationIndeces := []int{-64, -32, -16, -8, 8, 16, 32, 64}

	expectedRotate := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var expectedErr error

	for _, rotationIndex := range rotationIndeces {
		actualRotate, actualErr := Rotate(list, rotationIndex)

		if !reflect.DeepEqual(expectedErr, actualErr) {
			t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
		}

		if !reflect.DeepEqual(expectedRotate, actualRotate) {
			t.Fatalf("Expected: %v\n Got: %v\n", expectedRotate, actualRotate)
		}
	}

}

func TestRemoveAt(t *testing.T) {
	list := []int{1, 2, 3, 4}
	removalPoint := 2

	expectedRemoveAt := []int{1, 3, 4}
	var expectedErr error

	actualRemoveAt, actualErr := RemoveAt(list, removalPoint)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRemoveAt, actualRemoveAt) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRemoveAt, actualRemoveAt)
	}
}

func TestRemoveAt_EmptyList(t *testing.T) {
	list := []int{}
	removalPoint := 2

	expectedRemoveAt := []int{}
	expectedErr := lists.ErrEmptyList

	actualRemoveAt, actualErr := RemoveAt(list, removalPoint)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRemoveAt, actualRemoveAt) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRemoveAt, actualRemoveAt)
	}
}

func TestRemoveAt_NegativeSplitPoint(t *testing.T) {
	list := []int{1, 2, 3, 4}
	removalPoint := -2

	expectedRemoveAt := []int{}
	expectedErr := ErrIndexOutOfRange

	actualRemoveAt, actualErr := RemoveAt(list, removalPoint)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRemoveAt, actualRemoveAt) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRemoveAt, actualRemoveAt)
	}
}

func TestRemoveAt_OutOfRangeSplitPoint(t *testing.T) {
	list := []int{1, 2, 3, 4}
	removalPoint := 20

	expectedRemoveAt := []int{}
	expectedErr := ErrIndexOutOfRange

	actualRemoveAt, actualErr := RemoveAt(list, removalPoint)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRemoveAt, actualRemoveAt) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRemoveAt, actualRemoveAt)
	}
}

func TestRemoveAt_ZeroSplitPoint(t *testing.T) {
	list := []int{1, 2, 3, 4}
	removalPoint := 0

	expectedRemoveAt := []int{}
	expectedErr := ErrIndexOutOfRange

	actualRemoveAt, actualErr := RemoveAt(list, removalPoint)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRemoveAt, actualRemoveAt) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRemoveAt, actualRemoveAt)
	}
}
