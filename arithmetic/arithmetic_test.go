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

		actualPrimeFactors, actualErr := PrimeFactors(Number)

		if !reflect.DeepEqual(expectedErr, actualErr) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
		}

		if !reflect.DeepEqual(expectedPrimeFactors, actualPrimeFactors) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedPrimeFactors, actualPrimeFactors)
		}
	}
}

func TestPrimeFactors_NegativeNumber(t *testing.T) {
	Number := -17

	expectedPrimeFactors := []int{}
	expectedErr := definitions.ErrNotPositiveNumber

	actualPrimeFactors, actualErr := PrimeFactors(Number)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedPrimeFactors, actualPrimeFactors) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedPrimeFactors, actualPrimeFactors)
	}
}

func TestPrimeFactors_OneTwoThree(t *testing.T) {
	Numbers := []int{1, 2, 3}
	expectedPrimeFactorsList := [][]int{{1}, {2}, {3}}

	var expectedErr error

	for index := 0; index < len(Numbers); index++ {
		Number := Numbers[index]
		expectedPrimeFactors := expectedPrimeFactorsList[index]

		actualPrimeFactors, actualErr := PrimeFactors(Number)

		if !reflect.DeepEqual(expectedErr, actualErr) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
		}

		if !reflect.DeepEqual(expectedPrimeFactors, actualPrimeFactors) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedPrimeFactors, actualPrimeFactors)
		}
	}
}

func TestPrimeFactorsMult(t *testing.T) {
	Numbers := []int{2, 33, 68, 97, 177, 315}
	expectedPrimeFactorsMultList := [][][]int{{{2, 1}}, {{3, 1}, {11, 1}}, {{2, 2}, {17, 1}}, {{97, 1}}, {{3, 1}, {59, 1}}, {{3, 2}, {5, 1}, {7, 1}}}

	var expectedErr error

	for index := 0; index < len(Numbers); index++ {
		Number := Numbers[index]
		expectedPrimeFactorsMult := expectedPrimeFactorsMultList[index]

		actualPrimeFactorsMult, actualErr := PrimeFactorsMult(Number)

		if !reflect.DeepEqual(expectedErr, actualErr) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
		}

		if !reflect.DeepEqual(expectedPrimeFactorsMult, actualPrimeFactorsMult) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedPrimeFactorsMult, actualPrimeFactorsMult)
		}
	}
}

func TestPrimeFactorsMult_NegativeNumber(t *testing.T) {
	Number := -17

	expectedPrimeFactorsMult := [][]int{}
	expectedErr := definitions.ErrNotPositiveNumber

	actualPrimeFactorsMult, actualErr := PrimeFactorsMult(Number)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedPrimeFactorsMult, actualPrimeFactorsMult) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedPrimeFactorsMult, actualPrimeFactorsMult)
	}
}

func TestPrimeFactorsMult_OneTwoThree(t *testing.T) {
	Numbers := []int{1, 2, 3}
	expectedPrimeFactorsMultList := [][][]int{{{1, 1}}, {{2, 1}}, {{3, 1}}}

	var expectedErr error

	for index := 0; index < len(Numbers); index++ {
		Number := Numbers[index]
		expectedPrimeFactorsMult := expectedPrimeFactorsMultList[index]

		actualPrimeFactorsMult, actualErr := PrimeFactorsMult(Number)

		if !reflect.DeepEqual(expectedErr, actualErr) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
		}

		if !reflect.DeepEqual(expectedPrimeFactorsMult, actualPrimeFactorsMult) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedPrimeFactorsMult, actualPrimeFactorsMult)
		}
	}
}

func TestTotientImproved(t *testing.T) {
	Ms := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedTotientsImproved := []int{1, 1, 2, 2, 4, 2, 6, 4, 6, 4}

	var expectedErr error

	for index := 0; index < len(Ms); index++ {
		m := Ms[index]
		expectedTotientImproved := expectedTotientsImproved[index]

		actualTotientImproved, actualErr := Totient(m)

		if !reflect.DeepEqual(expectedErr, actualErr) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
		}

		if !reflect.DeepEqual(expectedTotientImproved, actualTotientImproved) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedTotientImproved, actualTotientImproved)
		}
	}
}

func TestTotientImproved_NegativeM(t *testing.T) {
	m := -10

	expectedTotientImproved := 0
	expectedErr := definitions.ErrNotPositiveNumber

	actualTotientImproved, actualErr := Totient(m)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedTotientImproved, actualTotientImproved) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedTotientImproved, actualTotientImproved)
	}
}

func TestTotientImproved_ZeroM(t *testing.T) {
	m := 0

	expectedTotientImproved := 0
	expectedErr := definitions.ErrNotPositiveNumber

	actualTotientImproved, actualErr := Totient(m)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedTotientImproved, actualTotientImproved) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedTotientImproved, actualTotientImproved)
	}
}

func BenchmarkTotient10090(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Totient(10090)
	}
}

func BenchmarkTotientImproved10090(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TotientImproved(10090)
	}
}
func TestPrimesR(t *testing.T) {
	m, n := 10, 20

	expectedPrimesR := []int{11, 13, 17, 19}
	var expectedErr error

	actualPrimesR, actualErr := PrimesR(m, n)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedPrimesR, actualPrimesR) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedPrimesR, actualPrimesR)
	}
}

