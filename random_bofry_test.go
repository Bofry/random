package random_test

import (
	"math/rand"
	"testing"

	"github.com/bofry/random" // 請確認套件路徑是否正確
)

const testRound = 1e8 // Number of rounds for statistical testing

func TestWeightedRandomIndex(t *testing.T) {
	seed := int64(5489)
	rng := random.New(rand.NewSource(seed)) // Create a random.Random instance

	// Test cases for weightedRandomIndex with different types and weights
	testCases := []struct {
		name                 string
		weights              []float64
		expectedDistribution map[int]float64
	}{
		{
			"Float64Weights",
			[]float64{2, 2, 2, 4},
			map[int]float64{0: 0.2, 1: 0.2, 2: 0.2, 3: 0.4},
		},
		{
			"Float32Weights",
			[]float64{2, 2, 2, 4},
			map[int]float64{0: 0.2, 1: 0.2, 2: 0.2, 3: 0.4},
		},
		{
			"Uint64Weights",
			[]float64{1, 2, 3, 4},
			map[int]float64{0: 0.1, 1: 0.2, 2: 0.3, 3: 0.4},
		},
		{
			"Uint32Weights",
			[]float64{1, 2, 3, 4},
			map[int]float64{0: 0.1, 1: 0.2, 2: 0.3, 3: 0.4},
		},
		{
			"Int64Weights",
			[]float64{1, 2, 3, 4},
			map[int]float64{0: 0.1, 1: 0.2, 2: 0.3, 3: 0.4},
		},
		{
			"Int32Weights",
			[]float64{1, 2, 3, 4},
			map[int]float64{0: 0.1, 1: 0.2, 2: 0.3, 3: 0.4},
		},
		{
			"IntWeights",
			[]float64{1, 2, 3, 4},
			map[int]float64{0: 0.1, 1: 0.2, 2: 0.3, 3: 0.4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Count the occurrences of each index
			counts := make(map[int]int)
			for i := 0; i < testRound; i++ {
				var index int
				switch tc.name {
				case "Float64Weights":
					index = rng.Float64w(tc.weights)
				case "Float32Weights":
					index = rng.Float32w(float32sToFloat64s(tc.weights))
				case "Uint64Weights":
					index = rng.Uint64w(uint64sToUint64s(tc.weights))
				case "Uint32Weights":
					index = rng.Uint32w(uint32sToUint32s(tc.weights))
				case "Int64Weights":
					index = rng.Int64w(int64sToint64s(tc.weights))
				case "Int32Weights":
					index = rng.Int32w(int32sToint32s(tc.weights))
				case "IntWeights":
					index = rng.Intw(intsToInts(tc.weights))
				}
				counts[index]++
			}

			// Check if the observed distribution is close to the expected distribution
			for index, expectedProb := range tc.expectedDistribution {
				observedProb := float64(counts[index]) / testRound
				tolerance := 0.01 // Allow 1% deviation
				if observedProb < expectedProb-tolerance || observedProb > expectedProb+tolerance {
					t.Errorf("Unexpected distribution for index %d: expected %.2f, got %.2f", index, expectedProb, observedProb)
				}
			}
		})
	}
}

// TestPanicCases tests panic conditions for weighted random functions.
func TestPanicCases(t *testing.T) {
	rng := random.New(rand.NewSource(seed))

	// Test cases that should panic
	panicTestCases := []struct {
		name     string
		function func()
	}{
		{"Float64w_EmptySlice", func() { rng.Float64w([]float64{}) }},
		{"Float64w_NegativeWeight", func() { rng.Float64w([]float64{1, -2, 3}) }},
		{"Float32w_EmptySlice", func() { rng.Float32w([]float32{}) }},
		{"Float32w_NegativeWeight", func() { rng.Float32w([]float32{1, -2, 3}) }},
		{"Uint64w_EmptySlice", func() { rng.Uint64w([]uint64{}) }},
		{"Uint32w_EmptySlice", func() { rng.Uint32w([]uint32{}) }},
		{"Int64w_EmptySlice", func() { rng.Int64w([]int64{}) }},
		{"Int64w_NegativeWeight", func() { rng.Int64w([]int64{1, -2, 3}) }},
		{"Int32w_EmptySlice", func() { rng.Int32w([]int32{}) }},
		{"Int32w_NegativeWeight", func() { rng.Int32w([]int32{1, -2, 3}) }},
		{"Intw_EmptySlice", func() { rng.Intw([]int{}) }},
		{"Intw_NegativeWeight", func() { rng.Intw([]int{1, -2, 3}) }},
	}

	for _, tc := range panicTestCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("%s did not panic as expected", tc.name)
				}
			}()
			tc.function()
		})
	}
}

// TestRandomNumberBounds tests if the generated numbers are within the expected bounds.
func TestRandomNumberBounds(t *testing.T) {
	rng := random.New(rand.NewSource(seed))
	// test cases
	intCases := []struct {
		name     string
		function func(int64) int64
		low      int64
		high     int64
		count    int
	}{
		{"Int63n", rng.Int63n, 1, 100, 1000},
		{"Int31n", func(n int64) int64 { return int64(rng.Int31n(int32(n))) }, 1, 100, 1000},
		{"Intn", func(n int64) int64 { return int64(rng.Intn(int(n))) }, 1, 100, 1000},
	}
	for _, c := range intCases {
		t.Run(c.name, func(t *testing.T) {
			for i := 0; i < c.count; i++ {
				num := c.function(c.high)
				if num < c.low || num >= c.high {
					t.Errorf("%s(%d) returned out of range value: %d", c.name, c.high, num)
				}
			}
		})
	}

	// Float Test Cases
	floatCases := []struct {
		name     string
		function func(float64) float64
		low      float64
		high     float64
		count    int
	}{
		{"Float64n", rng.Float64n, 1.0, 100.0, 1000},
		{"Float32n", func(n float64) float64 { return float64(rng.Float32n(float32(n))) }, 1.0, 100.0, 1000},
	}

	for _, c := range floatCases {
		t.Run(c.name, func(t *testing.T) {
			for i := 0; i < c.count; i++ {
				num := c.function(c.high)
				if num < c.low || num >= c.high {
					t.Errorf("%s(%v) returned out of range value: %v", c.name, c.high, num)
				}
			}
		})
	}
}

// Helper functions for type conversion
func float32sToFloat64s(input []float64) []float32 {
	output := make([]float32, len(input))
	for i, v := range input {
		output[i] = float32(v)
	}
	return output
}

func uint64sToUint64s(input []float64) []uint64 {
	output := make([]uint64, len(input))
	for i, v := range input {
		output[i] = uint64(v)
	}
	return output
}

func uint32sToUint32s(input []float64) []uint32 {
	output := make([]uint32, len(input))
	for i, v := range input {
		output[i] = uint32(v)
	}
	return output
}

func int64sToint64s(input []float64) []int64 {
	output := make([]int64, len(input))
	for i, v := range input {
		output[i] = int64(v)
	}
	return output
}

func int32sToint32s(input []float64) []int32 {
	output := make([]int32, len(input))
	for i, v := range input {
		output[i] = int32(v)
	}
	return output
}

func intsToInts(input []float64) []int {
	output := make([]int, len(input))
	for i, v := range input {
		output[i] = int(v)
	}
	return output
}
