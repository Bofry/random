// It was originally implemented in golang by gosl.
// Due to lightweight requirements, we only copy the rnd function from gosl.

// Copyright (c) 2016, Dorival Pedroso.
// <https://github.com/cpmech/gosl>

// Copyright (c) 2016, Dorival Pedroso. All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:

// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.

// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.

// * Neither the name of Gosl nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package random

// Int63r - generates pseudo random int64 between low and high.
//  input:
//   low  -- lower limit.
//   high -- upper limit.
//  returns:
//	 -- an int64 between [low, high].
func (r *Random) Int63r(low, high int64) int64 {
	return r.Int63()%(high-low+1) + low
}

// Int63s -
// generates pseudo Randomom integers between low and high.
//  input:
//   low    -- lower limit.
//   high   -- upper limit.
//  output:
//   values -- slice to be filled with len(values) numbers.
func (r *Random) Int63s(values []int64, low, high int64) {
	if len(values) < 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		values[i] = r.Int63r(low, high)
	}
}

// Int63Shuffle -
// shuffles a slice of integers.
func (r *Random) Int63Shuffle(values []int64) {
	var tmp int64
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
}

// Uint32 -
// generates pseudo Randomom uint32 between low and high.
//  input:
//   low  -- lower limit.
//   high -- upper limit.
//  returns:
//	 -- an uint32 between [low, high].
func (r *Random) Uint32r(low, high uint32) uint32 {
	return r.Uint32()%(high-low+1) + low
}

// Uint32s -
// generates pseudo Randomom integers between low and high.
//  input:
//   low    -- lower limit.
//   high   -- upper limit.
//  output:
//   values -- slice to be filled with len(values) numbers.
func (r *Random) Uint32s(values []uint32, low, high uint32) {
	if len(values) < 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		values[i] = r.Uint32r(low, high)
	}
}

// Uint32Shuffle -
// shuffles a slice of integers.
func (r *Random) Uint32Shuffle(values []uint32) {
	var tmp uint32
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
}

// Uint64r -
// generates pseudo Randomom uint64 between low and high.
//  input:
//   low  -- lower limit.
//   high -- upper limit.
//  returns:
//	 -- an uint64 between [low, high].
func (r *Random) Uint64r(low, high uint64) uint64 {
	return r.Uint64()%(high-low+1) + low
}

// Uint64s -
// generates pseudo Randomom integers between low and high.
//  input:
//   low    -- lower limit.
//   high   -- upper limit.
//  output:
//   values -- slice to be filled with len(values) numbers.
func (r *Random) Uint64s(values []uint64, low, high uint64) {
	if len(values) < 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		values[i] = r.Uint64r(low, high)
	}
}

// Uint64Shuffle -
// shuffles a slice of integers.
func (r *Random) Uint64Shuffle(values []uint64) {
	var tmp uint64
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
}

// Int31r -
// is int range generates pseudo Randomom int32 between low and high.
//  input:
//   low  -- lower limit.
//   high -- upper limit.
//  returns:
//	 -- an int32 between [low, high].
func (r *Random) Int31r(low, high int32) int32 {
	return r.Int31()%(high-low+1) + low
}

// Int31s -
// generates pseudo Randomom integers between low and high.
//  input:
//   low    -- lower limit.
//   high   -- upper limit.
//  output:
//   values -- slice to be filled with len(values) numbers.
func (r *Random) Int31s(values []int32, low, high int32) {
	if len(values) < 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		values[i] = r.Int31r(low, high)
	}
}

// Int31Shuffle -
// shuffles a slice of integers.
func (r *Random) Int31Shuffle(values []int32) {
	var tmp int32
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
}

// Intr -
// is int range generates pseudo Randomom integer between low and high.
//  input:
//   low  -- lower limit.
//   high -- upper limit.
//  returns:
//	 -- an int between [low, high].
func (r *Random) Intr(low, high int) int {
	return r.Int()%(high-low+1) + low
}

// Ints -
// generates pseudo Randomom integers between low and high.
//  input:
//   low    -- lower limit.
//   high   -- upper limit.
//  output:
//   values -- slice to be filled with len(values) numbers.
func (r *Random) Ints(values []int, low, high int) {
	if len(values) < 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		values[i] = r.Intr(low, high)
	}
}

// IntShuffle -
// shuffles a slice of integers.
func (r *Random) IntShuffle(values []int) {
	var j, tmp int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
}

// Float64r -
// generates a pseudo Randomom real number between low and high; i.e. in [low, right)
//  input:
//   low  -- lower limit. (closed)
//   high -- upper limit. (open)
//  returns:
//	 -- an int between [low, high].
func (r *Random) Float64r(low, high float64) float64 {
	return low + (high-low)*r.Float64()
}

// Float64s -
// generates pseudo Randomom real numbers between low and high; i.e. in [low, right)
//  input:
//   low  -- lower limit. (closed)
//   high -- upper limit. (open)
//  output:
//   values -- slice to be filled with len(values) numbers
func (r *Random) Float64s(values []float64, low, high float64) {
	for i := 0; i < len(values); i++ {
		values[i] = low + (high-low)*r.Float64()
	}
}

// Float64Shuffle -
// shuffles a slice of float point numbers
func (r *Random) Float64Shuffle(values []float64) {
	var tmp float64
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
}

// Float32r generates a pseudo Randomom real number between low and high; i.e. in [low, right)
//  Input:
//   low  -- lower limit. (closed)
//   high -- upper limit. (open)
//  returns:
//	 -- an int between [low, high].
func (r *Random) Float32r(low, high float32) float32 {
	return low + (high-low)*r.Float32()
}

// Float32s -
// generates pseudo Randomom real numbers between low and high; i.e. in [low, right)
//  input:
//   low  -- lower limit. (closed)
//   high -- upper limit. (open)
//  output:
//   values -- slice to be filled with len(values) numbers.
func (r *Random) Float32s(values []float32, low, high float32) {
	for i := 0; i < len(values); i++ {
		values[i] = low + (high-low)*r.Float32()
	}
}

// Float32Shuffle -
// shuffles a slice of float point numbers.
func (r *Random) Float32Shuffle(values []float32) {
	var tmp float32
	var j int
	for i := len(values) - 1; i > 0; i-- {
		j = r.Int() % i
		tmp = values[j]
		values[j] = values[i]
		values[i] = tmp
	}
}

// FlipCoin -
// generates a Bernoulli variable; throw a coin with probability p.
func (r *Random) FlipCoin(p float64) bool {
	if p == 1.0 {
		return true
	}
	if p == 0.0 {
		return false
	}
	if r.Float64() <= p {
		return true
	}
	return false
}
