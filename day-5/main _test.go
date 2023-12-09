package main

import (
	"math"
	"slices"
	"testing"
)

func CompilePieceWiseFunction(source []FuncPiece) func(input int) int {
	return func(input int) int {
		for _, piece := range source {
			if input >= piece.from && input < piece.to {
				return input + piece.a
			}
		}
		return 0
	}
}
func testCaseMapPieceRangeToTargetDomain(
	t *testing.T,
	name string,
	srcPiece FuncPiece,
	targetF []FuncPiece,
	expectedF []FuncPiece,
) {
	t.Run(name, func(t *testing.T) {
		size := srcPiece.to - srcPiece.from
		seg := make([]FuncPiece, 0)

		MapPieceRangeToTargetDomain(srcPiece, targetF, &seg)
		srcFn := CompilePieceWiseFunction([]FuncPiece{srcPiece})

		targetFn := CompilePieceWiseFunction(targetF)
		resFn := CompilePieceWiseFunction(seg)
		expectedFn := CompilePieceWiseFunction(seg)

		composedResSlice := make([]int, size)
		inputs := make([]int, size)
		expectedResSlice := make([]int, size)
		resultSlice := make([]int, size)

		for i := 0; i < size; i++ {
			inputs[i] = i
			composedResSlice[i] = targetFn(srcFn(srcPiece.from + i))
			expectedResSlice[i] = expectedFn(srcPiece.from + i)
			resultSlice[i] = resFn(srcPiece.from + i)
		}

		if !slices.Equal(seg, expectedF) {
			t.Error("Expected to match", seg, expectedF)
			t.Log("inputs  ", inputs)

			t.Log("composed", composedResSlice)
			t.Log("expected", expectedResSlice)
			t.Log("result  ", resultSlice)
		}
	})
}

func TestMapPieceRangeToTargetDomain(t *testing.T) {
	testCaseMapPieceRangeToTargetDomain(
		t,
		"A",
		FuncPiece{0, 10, 0},
		[]FuncPiece{{0, 20, 0}, {20, 30, 10}, {30, math.MaxInt, 0}},
		[]FuncPiece{{0, 10, 0}},
	)
	testCaseMapPieceRangeToTargetDomain(
		t,
		"B",
		FuncPiece{0, 20, 0},
		[]FuncPiece{{0, 10, 0}, {10, 30, 10}, {30, math.MaxInt, 0}},
		[]FuncPiece{{0, 10, 0}, {10, 20, 10}},
	)
	testCaseMapPieceRangeToTargetDomain(
		t,
		"Ba",
		FuncPiece{0, 20, 3},
		[]FuncPiece{{0, 10, 0}, {10, 30, 10}, {30, math.MaxInt, 0}},
		[]FuncPiece{{0, 7, 3}, {7, 20, 13}},
	)
	testCaseMapPieceRangeToTargetDomain(
		t,
		"C",
		FuncPiece{5, 15, 0},
		[]FuncPiece{{0, 10, 0}, {10, 30, 10}, {30, math.MaxInt, 0}},
		[]FuncPiece{{5, 10, 0}, {10, 15, 10}},
	)
	testCaseMapPieceRangeToTargetDomain(
		t,
		"D",
		FuncPiece{5, 35, 0},
		[]FuncPiece{{0, 10, 0}, {10, 30, 10}, {30, math.MaxInt, 0}},
		[]FuncPiece{{5, 10, 0}, {10, 30, 10}, {30, 35, 0}},
	)
	testCaseMapPieceRangeToTargetDomain(
		t,
		"E",
		FuncPiece{5, 15, 1},
		[]FuncPiece{{0, 10, 0}, {10, 30, 10}, {30, math.MaxInt, 0}},
		[]FuncPiece{{5, 9, 1}, {9, 15, 11}},
	)
	testCaseMapPieceRangeToTargetDomain(
		t,
		"F",
		FuncPiece{5, 35, 1},
		[]FuncPiece{{0, 10, 0}, {10, 30, 10}, {30, math.MaxInt, 0}},
		[]FuncPiece{{5, 9, 1}, {9, 29, 11}, {29, 35, 1}},
	)
	testCaseMapPieceRangeToTargetDomain(
		t,
		"G",
		FuncPiece{10, 20, 10},
		[]FuncPiece{{0, 20, 0}, {20, 30, 10}, {30, math.MaxInt, 0}},
		[]FuncPiece{{10, 20, 20}},
	)
	testCaseMapPieceRangeToTargetDomain(
		t,
		"H",
		FuncPiece{11, 20, 0},
		[]FuncPiece{{0, 20, 0}, {20, 30, 10}, {30, math.MaxInt, 0}},
		[]FuncPiece{{11, 20, 0}},
	)
	testCaseMapPieceRangeToTargetDomain(
		t,
		"I",
		FuncPiece{20, 60, 2},
		[]FuncPiece{{0, 20, 0}, {20, 30, 10}, {30, math.MaxInt, 0}},
		[]FuncPiece{{20, 28, 12}, {28, 60, 2}},
	)
}

func testCaseComposePieceWiseFuncs(t *testing.T, sources []FuncPiece, targets []FuncPiece) {
	results := ComposePieceWiseFuncs(targets, sources)
	sourceCompiled := CompilePieceWiseFunction(sources)
	targetCompiled := CompilePieceWiseFunction(targets)
	resultsCompiled := CompilePieceWiseFunction(results)

	mergedRes := make([]int, 40)
	testRes := make([]int, 40)
	sourceRed := make([]int, 40)
	input := make([]int, 40)
	for i := 0; i < 40; i++ {
		input[i] = i
		sourceRed[i] = sourceCompiled(i)
		mergedRes[i] = targetCompiled(sourceRed[i])
		testRes[i] = resultsCompiled(i)
	}

	if !slices.Equal(mergedRes, testRes) {
		t.Error("Expected to match")
		t.Log("input   ", input)
		t.Log("expected", mergedRes)
		t.Log("results ", testRes)

	}

}

func TestComposePieceWiseFuncs(t *testing.T) {
	t.Run("A", func(t *testing.T) {
		testCaseComposePieceWiseFuncs(t,
			[]FuncPiece{{0, 10, 0}, {10, 20, 10}, {20, math.MaxInt, 0}},
			[]FuncPiece{{0, 20, 0}, {20, 30, 10}, {30, math.MaxInt, 0}},
		)
	})
	t.Run("B", func(t *testing.T) {
		testCaseComposePieceWiseFuncs(t,
			[]FuncPiece{{0, 10, 0}, {10, 21, 10}, {21, math.MaxInt, 0}},
			[]FuncPiece{{0, 20, 0}, {20, 30, 10}, {30, math.MaxInt, 0}},
		)
	})
	t.Run("C", func(t *testing.T) {
		testCaseComposePieceWiseFuncs(t,
			[]FuncPiece{{0, 10, 0}, {10, 20, 10}, {20, math.MaxInt - 3, 2}},
			[]FuncPiece{{0, 20, 0}, {20, 30, 10}, {30, math.MaxInt - 3, 0}},
		)
	})
	t.Run("D", func(t *testing.T) {
		testCaseComposePieceWiseFuncs(t,
			[]FuncPiece{{0, 10, 0}, {10, 30, 10}, {30, math.MaxInt, 2}},
			[]FuncPiece{{0, 20, 0}, {20, 30, 10}, {30, math.MaxInt, 0}},
		)
	})
}

func BenchmarkMain(b *testing.B) {
	main()

}
