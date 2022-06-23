package random

// Uint64n -
//  input:
//   n	-- upper limit.
//  returns:
//	 -- an uint64 between [0, n).
// 	panics:
//	 --	if n<= 0.
func (r *Random) Uint64n(n uint64) uint64 {
	if n <= 0 {
		panic("invalid argument to Uint64n")
	}
	return r.Uint64() % n
}

// Uint32n -
//  input:
//   n	-- upper limit.
//  returns:
//	 -- an uint32 between [0, n).
// 	panics:
//	 --	if n<= 0.
func (r *Random) Uint32n(n uint32) uint32 {
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
func (r *Random) Float64n(n float64) float64 {
	return n * r.Float64()
}

// Float32n -
//  input:
//   n	-- upper limit.
//  returns:
//	 -- an float32 between [0, n).
// 	panics:
//	 --	if n<= 0.
func (r *Random) Float32n(n float32) float32 {
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
func (r *Random) Float64w(w []float64) int {
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
func (r *Random) Float32w(w []float32) int {
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
func (r *Random) Uint64w(w []uint64) int {
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
func (r *Random) Uint32w(w []uint32) int {
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
func (r *Random) Int64w(w []int64) int {
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
func (r *Random) Int32w(w []int32) int {
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
func (r *Random) Intw(w []int) int {
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
