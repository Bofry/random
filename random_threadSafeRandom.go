package random

import (
	"math/rand"
	"sync"
)

type threadSafeRandom struct {
	lk   sync.Mutex
	rand *rand.Rand
}

// NewthreadSafeRandom -
// 	returns:
//	 --	a new thread safe based on Random.
func NewthreadSafeRandom(src rand.Source) threadSafeRandom {
	return threadSafeRandom{
		rand: rand.New(src),
	}
}

// Seed -
// Seed uses the provided seed value to initialize the generator to a deterministic state.
// Seed should not be called concurrently with any other Rand method.
func (r *threadSafeRandom) Seed(seed int64) {
	r.rand.Seed(seed)
}

// Int63 -
// 	returns:
//	 --	an int64 between [-2^63, 2^63-1].
func (r *threadSafeRandom) Int63() int64 {
	r.lk.Lock()
	val := r.rand.Int63()
	r.lk.Unlock()
	return val
}

// Uint64 -
// 	returns:
//	 --	an Uint64 between [0, 2^64-1].
func (r *threadSafeRandom) Uint64() uint64 {
	r.lk.Lock()
	val := r.rand.Uint64()
	r.lk.Unlock()
	return val
}

// Uint32 -
// 	returns:
//	 --	an Uint32 between [0, 2^32-1].
func (r *threadSafeRandom) Uint32() uint32 {
	r.lk.Lock()
	val := r.rand.Uint32()
	r.lk.Unlock()
	return val
}

// Int31 -
// 	returns:
//	 --	an int64 between [-2^31, 2^31-1].
func (r *threadSafeRandom) Int31() int32 {
	r.lk.Lock()
	val := r.rand.Int31()
	r.lk.Unlock()
	return val
}

// Int -
// 	returns:
//	 --	an Int between [-2^31, 2^31-1].
func (r *threadSafeRandom) Int() int {
	r.lk.Lock()
	val := r.rand.Int()
	r.lk.Unlock()
	return val
}

// Int63n -
//	input:
//	 n	-- upper limit.
// 	returns:
//	 --	an int64 between [0, n).
// 	panics:
//	 --	if n<= 0.
func (r *threadSafeRandom) Int63n(n int64) int64 {
	r.lk.Lock()
	val := r.rand.Int63n(n)
	r.lk.Unlock()
	return val
}

// Int31n -
//	input:
//	 n	-- upper limit.
// 	returns:
//	 --	an int32 between [0, n).
// 	panics:
//	 --	if n<= 0.
func (r *threadSafeRandom) Int31n(n int32) int32 {
	r.lk.Lock()
	val := r.rand.Int31n(n)
	r.lk.Unlock()
	return val
}

// Intn -
//	input:
//	 n	-- upper limit.
// 	returns:
//	 --	an int between [0, n).
func (r *threadSafeRandom) Intn(n int) int {
	r.lk.Lock()
	val := r.rand.Intn(n)
	r.lk.Unlock()
	return val
}

// Float64 -
// 	returns:
//	 --	an float64 between [0.0, 1.0).
func (r *threadSafeRandom) Float64() float64 {
	r.lk.Lock()
	val := r.rand.Float64()
	r.lk.Unlock()
	return val
}

// Float32 -
//	--	returns:
//	 --	an float32 between [0.0, 1.0).
func (r *threadSafeRandom) Float32() float32 {
	r.lk.Lock()
	val := r.rand.Float32()
	r.lk.Unlock()
	return val
}

// Perm -
// 	returns:
//	 --	as a slice of n ints in the half-open interval [0...n).
func (r *threadSafeRandom) Perm(n int) []int {
	r.lk.Lock()
	val := r.rand.Perm(n)
	r.lk.Unlock()
	return val
}

// Shuffle - Shuffle pseudo-randomizes the order of elements.
//	input:
//	 n	-- the number of elements.
//	 swap	-- swaps the elements with indexes i and j.
// 	panic: n < 0
func (r *threadSafeRandom) Shuffle(n int, swap func(i, j int)) {
	r.lk.Lock()
	r.rand.Shuffle(n, swap)
	r.lk.Unlock()
}

// Read - generates len(p) random bytes and writes them into p.
//	returns:
//	 --	len(p) and a nil error.
func (r *threadSafeRandom) Read(p []byte) (n int, err error) {
	r.lk.Lock()
	n, err = r.rand.Read(p)
	r.lk.Unlock()
	return n, err
}

