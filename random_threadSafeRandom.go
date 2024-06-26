package random

import (
	"math/rand"
	"sync"
)

var _ rand.Source64 = (*threadSafeRandom)(nil) // Ensures Rand complies with rand.Source64

// threadSafeRandom provides a random number generator safe for concurrent use.
type threadSafeRandom struct {
	lk   sync.RWMutex
	rand *rand.Rand
}

// NewthreadSafeRandom returns a new threadSafeRandom.
func NewthreadSafeRandom(src rand.Source) threadSafeRandom {
	return threadSafeRandom{
		rand: rand.New(src),
	}
}

// Seed initializes the generator to a deterministic state.
func (r *threadSafeRandom) Seed(seed int64) {
	r.lk.Lock()
	r.rand.Seed(seed)
	r.lk.Unlock()
}

// Int63 returns a non-negative pseudo-random int64.
func (r *threadSafeRandom) Int63() int64 {
	r.lk.RLock()
	val := r.rand.Int63()
	r.lk.RUnlock()
	return val
}

// Uint64 returns a non-negative pseudo-random uint64.
func (r *threadSafeRandom) Uint64() uint64 {
	r.lk.RLock()
	val := r.rand.Uint64()
	r.lk.RUnlock()
	return val
}

// Uint32 returns a non-negative pseudo-random uint32.
func (r *threadSafeRandom) Uint32() uint32 {
	r.lk.RLock()
	val := r.rand.Uint32()
	r.lk.RUnlock()
	return val
}

// Int31 returns a non-negative pseudo-random int32.
func (r *threadSafeRandom) Int31() int32 {
	r.lk.RLock()
	val := r.rand.Int31()
	r.lk.RUnlock()
	return val
}

// Int returns a non-negative pseudo-random int.
func (r *threadSafeRandom) Int() int {
	r.lk.RLock()
	val := r.rand.Int()
	r.lk.RUnlock()
	return val
}

// Int63n returns a non-negative pseudo-random int64 in [0, n).
// Panics if n <= 0.
func (r *threadSafeRandom) Int63n(n int64) int64 {
	r.lk.RLock()
	val := r.rand.Int63n(n)
	r.lk.RUnlock()
	return val
}

// Int31n returns a non-negative pseudo-random int32 in [0, n).
// Panics if n <= 0.
func (r *threadSafeRandom) Int31n(n int32) int32 {
	r.lk.RLock()
	val := r.rand.Int31n(n)
	r.lk.RUnlock()
	return val
}

// Intn returns a non-negative pseudo-random int in [0, n).
func (r *threadSafeRandom) Intn(n int) int {
	r.lk.RLock()
	val := r.rand.Intn(n)
	r.lk.RUnlock()
	return val
}

// Float64 returns a pseudo-random float64 in [0.0, 1.0).
func (r *threadSafeRandom) Float64() float64 {
	r.lk.RLock()
	val := r.rand.Float64()
	r.lk.RUnlock()
	return val
}

// Float32 returns a pseudo-random float32 in [0.0, 1.0).
func (r *threadSafeRandom) Float32() float32 {
	r.lk.RLock()
	val := r.rand.Float32()
	r.lk.RUnlock()
	return val
}

// Perm returns a slice of n ints in [0, n).
func (r *threadSafeRandom) Perm(n int) []int {
	r.lk.Lock()
	val := r.rand.Perm(n)
	r.lk.Unlock()
	return val
}

// Shuffle pseudo-randomizes the order of elements.
func (r *threadSafeRandom) Shuffle(n int, swap func(i, j int)) {
	r.lk.Lock()
	r.rand.Shuffle(n, swap)
	r.lk.Unlock()
}

// Read generates len(p) random bytes and writes them into p.
func (r *threadSafeRandom) Read(p []byte) (n int, err error) {
	r.lk.Lock()
	n, err = r.rand.Read(p)
	r.lk.Unlock()
	return n, err
}

// Int63r returns a pseudo-random int64 in [low, high].
func (r *threadSafeRandom) Int63r(low, high int64) int64 {
	r.lk.RLock()
	val := r.Int63()%(high-low+1) + low
	r.lk.RUnlock()
	return val
}

// Int63s fills the values slice with pseudo-random int64 in [low, high].
func (r *threadSafeRandom) Int63s(values []int64, low, high int64) {
	r.lk.Lock()
	if len(values) < 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		values[i] = r.Int63r(low, high)
	}
	r.lk.Unlock()
}

