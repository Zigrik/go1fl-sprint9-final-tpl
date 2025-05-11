package main

// Пишите тесты в этом файле
import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		size   int
		answer int
	}{
		{-1, 0},
		{0, 0},
		{1, 1},
		{5, 5},
		{10000, 10000},
	}
	for _, v := range tests {
		arr := generateRandomElements(v.size)
		require.Equal(t, v.answer, len(arr))
		if v.answer > 0 {
			require.Greater(t, slices.Min(arr), 0, "all generated numbers must be positive")
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		array  []int
		answer int
	}{
		{[]int{}, 0},
		{[]int{7}, 7},
		{[]int{1, 2, 3, 4, 567, 435643276, 1214, 1231243, 34}, 435643276},
		{[]int{0}, 0},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1},
	}
	for _, v := range tests {
		maxTestNum := maximum(v.array)
		require.Equal(t, v.answer, maxTestNum)
	}
}
