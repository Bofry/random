package random

import (
	"math/rand"
)

// Uint64n returns a non-negative pseudo-random uint64 value in [0, n).
// Panics if n <= 0.
func (r *Random) Uint64n(n uint64) uint64 {
	if n <= 0 {
		panic("invalid argument to Uint64n")
	}
	return r.rand.Uint64() % n
}

// Uint32n returns a non-negative pseudo-random uint32 value in [0, n).
// Panics if n <= 0.
func (r *Random) Uint32n(n uint32) uint32 {
	if n <= 0 {
		panic("invalid argument to Uint32n")
	}
	return r.rand.Uint32() % n
}

// Float64n returns a pseudo-random float64 value in [0.0, n).
func (r *Random) Float64n(n float64) float64 {
	return n * r.rand.Float64()
}

// Float32n returns a pseudo-random float32 value in [0.0, n).
// Panics if n <= 0.
func (r *Random) Float32n(n float32) float32 {
	if n <= 0 {
		panic("invalid argument to Float32n")
	}
	return n * r.rand.Float32()
}

// ----------------------------------------------------------------------------
// Weighted Random Selection
// ----------------------------------------------------------------------------

// Float64w randomly picks an index in the range [0, len(w)-1] based on the weights in slice w.
// The probability of picking index i is w[i] / sum(w).
// Panics if w is empty or contains non-positive values.
func (r *Random) Float64w(w []float64) int {
	return weightedRandomIndexFloat64(r.rand, w)
}

// Float32w randomly picks an index in the range [0, len(w)-1] based on the weights in slice w.
// The probability of picking index i is w[i] / sum(w).
// Panics if w is empty or contains non-positive values.
func (r *Random) Float32w(w []float32) int {
	return weightedRandomIndexFloat32(r.rand, w)
}

// Uint64w randomly picks an index in the range [0, len(w)-1] based on the weights in slice w.
// The probability of picking index i is w[i] / sum(w).
// Panics if w is empty.
func (r *Random) Uint64w(w []uint64) int {
	return weightedRandomIndexUint64(r.rand, w)
}

// Uint32w randomly picks an index in the range [0, len(w)-1] based on the weights in slice w.
// The probability of picking index i is w[i] / sum(w).
// Panics if w is empty.
func (r *Random) Uint32w(w []uint32) int {
	return weightedRandomIndexUint32(r.rand, w)
}

// Int64w randomly picks an index in the range [0, len(w)-1] based on the weights in slice w.
// The probability of picking index i is w[i] / sum(w).
// Panics if w is empty or contains non-positive values.
func (r *Random) Int64w(w []int64) int {
	return weightedRandomIndexInt64(r.rand, w)
}

// Int32w randomly picks an index in the range [0, len(w)-1] based on the weights in slice w.
// The probability of picking index i is w[i] / sum(w).
// Panics if w is empty or contains non-positive values.
func (r *Random) Int32w(w []int32) int {
	return weightedRandomIndexInt32(r.rand, w)
}

// Intw randomly picks an index in the range [0, len(w)-1] based on the weights in slice w.
// The probability of picking index i is w[i] / sum(w).
// Panics if w is empty or contains non-positive values.
func (r *Random) Intw(w []int) int {
	return weightedRandomIndexInt(r.rand, w)
}

// weightedRandomIndexFloat64 selects a random index based on float64 weights.
func weightedRandomIndexFloat64(rng *rand.Rand, weights []float64) int {
	if len(weights) == 0 {
		panic("empty weights slice")
	}
	if len(weights) == 1 {
		return 0
	}

	var totalWeight float64
	for _, w := range weights {
		if w <= 0 {
			panic("weights must be positive")
		}
		totalWeight += w
	}

	target := rng.Float64() * totalWeight
	for i, w := range weights {
		if target < w {
			return i
		}
		target -= w
	}
	return len(weights) - 1 // Should not reach here if weights are valid
}

// weightedRandomIndexFloat32 selects a random index based on float32 weights.
func weightedRandomIndexFloat32(rng *rand.Rand, weights []float32) int {
	// Convert float32 weights to float64
	weights64 := make([]float64, len(weights))
	for i, w := range weights {
		weights64[i] = float64(w)
	}
	return weightedRandomIndexFloat64(rng, weights64)
}

// weightedRandomIndexUint64 selects a random index based on uint64 weights.
func weightedRandomIndexUint64(rng *rand.Rand, weights []uint64) int {
	if len(weights) == 0 {
		panic("empty weights slice")
	}
	if len(weights) == 1 {
		return 0
	}

	var totalWeight uint64
	for _, w := range weights {
		totalWeight += w
	}

	// Scale random value to the total weight range
	target := rng.Uint64() % totalWeight

	for i, w := range weights {
		if target < w {
			return i
		}
		target -= w
	}
	return len(weights) - 1 // Should not reach here if weights are valid
}

// weightedRandomIndexUint32 selects a random index based on uint32 weights.
func weightedRandomIndexUint32(rng *rand.Rand, weights []uint32) int {
	// Convert uint32 weights to uint64
	weights64 := make([]uint64, len(weights))
	for i, w := range weights {
		weights64[i] = uint64(w)
	}
	return weightedRandomIndexUint64(rng, weights64)
}

// weightedRandomIndexInt64 selects a random index based on int64 weights.
func weightedRandomIndexInt64(rng *rand.Rand, weights []int64) int {
	// Convert int64 weights to float64
	weightsF64 := make([]float64, len(weights))
	for i, w := range weights {
		if w < 0 {
			panic("weights must be positive")
		}
		weightsF64[i] = float64(w)
	}
	return weightedRandomIndexFloat64(rng, weightsF64)
}

// weightedRandomIndexInt32 selects a random index based on int32 weights.
func weightedRandomIndexInt32(rng *rand.Rand, weights []int32) int {
	// Convert int32 weights to int64
	weights64 := make([]int64, len(weights))
	for i, w := range weights {
		if w < 0 {
			panic("weights must be positive")
		}
		weights64[i] = int64(w)
	}
	return weightedRandomIndexInt64(rng, weights64)
}

// weightedRandomIndexInt selects a random index based on int weights.
func weightedRandomIndexInt(rng *rand.Rand, weights []int) int {
	// Convert int weights to float64
	weightsF64 := make([]float64, len(weights))
	for i, w := range weights {
		if w < 0 {
			panic("weights must be positive")
		}
		weightsF64[i] = float64(w)
	}
	return weightedRandomIndexFloat64(rng, weightsF64)
}