func TestPrimesR_ZeroIndex(t *testing.T) {
	m, n := 10, 0

	expectedPrimesR := []int{}
	expectedErr := definitions.ErrNotPositiveNumber

	actualPrimesR, actualErr := PrimesR(m, n)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedPrimesR, actualPrimesR) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedPrimesR, actualPrimesR)
	}
}

func TestPrimesR_NegativeIndex(t *testing.T) {
	m, n := -10, -20

	expectedPrimesR := []int{}
	expectedErr := definitions.ErrNotPositiveNumber

	actualPrimesR, actualErr := PrimesR(m, n)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedPrimesR, actualPrimesR) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedPrimesR, actualPrimesR)
	}
}

func TestPrimesR_MGreaterThanN(t *testing.T) {
	m, n := 20, 10

	expectedPrimesR := []int{11, 13, 17, 19}
	var expectedErr error

	actualPrimesR, actualErr := PrimesR(m, n)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedPrimesR, actualPrimesR) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedPrimesR, actualPrimesR)
	}
}

func TestGoldbach(t *testing.T) {
	numbers := []int{10, 14, 28, 1382, 1928}

	expectedGoldbachList := [][]int{{3, 7}, {3, 11}, {5, 23}, {61, 1321}, {61, 1867}}
	var expectedErr error

	for index := 0; index < len(numbers); index++ {
		number := numbers[index]
		expectedGoldbach := expectedGoldbachList[index]

		actualGoldbach, actualErr := Goldbach(number)

		if !reflect.DeepEqual(expectedErr, actualErr) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
		}

		if !reflect.DeepEqual(expectedGoldbach, actualGoldbach) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedGoldbach, actualGoldbach)
		}
	}
}

func TestGoldbach_Negative(t *testing.T) {
	number := -28

	expectedGoldbach := []int{}
	expectedErr := definitions.ErrAtLeastTwo

	actualGoldbach, actualErr := Goldbach(number)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedGoldbach, actualGoldbach) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedGoldbach, actualGoldbach)
	}
}

func TestGoldbach_LessThan3(t *testing.T) {
	numbers := []int{2, 1, 0}

	expectedGoldbach := []int{}
	expectedErr := definitions.ErrAtLeastTwo

	for index := 0; index < len(numbers); index++ {
		number := numbers[index]

		actualGoldbach, actualErr := Goldbach(number)

		if !reflect.DeepEqual(expectedErr, actualErr) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
		}

		if !reflect.DeepEqual(expectedGoldbach, actualGoldbach) {
			t.Fatalf("Expected: %v\n Actual: %v\n", expectedGoldbach, actualGoldbach)
		}
	}
}

func TestGoldbach_Odd(t *testing.T) {
	number := 11

	expectedGoldbach := []int{}
	expectedErr := definitions.ErrOddNumber

	actualGoldbach, actualErr := Goldbach(number)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedGoldbach, actualGoldbach) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedGoldbach, actualGoldbach)
	}
}

func TestGoldbachList(t *testing.T) {
	lower, upper := 9, 20

	expectedGoldbachList := [][]int{{3, 7}, {5, 7}, {3, 11}, {3, 13}, {5, 13}, {3, 17}}
	var expectedErr error

	actualGoldbachList, actualErr := GoldbachList(lower, upper)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedGoldbachList, actualGoldbachList) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedGoldbachList, actualGoldbachList)
	}
}

func TestGoldbachList_Zero(t *testing.T) {
	lower, upper := 9, 0

	expectedGoldbachList := [][]int{}
	expectedErr := definitions.ErrAtLeastTwo

	actualGoldbachList, actualErr := GoldbachList(lower, upper)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedGoldbachList, actualGoldbachList) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedGoldbachList, actualGoldbachList)
	}
}

func TestGoldbachList_Negative(t *testing.T) {
	lower, upper := 9, -20

	expectedGoldbachList := [][]int{}
	expectedErr := definitions.ErrAtLeastTwo

	actualGoldbachList, actualErr := GoldbachList(lower, upper)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedGoldbachList, actualGoldbachList) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedGoldbachList, actualGoldbachList)
	}
}

func TestGoldbachList_LessThan3(t *testing.T) {
	lower, upper := 9, 2

	expectedGoldbachList := [][]int{}
	expectedErr := definitions.ErrAtLeastTwo

	actualGoldbachList, actualErr := GoldbachList(lower, upper)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedGoldbachList, actualGoldbachList) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedGoldbachList, actualGoldbachList)
	}
}

func TestGoldbachList_UpperLessThanLower(t *testing.T) {
	lower, upper := 20, 9

	expectedGoldbachList := [][]int{{3, 7}, {5, 7}, {3, 11}, {3, 13}, {5, 13}, {3, 17}}
	var expectedErr error

	actualGoldbachList, actualErr := GoldbachList(lower, upper)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedGoldbachList, actualGoldbachList) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedGoldbachList, actualGoldbachList)
	}
}