// Int63r - generates pseudo random int64 between low and high.
//  input:
//   low  -- lower limit.
//   high -- upper limit.
//  returns:
//	 -- an int64 between [low, high].
func (r *threadSafeRandom) Int63r(low, high int64) int64 {
	r.lk.Lock()
	val := r.Int63()%(high-low+1) + low
	r.lk.Unlock()
	return val
}

// Int63s -
// generates pseudo Randomom integers between low and high.
//  input:
//   low    -- lower limit.
//   high   -- upper limit.
//  output:
//   values -- slice to be filled with len(values) numbers.
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

// Int63Shuffle -
// shuffles a slice of integers.
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

// Uint32 -
// generates pseudo Randomom uint32 between low and high.
//  input:
//   low  -- lower limit.
//   high -- upper limit.
//  returns:
//	 -- an uint32 between [low, high].
func (r *threadSafeRandom) Uint32r(low, high uint32) uint32 {
	r.lk.Lock()
	val := r.Uint32()%(high-low+1) + low
	r.lk.Unlock()
	return val
}

// Uint32s -
// generates pseudo Randomom integers between low and high.
//  input:
//   low    -- lower limit.
//   high   -- upper limit.
//  output:
//   values -- slice to be filled with len(values) numbers.
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

// Uint32Shuffle -
// shuffles a slice of integers.
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

// Uint64r -
// generates pseudo Randomom uint64 between low and high.
//  input:
//   low  -- lower limit.
//   high -- upper limit.
//  returns:
//	 -- an uint64 between [low, high].
func (r *threadSafeRandom) Uint64r(low, high uint64) uint64 {
	r.lk.Lock()
	val := r.Uint64()%(high-low+1) + low
	r.lk.Unlock()
	return val
}

// Uint64s -
// generates pseudo Randomom integers between low and high.
//  input:
//   low    -- lower limit.
//   high   -- upper limit.
//  output:
//   values -- slice to be filled with len(values) numbers.
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

// Uint64Shuffle -
// shuffles a slice of integers.
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

// Int31r -
// is int range generates pseudo Randomom int32 between low and high.
//  input:
//   low  -- lower limit.
//   high -- upper limit.
//  returns:
//	 -- an int32 between [low, high].
func (r *threadSafeRandom) Int31r(low, high int32) int32 {
	r.lk.Lock()
	val := r.Int31()%(high-low+1) + low
	r.lk.Unlock()
	return val
}

// Int31s -
// generates pseudo Randomom integers between low and high.
//  input:
//   low    -- lower limit.
//   high   -- upper limit.
//  output:
//   values -- slice to be filled with len(values) numbers.
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

// Int31Shuffle -
// shuffles a slice of integers.
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

// Intr -
// is int range generates pseudo Randomom integer between low and high.
//  input:
//   low  -- lower limit.
//   high -- upper limit.
//  returns:
//	 -- an int between [low, high].
func (r *threadSafeRandom) Intr(low, high int) int {
	r.lk.Lock()
	val := r.Int()%(high-low+1) + low
	r.lk.Unlock()
	return val
}

// Ints -
// generates pseudo Randomom integers between low and high.
//  input:
//   low    -- lower limit.
//   high   -- upper limit.
//  output:
//   values -- slice to be filled with len(values) numbers.
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

// IntShuffle -
// shuffles a slice of integers.
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

// Float64r -
// generates a pseudo Randomom real number between low and high; i.e. in [low, right)
//  input:
//   low  -- lower limit. (closed)
//   high -- upper limit. (open)
//  returns:
//	 -- an int between [low, high].
func (r *threadSafeRandom) Float64r(low, high float64) float64 {
	r.lk.Lock()
	val := low + (high-low)*r.Float64()
	r.lk.Unlock()
	return val
}

// Float64s -
// generates pseudo Randomom real numbers between low and high; i.e. in [low, right)
//  input:
//   low  -- lower limit. (closed)
//   high -- upper limit. (open)
//  output:
//   values -- slice to be filled with len(values) numbers
func (r *threadSafeRandom) Float64s(values []float64, low, high float64) {
	r.lk.Lock()
	for i := 0; i < len(values); i++ {
		values[i] = low + (high-low)*r.Float64()
	}
	r.lk.Unlock()
}

// Float64Shuffle -
// shuffles a slice of float point numbers
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

