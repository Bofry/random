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

import "math/big"

// Int63r generates a pseudo-random int64 between low (inclusive) and high (inclusive).
func (r *Random) Int63r(low, high int64) int64 {
	if low > high {
		low, high = high, low // Swap if low is greater than high
	}
	rangeBig := big.NewInt(0).Sub(big.NewInt(high), big.NewInt(low))
	rangeBig.Add(rangeBig, big.NewInt(1)) // Add 1 to include high
	offsetBig := big.NewInt(0).Mod(big.NewInt(r.Int63()), rangeBig)
	resultBig := big.NewInt(0).Add(offsetBig, big.NewInt(low))
	return resultBig.Int64()
}

// Int63s generates a slice of pseudo-random int64 values between low (inclusive) and high (inclusive).
func (r *Random) Int63s(values []int64, low, high int64) {
	for i := range values {
		values[i] = r.Int63r(low, high)
	}
}

// Int63Shuffle shuffles a slice of int64 values.
func (r *Random) Int63Shuffle(values []int64) {
	r.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })
}

// Uint32r generates a pseudo-random uint32 between low (inclusive) and high (inclusive).
func (r *Random) Uint32r(low, high uint32) uint32 {
	if low > high {
		low, high = high, low
	}
	return low + r.Uint32()%(high-low+1) // No overflow risk with uint32
}

// Uint32s generates a slice of pseudo-random uint32 values between low (inclusive) and high (inclusive).
func (r *Random) Uint32s(values []uint32, low, high uint32) {
	for i := range values {
		values[i] = r.Uint32r(low, high)
	}
}

// Uint32Shuffle shuffles a slice of uint32 values.
func (r *Random) Uint32Shuffle(values []uint32) {
	r.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })
}

// Uint64r generates a pseudo-random uint64 between low (inclusive) and high (inclusive).
func (r *Random) Uint64r(low, high uint64) uint64 {
	if low > high {
		low, high = high, low
	}
	rangeBig := big.NewInt(0).Sub(big.NewInt(int64(high)), big.NewInt(int64(low)))
	rangeBig.Add(rangeBig, big.NewInt(1))
	offsetBig := big.NewInt(0).Mod(big.NewInt(int64(r.Uint64())), rangeBig)
	resultBig := big.NewInt(0).Add(offsetBig, big.NewInt(int64(low)))
	return resultBig.Uint64()
}

// Uint64s generates a slice of pseudo-random uint64 values between low (inclusive) and high (inclusive).
func (r *Random) Uint64s(values []uint64, low, high uint64) {
	for i := range values {
		values[i] = r.Uint64r(low, high)
	}
}

// Uint64Shuffle shuffles a slice of uint64 values.
func (r *Random) Uint64Shuffle(values []uint64) {
	r.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })
}

// Int31r generates a pseudo-random int32 between low (inclusive) and high (inclusive).
func (r *Random) Int31r(low, high int32) int32 {
	if low > high {
		low, high = high, low
	}
	return low + r.Int31n(high-low+1) // No overflow risk with int31n
}

// Int31s generates a slice of pseudo-random int32 values between low (inclusive) and high (inclusive).
func (r *Random) Int31s(values []int32, low, high int32) {
	if len(values) < 1 {
		return
	}
	for i := 0; i < len(values); i++ {
		values[i] = r.Int31r(low, high)
	}
}

// Int31Shuffle shuffles a slice of int32 values.
func (r *Random) Int31Shuffle(values []int32) {
	r.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })
}

// Intr generates a pseudo-random int between low (inclusive) and high (inclusive).
func (r *Random) Intr(low, high int) int {
	if low > high {
		low, high = high, low
	}
	return low + r.Intn(high-low+1) // No overflow risk with Intn
}

// Ints generates a slice of pseudo-random int values between low (inclusive) and high (inclusive).
func (r *Random) Ints(values []int, low, high int) {
	for i := range values {
		values[i] = r.Intr(low, high)
	}
}

// IntShuffle shuffles a slice of int values.
func (r *Random) IntShuffle(values []int) {
	r.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })
}

// Float64r generates a pseudo-random float64 in the range [low, high).
func (r *Random) Float64r(low, high float64) float64 {
	if low > high {
		low, high = high, low
	}
	return low + (high-low)*r.Float64()
}

// Float64s fills a slice with pseudo-random float64 values in the range [low, high).
func (r *Random) Float64s(values []float64, low, high float64) {
	for i := range values {
		values[i] = r.Float64r(low, high)
	}
}

// Float64Shuffle shuffles a slice of float64 values.
func (r *Random) Float64Shuffle(values []float64) {
	r.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })
}

// Float32r generates a pseudo-random float32 in the range [low, high).
func (r *Random) Float32r(low, high float32) float32 {
	if low > high {
		low, high = high, low
	}
	return low + (high-low)*r.Float32()
}

// Float32s fills a slice with pseudo-random float32 values in the range [low, high).
func (r *Random) Float32s(values []float32, low, high float32) {
	for i := range values {
		values[i] = r.Float32r(low, high)
	}
}

// Float32Shuffle shuffles a slice of float32 values.
func (r *Random) Float32Shuffle(values []float32) {
	r.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })
}

// FlipCoin simulates a coin flip with the given probability p of heads (true).
func (r *Random) FlipCoin(p float64) bool {
	if p == 1.0 {
		return true
	}
	if p == 0.0 {
		return false
	}
	return r.Float64() <= p
}