// Int63Shuffle shuffles a slice of int64.
func (r *threadSafeRandom) Int63Shuffle(values []int64) {
	r.lk.Lock()
	var tmp int64
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
	r.lk.Unlock()
}

// Uint32r returns a pseudo-random uint32 in [low, high].
func (r *threadSafeRandom) Uint32r(low, high uint32) uint32 {
	r.lk.RLock()
	val := r.Uint32()%(high-low+1) + low
	r.lk.RUnlock()
	return val
}

// Uint32s fills the values slice with pseudo-random uint32 in [low, high].
func (r *threadSafeRandom) Uint32s(values []uint32, low, high uint32) {
	r.lk.Lock()
	if len(values) < 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		values[i] = r.Uint32r(low, high)
	}
	r.lk.Unlock()
}

// Uint32Shuffle shuffles a slice of uint32.
func (r *threadSafeRandom) Uint32Shuffle(values []uint32) {
	r.lk.Lock()
	var tmp uint32
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
	r.lk.Unlock()
}

// Uint64r returns a pseudo-random uint64 in [low, high].
func (r *threadSafeRandom) Uint64r(low, high uint64) uint64 {
	r.lk.RLock()
	val := r.Uint64()%(high-low+1) + low
	r.lk.RUnlock()
	return val
}

// Uint64s fills the values slice with pseudo-random uint64 in [low, high].
func (r *threadSafeRandom) Uint64s(values []uint64, low, high uint64) {
	r.lk.Lock()
	if len(values) < 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		values[i] = r.Uint64r(low, high)
	}
	r.lk.Unlock()
}

// Uint64Shuffle shuffles a slice of uint64.
func (r *threadSafeRandom) Uint64Shuffle(values []uint64) {
	r.lk.Lock()
	var tmp uint64
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
	r.lk.Unlock()
}

// Int31r returns a pseudo-random int32 in [low, high].
func (r *threadSafeRandom) Int31r(low, high int32) int32 {
	r.lk.RLock()
	val := r.Int31()%(high-low+1) + low
	r.lk.RUnlock()
	return val
}

// Int31s fills the values slice with pseudo-random int32 in [low, high].
func (r *threadSafeRandom) Int31s(values []int32, low, high int32) {
	r.lk.Lock()
	if len(values) < 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		values[i] = r.Int31r(low, high)
	}
	r.lk.Unlock()
}

// Int31Shuffle shuffles a slice of int32.
func (r *threadSafeRandom) Int31Shuffle(values []int32) {
	r.lk.Lock()
	var tmp int32
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
	r.lk.Unlock()
}

// Intr returns a pseudo-random int in [low, high].
func (r *threadSafeRandom) Intr(low, high int) int {
	r.lk.RLock()
	val := r.Int()%(high-low+1) + low
	r.lk.RUnlock()
	return val
}

// Ints fills the values slice with pseudo-random int in [low, high].
func (r *threadSafeRandom) Ints(values []int, low, high int) {
	r.lk.Lock()
	if len(values) < 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		values[i] = r.Intr(low, high)
	}
	r.lk.Unlock()
}

// IntShuffle shuffles a slice of int.
func (r *threadSafeRandom) IntShuffle(values []int) {
	r.lk.Lock()
	var j, tmp int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
	r.lk.Unlock()
}

// Float64r returns a pseudo-random float64 in [low, high).
func (r *threadSafeRandom) Float64r(low, high float64) float64 {
	r.lk.RLock()
	val := low + (high-low)*r.Float64()
	r.lk.RUnlock()
	return val
}

// Float64s fills the values slice with pseudo-random float64 in [low, high).
func (r *threadSafeRandom) Float64s(values []float64, low, high float64) {
	r.lk.Lock()
	for i := 0; i < len(values); i++ {
		values[i] = low + (high-low)*r.Float64()
	}
	r.lk.Unlock()
}

// Float64Shuffle shuffles a slice of float64.
func (r *threadSafeRandom) Float64Shuffle(values []float64) {
	r.lk.Lock()
	var tmp float64
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
	r.lk.Unlock()
}

// Float32r returns a pseudo-random float32 in [low, high).
func (r *threadSafeRandom) Float32r(low, high float32) float32 {
	r.lk.RLock()
	val := low + (high-low)*r.Float32()
	r.lk.RUnlock()
	return val
}