// Float32r generates a pseudo Randomom real number between low and high; i.e. in [low, right)
//  Input:
//   low  -- lower limit. (closed)
//   high -- upper limit. (open)
//  returns:
//	 -- an int between [low, high].
func (r *threadSafeRandom) Float32r(low, high float32) float32 {
	r.lk.Lock()
	val := low + (high-low)*r.Float32()
	r.lk.Unlock()
	return val
}

// Float32s -
// generates pseudo Randomom real numbers between low and high; i.e. in [low, right)
//  input:
//   low  -- lower limit. (closed)
//   high -- upper limit. (open)
//  output:
//   values -- slice to be filled with len(values) numbers.
func (r *threadSafeRandom) Float32s(values []float32, low, high float32) {
	r.lk.Lock()
	for i := 0; i < len(values); i++ {
		values[i] = low + (high-low)*r.Float32()
	}
	r.lk.Unlock()
}

// Float32Shuffle -
// shuffles a slice of float point numbers.
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

// FlipCoin -
// generates a Bernoulli variable; throw a coin with probability p.
func (r *threadSafeRandom) FlipCoin(p float64) bool {
	r.lk.Lock()
	if p == 1.0 {
		return true
	}
	if p == 0.0 {
		return false
	}
	if r.Float64() <= p {
		return true
	}

	r.lk.Unlock()
	return false
}

// Uint32n -
//  input:
//   n	-- upper limit.
//  returns:
//	 -- an uint32 between [0, n).
// 	panics:
//	 --	if n<= 0.
func (r *threadSafeRandom) Uint32n(n uint32) uint32 {
	if n <= 0 {
		panic("invalid argument to Uint32n")
	}
	return r.Uint32() % n
}

// Float64n -
//  input:
//   n	-- upper limit.
//  returns:
//	 -- an float64 between [0, n).
func (r *threadSafeRandom) Float64n(n float64) float64 {
	return n * r.Float64()
}

// Float32n -
//  input:
//   n	-- upper limit.
//  returns:
//	 -- an float32 between [0, n).
// 	panics:
//	 --	if n<= 0.
func (r *threadSafeRandom) Float32n(n float32) float32 {
	return (n) * r.Float32()
}

// Float64w -
// Ｗhich randomly picks an index in the range [0, w.length - 1] (inclusive) and returns it.
// The probability of picking an index i is w[i] / sum(w).
//  input:
//   w	-- an array of positive float64 w.
//  returns:
//	 -- picks an index in the range [0, len(w)-1]..
// 	panics:
//	 --	if n<= 0.
//   -- if w[i] is not positive.
func (r *threadSafeRandom) Float64w(w []float64) int {
	l := len(w)
	if l == 0 {
		panic("invalid argument to Float64w")
	}
	if l == 1 {
		return 0
	}
	var totalWeigehts float64
	for _, v := range w {
		if v < 0 {
			panic("Float64w: w[i] is not positive.")
		}
		totalWeigehts += v
	}

	dice := r.Float64n(totalWeigehts)
	for i := 0; i < l-1; i++ {
		if dice < w[i] {
			return i
		}
		dice -= w[i]
	}
	return l - 1
}

// Float32w -
// Ｗhich randomly picks an index in the range [0, w.length - 1] (inclusive) and returns it.
// The probability of picking an index i is w[i] / sum(w).
//  input:
//   w	-- an array of positive float32 w.
//  returns:
//	 -- picks an index in the range [0, len(w)-1]..
// 	panics:
//	 --	if n<= 0.
//   -- if w[i] is not positive.
func (r *threadSafeRandom) Float32w(w []float32) int {
	l := len(w)
	if l == 0 {
		panic("invalid argument to Float32w")
	}
	if l == 1 {
		return 0
	}
	var totalWeigehts float32
	for _, v := range w {
		if v < 0 {
			panic("Float32w: w[i] is not positive.")
		}
		totalWeigehts += v
	}

	dice := r.Float32n(totalWeigehts)
	for i := 0; i < l-1; i++ {
		if dice < w[i] {
			return i
		}
		dice -= w[i]
	}
	return l - 1
}

// Uint64w -
// Ｗhich randomly picks an index in the range [0, w.length - 1] (inclusive) and returns it.
// The probability of picking an index i is w[i] / sum(w).
//  input:
//   w	-- an array of positive uint64 w.
//  returns:
//	 -- picks an index in the range [0, len(w)-1]..
// 	panics:
//	 --	if n<= 0.
func (r *threadSafeRandom) Uint64w(w []uint64) int {
	l := len(w)
	if l == 0 {
		panic("invalid argument to Uint64w")
	}
	if l == 1 {
		return 0
	}
	var totalWeigehts uint64
	for _, v := range w {
		totalWeigehts += v
	}

	dice := r.Uint64n(totalWeigehts)
	for i := 0; i < l-1; i++ {
		if dice < w[i] {
			return i
		}
		dice -= w[i]
	}
	return l - 1
}

