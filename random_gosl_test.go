package random_test

import (
	"reflect"
	"testing"
)

// TestRangeInt63 tests Int63r for different ranges and data types.
func TestRangeInt63(t *testing.T) {
	testCases := []struct {
		name  string
		low   int64
		high  int64
		count int
	}{
		{"Positive Range", 0, 100, 1000},
		{"Negative Range", -100, -1, 1000},
		{"Zero Range", 0, 0, 100},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i < tc.count; i++ {
				num := rng.Int63r(tc.low, tc.high)
				if num < tc.low || num > tc.high {
					t.Errorf("Int63r(%d, %d) returned out of range value: %d", tc.low, tc.high, num)
				}
			}
		})
	}
}

// TestInt63s tests the Int63s function to ensure it generates int64 slices correctly.
func TestInt63s(t *testing.T) {
	// Test with an empty slice
	values := make([]int64, 0)
	rng.Int63s(values, 0, 10)
	if len(values) != 0 {
		t.Error("Int63s should not modify an empty slice")
	}

	// Test with a non-empty slice
	values = make([]int64, 5)
	rng.Int63s(values, 0, 10)
	for _, v := range values {
		if v < 0 || v > 10 {
			t.Errorf("Int63s generated a value out of range: %d", v)
		}
	}
}

// TestInt63Shuffle tests the Int63Shuffle function to ensure it shuffles int64 slices.
func TestInt63Shuffle(t *testing.T) {
	original := []int64{1, 2, 3, 4, 5}
	shuffled := make([]int64, len(original))
	copy(shuffled, original)
	rng.Int63Shuffle(shuffled)
	if reflect.DeepEqual(original, shuffled) {
		t.Error("Int63Shuffle did not shuffle the slice")
	}
}

// Similar tests for Uint32r, Uint32s, Uint32Shuffle, Uint64r, Uint64s, Uint64Shuffle...
// Similar tests for Int31r, Int31s, Int31Shuffle, Intr, Ints, IntShuffle...
// Similar tests for Float64r, Float64s, Float64Shuffle, Float32r, Float32s, Float32Shuffle...

// TestFlipCoin tests the FlipCoin function for different probabilities.
func TestFlipCoin(t *testing.T) {
	rng.Seed(seed)

	// Test with p = 0
	for i := 0; i < 100; i++ {
		if rng.FlipCoin(0) {
			t.Error("FlipCoin(0) should always return false")
		}
	}

	// Test with p = 1
	for i := 0; i < 100; i++ {
		if !rng.FlipCoin(1) {
			t.Error("FlipCoin(1) should always return true")
		}
	}

	// Test with p = 0.5
	heads := 0
	tails := 0
	for i := 0; i < 10000; i++ {
		if rng.FlipCoin(0.5) {
			heads++
		} else {
			tails++
		}
	}
	// Allow some deviation due to randomness
	if heads < 4500 || heads > 5500 {
		t.Errorf("FlipCoin(0.5) head/tail ratio is too skewed: %d heads, %d tails", heads, tails)
	}
}
