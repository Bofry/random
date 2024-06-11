package random

import (
	"math/rand"
)

var _ rand.Source64 = (*Random)(nil) // Ensures Rand complies with rand.Source64

// Random is a random number generator.
type Random struct {
	rand *rand.Rand
}

// New creates a new Random instance.
func New(src rand.Source) *Random {
	return &Random{
		rand: rand.New(src),
	}
}

// Seed uses the provided seed value to initialize the generator to a deterministic state.
// Seed should not be called concurrently with any other Rand method.
func (r *Random) Seed(seed int64) {
	r.rand.Seed(seed)
}

// Int63 returns a non-negative pseudo-random 64-bit integer as an int64.
func (r *Random) Int63() int64 {
	return r.rand.Int63()
}

// Uint64 returns a non-negative pseudo-random 64-bit integer as a uint64.
func (r *Random) Uint64() uint64 {
	return r.rand.Uint64()
}

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func (r *Random) Uint32() uint32 {
	return r.rand.Uint32()
}

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func (r *Random) Int31() int32 {
	return r.rand.Int31()
}

// Int returns a non-negative pseudo-random int.
func (r *Random) Int() int {
	return r.rand.Int()
}

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Random) Int63n(n int64) int64 {
	return r.rand.Int63n(n)
}

// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Random) Int31n(n int32) int32 {
	return r.rand.Int31n(n)
}

// Intn returns, as an int, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Random) Intn(n int) int {
	return r.rand.Intn(n)
}

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func (r *Random) Float64() float64 {
	return r.rand.Float64()
}

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func (r *Random) Float32() float32 {
	return r.rand.Float32()
}

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n).
func (r *Random) Perm(n int) []int {
	return r.rand.Perm(n)
}

// Shuffle pseudo-randomizes the order of elements using the provided swap function.
func (r *Random) Shuffle(n int, swap func(i, j int)) {
	r.rand.Shuffle(n, swap)
}

// Read generates len(p) random bytes and writes them into p. It always returns len(p) and a nil error.
func (r *Random) Read(p []byte) (n int, err error) {
	return r.rand.Read(p)
}