// Uint32w -
// Ｗhich randomly picks an index in the range [0, w.length - 1] (inclusive) and returns it.
// The probability of picking an index i is w[i] / sum(w).
//  input:
//   w	-- an array of positive uint32 w.
//  returns:
//	 -- picks an index in the range [0, len(w)-1]..
// 	panics:
//	 --	if n<= 0.
func (r *threadSafeRandom) Uint32w(w []uint32) int {
	l := len(w)
	if l == 0 {
		panic("invalid argument to Uint32w")
	}
	if l == 1 {
		return 0
	}
	var totalWeigehts uint32
	for _, v := range w {
		totalWeigehts += v
	}

	dice := r.Uint32n(totalWeigehts)
	for i := 0; i < l-1; i++ {
		if dice < w[i] {
			return i
		}
		dice -= w[i]
	}
	return l - 1
}

// Int64w -
// Ｗhich randomly picks an index in the range [0, w.length - 1] (inclusive) and returns it.
// The probability of picking an index i is w[i] / sum(w).
//  input:
//   w	-- an array of positive int64 w.
//  returns:
//	 -- picks an index in the range [0, len(w)-1]..
// 	panics:
//	 --	if n<= 0.
//   -- if w[i] is not positive.
func (r *threadSafeRandom) Int64w(w []int64) int {
	l := len(w)
	if l == 0 {
		panic("invalid argument to Int64w")
	}
	if l == 1 {
		return 0
	}
	var totalWeigehts int64
	for _, v := range w {
		if v < 0 {
			panic("Int64w: w[i] is not positive.")
		}
		totalWeigehts += v
	}

	dice := r.Int63n(totalWeigehts)
	for i := 0; i < l-1; i++ {
		if dice < w[i] {
			return i
		}
		dice -= w[i]
	}
	return l - 1
}

// Int32w -
// Ｗhich randomly picks an index in the range [0, w.length - 1] (inclusive) and returns it.
// The probability of picking an index i is w[i] / sum(w).
//  input:
//   w	-- an array of positive int32 w.
//  returns:
//	 -- picks an index in the range [0, len(w)-1]..
// 	panics:
//	 --	if n<= 0.
//   -- if w[i] is not positive.
func (r *threadSafeRandom) Int32w(w []int32) int {
	l := len(w)
	if l == 0 {
		panic("invalid argument to Int32w")
	}
	if l == 1 {
		return 0
	}
	var totalWeigehts int32
	for _, v := range w {
		if v < 0 {
			panic("Int32w: w[i] is not positive.")
		}
		totalWeigehts += v
	}

	dice := r.Int31n(totalWeigehts)
	for i := 0; i < l-1; i++ {
		if dice < w[i] {
			return i
		}
		dice -= w[i]
	}
	return l - 1
}

// Intw -
// Ｗhich randomly picks an index in the range [0, w.length - 1] (inclusive) and returns it.
// The probability of picking an index i is w[i] / sum(w).
//  input:
//   w	-- an array of positive int w.
//  returns:
//	 -- picks an index in the range [0, len(w)-1]..
// 	panics:
//	 --	if n<= 0.
//   -- if w[i] is not positive.
func (r *threadSafeRandom) Intw(w []int) int {
	l := len(w)
	if l == 0 {
		panic("invalid argument to Intw")
	}
	if l == 1 {
		return 0
	}
	var totalWeigehts int
	for _, v := range w {
		if v < 0 {
			panic("Intw: w[i] is not positive.")
		}
		totalWeigehts += v
	}

	dice := r.Intn(totalWeigehts)
	for i := 0; i < l-1; i++ {
		if dice < w[i] {
			return i
		}
		dice -= w[i]
	}
	return l - 1
}

// Uint64n -
//  input:
//   n	-- upper limit.
//  returns:
//	 -- an uint64 between [0, n).
// 	panics:
//	 --	if n<= 0.
func (r *threadSafeRandom) Uint64n(n uint64) uint64 {
	if n <= 0 {
		panic("invalid argument to Uint64n")
	}
	return r.Uint64() % n
}
