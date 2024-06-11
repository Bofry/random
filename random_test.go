package random_test

import (
	"math/rand"
	"testing"

	"github.com/bofry/random/mt19937"

	"github.com/bofry/random"
)

var (
	seed             = int64(5489)
	rng              = random.New(rand.NewSource(seed))
	rng_safe         = random.NewthreadSafeRandom(rand.NewSource(seed))
	rng_mt19937      = random.New(mt19937.New())
	rng_safe_mt19937 = random.NewthreadSafeRandom(mt19937.New())
)

func TestRandom(t *testing.T) {
	seed := int64(5489)
	rng := random.New(rand.NewSource(seed))
	rngMT := random.New(mt19937.New())

	testCases := []struct {
		name     string
		rng      *random.Random // Random number generator to test
		function func() interface{}
		check    func(interface{}) bool // Function to check the result
	}{
		{"Int63", rng, func() interface{} { return rng.Int63() }, func(v interface{}) bool { return v.(int64) >= 0 }},
		{"Uint64", rng, func() interface{} { return rng.Uint64() }, func(v interface{}) bool { return v.(uint64) >= 0 }},
		{"Uint32", rng, func() interface{} { return rng.Uint32() }, func(v interface{}) bool { return v.(uint32) >= 0 }},
		{"Int31", rng, func() interface{} { return rng.Int31() }, func(v interface{}) bool { return v.(int32) >= 0 }},
		{"Int", rng, func() interface{} { return rng.Int() }, func(v interface{}) bool { return v.(int) >= 0 }},
		{"Float64", rng, func() interface{} { return rng.Float64() }, func(v interface{}) bool { return v.(float64) >= 0 && v.(float64) < 1 }},
		{"Float32", rng, func() interface{} { return rng.Float32() }, func(v interface{}) bool { return v.(float32) >= 0 && v.(float32) < 1 }},
		// Add more test cases as needed
		{"Int63 (MT19937)", rngMT, func() interface{} { return rngMT.Int63() }, func(v interface{}) bool { return v.(int64) >= 0 }},
		{"Uint64 (MT19937)", rngMT, func() interface{} { return rngMT.Uint64() }, func(v interface{}) bool { return v.(uint64) >= 0 }},
		{"Uint32 (MT19937)", rngMT, func() interface{} { return rngMT.Uint32() }, func(v interface{}) bool { return v.(uint32) >= 0 }},
		{"Int31 (MT19937)", rngMT, func() interface{} { return rngMT.Int31() }, func(v interface{}) bool { return v.(int32) >= 0 }},
		{"Int (MT19937)", rngMT, func() interface{} { return rngMT.Int() }, func(v interface{}) bool { return v.(int) >= 0 }},
		{"Float64 (MT19937)", rngMT, func() interface{} { return rngMT.Float64() }, func(v interface{}) bool { return v.(float64) >= 0 && v.(float64) < 1 }},
		{"Float32 (MT19937)", rngMT, func() interface{} { return rngMT.Float32() }, func(v interface{}) bool { return v.(float32) >= 0 && v.(float32) < 1 }},
		// Add more MT19937 test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Check if the return value meets the expectation
			if !tc.check(tc.function()) {
				t.Errorf("%s failed", tc.name)
			}
		})
	}

	// Test panic cases for Intn family functions
	panicTestCases := []struct {
		name     string
		rng      *random.Random
		function func()
	}{
		{"Int63n", rng, func() { rng.Int63n(-1) }},
		{"Int31n", rng, func() { rng.Int31n(-1) }},
		{"Intn", rng, func() { rng.Intn(-1) }},
		{"Int63n (MT19937)", rngMT, func() { rngMT.Int63n(-1) }},
		{"Int31n (MT19937)", rngMT, func() { rngMT.Int31n(-1) }},
		{"Intn (MT19937)", rngMT, func() { rngMT.Intn(-1) }},
		// Add more panic test cases as needed
	}

	for _, tc := range panicTestCases {
		t.Run(tc.name+"_Panic", func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("%s did not panic as expected", tc.name)
				}
			}()
			tc.function()
		})
	}
}

func TestPermAndShuffle(t *testing.T) {
	// Test Perm
	rng := random.New(rand.NewSource(seed))
	for i := 0; i < 100; i++ {
		n := rng.Perm(5)
		if len(n) != 5 {
			t.Errorf("Perm: expected length 5, got %d", len(n))
		}
		for _, v := range n {
			if v < 0 || v >= 5 {
				t.Errorf("Perm: value out of range [0, 5): %d", v)
			}
		}
	}

	// Test Shuffle
	rng.Seed(seed)
	n := rng.Perm(5)
	original := make([]int, len(n))
	copy(original, n)
	rng.Shuffle(len(n), func(i, j int) { n[i], n[j] = n[j], n[i] })
	if isEqual(original, n) {
		t.Errorf("Shuffle: did not change the order of the slice")
	}
}

func TestRead(t *testing.T) {
	// Test Read
	rng := random.New(rand.NewSource(seed))
	n := make([]byte, 5)
	_, err := rng.Read(n)
	if err != nil {
		t.Errorf("Read: unexpected error: %v", err)
	}
}

// Helper function to compare two slices for equality
func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