// Float32s fills the values slice with pseudo-random float32 in [low, high).
func (r *threadSafeRandom) Float32s(values []float32, low, high float32) {
	r.lk.Lock()
	for i := 0; i < len(values); i++ {
		values[i] = low + (high-low)*r.Float32()
	}
	r.lk.Unlock()
}

// Float32Shuffle shuffles a slice of float32.
func (r *threadSafeRandom) Float32Shuffle(values []float32) {
	r.lk.Lock()
	var tmp float32
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
	r.lk.Unlock()
}

// FlipCoin returns true with probability p.
func (r *threadSafeRandom) FlipCoin(p float64) bool {
	r.lk.RLock()
	defer r.lk.RUnlock()

	if p == 1.0 {
		return true
	}
	if p == 0.0 {
		return false
	}
	return r.Float64() <= p
}

// Uint64n returns a non-negative pseudo-random uint64 in [0, n).
// Panics if n <= 0.
func (r *threadSafeRandom) Uint64n(n uint64) uint64 {
	if n <= 0 {
		panic("invalid argument to Uint64n")
	}
	r.lk.RLock()
	defer r.lk.RUnlock()
	return r.rand.Uint64() % n
}

// Uint32n returns a non-negative pseudo-random uint32 in [0, n).
// Panics if n <= 0.
func (r *threadSafeRandom) Uint32n(n uint32) uint32 {
	if n <= 0 {
		panic("invalid argument to Uint32n")
	}
	return r.Uint32() % n
}

// Float64n returns a pseudo-random float64 in [0.0, n).
func (r *threadSafeRandom) Float64n(n float64) float64 {
	return n * r.Float64()
}

// Float32n returns a pseudo-random float32 in [0.0, n).
// Panics if n <= 0.
func (r *threadSafeRandom) Float32n(n float32) float32 {
	if n <= 0 {
		panic("invalid argument to Float32n")
	}
	return n * r.Float32()
}

// ----------------------------------------------------------------------------
// Weighted Random Selection
// ----------------------------------------------------------------------------

// Float64w randomly picks an index in [0, len(w)-1] based on the weights in slice w.
// The probability of picking index i is w[i] / sum(w).
// Panics if w is empty or contains non-positive values.
func (r *threadSafeRandom) Float64w(w []float64) int {
	return weightedRandomIndexFloat64(r.rand, w)
}

// Float32w randomly picks an index in [0, len(w)-1] based on the weights in slice w.
// The probability of picking index i is w[i] / sum(w).
// Panics if w is empty or contains non-positive values.
func (r *threadSafeRandom) Float32w(w []float32) int {
	return weightedRandomIndexFloat32(r.rand, w)
}

// Uint64w randomly picks an index in [0, len(w)-1] based on the weights in slice w.
// The probability of picking index i is w[i] / sum(w).
// Panics if w is empty.
func (r *threadSafeRandom) Uint64w(w []uint64) int {
	return weightedRandomIndexUint64(r.rand, w)
}

// Uint32w randomly picks an index in [0, len(w)-1] based on the weights in slice w.
// The probability of picking index i is w[i] / sum(w).
// Panics if w is empty.
func (r *threadSafeRandom) Uint32w(w []uint32) int {
	return weightedRandomIndexUint32(r.rand, w)
}

// Int64w randomly picks an index in [0, len(w)-1] based on the weights in slice w.
// The probability of picking index i is w[i] / sum(w).
// Panics if w is empty or contains non-positive values.
func (r *threadSafeRandom) Int64w(w []int64) int {
	return weightedRandomIndexInt64(r.rand, w)
}

// Int32w randomly picks an index in [0, len(w)-1] based on the weights in slice w.
// The probability of picking index i is w[i] / sum(w).
// Panics if w is empty or contains non-positive values.
func (r *threadSafeRandom) Int32w(w []int32) int {
	return weightedRandomIndexInt32(r.rand, w)
}

// Intw randomly picks an index in [0, len(w)-1] based on the weights in slice w.
// The probability of picking index i is w[i] / sum(w).
// Panics if w is empty or contains non-positive values.
func (r *threadSafeRandom) Intw(w []int) int {
	return weightedRandomIndexInt(r.rand, w)
}
