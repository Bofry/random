package random

import (
	"math/rand"
)

// Random -
// A Random Number Generator
type Random struct {
	rand *rand.Rand
}

// New -
// 	returns:
//	 --	a new Random.
func New(src rand.Source) *Random {
	return &Random{
		rand: rand.New(src),
	}
}

// Seed -
// Seed uses the provided seed value to initialize the generator to a deterministic state.
// Seed should not be called concurrently with any other Rand method.
func (r *Random) Seed(seed int64) {
	r.rand.Seed(seed)
}

// Int63 -
// 	returns:
//	 --	an int64 between [-2^63, 2^63-1].
func (r *Random) Int63() int64 {
	val := r.rand.Int63()
	return val
}

// Uint64 -
// 	returns:
//	 --	an Uint64 between [0, 2^64-1].
func (r *Random) Uint64() uint64 {
	return r.rand.Uint64()
}

// Uint32 -
// 	returns:
//	 --	an Uint32 between [0, 2^32-1].
func (r *Random) Uint32() uint32 {
	return r.rand.Uint32()
}

// Int31 -
// 	returns:
//	 --	an int64 between [-2^31, 2^31-1].
func (r *Random) Int31() int32 {
	return r.rand.Int31()
}

// Int -
// 	returns:
//	 --	an Int between [-2^31, 2^31-1].
func (r *Random) Int() int {
	return r.rand.Int()
}

// Int63n -
//	input:
//	 n	-- upper limit.
// 	returns:
//	 --	an int64 between [0, n).
// 	panics:
//	 --	if n<= 0.
func (r *Random) Int63n(n int64) int64 {
	return r.rand.Int63n(n)
}

// Int31n -
//	input:
//	 n	-- upper limit.
// 	returns:
//	 --	an int32 between [0, n).
// 	panics:
//	 --	if n<= 0.
func (r *Random) Int31n(n int32) int32 {
	return r.rand.Int31n(n)
}

// Intn -
//	input:
//	 n	-- upper limit.
// 	returns:
//	 --	an int between [0, n).
func (r *Random) Intn(n int) int {
	return r.rand.Intn(n)
}

// Float64 -
// 	returns:
//	 --	an float64 between [0.0, 1.0).
func (r *Random) Float64() float64 {
	return r.rand.Float64()
}

// Float32 -
//	--	returns:
//	 --	an float32 between [0.0, 1.0).
func (r *Random) Float32() float32 {
	return r.rand.Float32()
}

// Perm -
// 	returns:
//	 --	as a slice of n ints in the half-open interval [0...n).
func (r *Random) Perm(n int) []int {
	return r.rand.Perm(n)
}

// Shuffle - Shuffle pseudo-randomizes the order of elements.
//	input:
//	 n	-- the number of elements.
//	 swap	-- swaps the elements with indexes i and j.
// 	panic: n < 0
func (r *Random) Shuffle(n int, swap func(i, j int)) {
	r.rand.Shuffle(n, swap)
}

// Read - generates len(p) random bytes and writes them into p.
//	returns:
//	 --	len(p) and a nil error.
func (r *Random) Read(p []byte) (n int, err error) {
	return r.rand.Read(p)
}
