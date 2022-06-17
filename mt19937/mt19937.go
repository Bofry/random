package mt19937

type Source struct {
	mt *mt19937
}

func New() *Source {
	return &Source{
		mt: new_mt19937(),
	}
}

// Seed -
// initializes mt[NN] with a seed
func (source *Source) Seed(seed int64) {
	source.mt.init_genrand64(uint64(seed))
}

// SeedbyArray -
// initialize by an array
func (source *Source) SeedbyArray(seeds []uint64) {
	source.mt.init_by_array64(seeds)
}

// Uint64 -
// returns: as an Uint64 [0..2^64-1].
func (source *Source) Uint64() uint64 {
	return source.mt.genrand64_int64()
}

// Int63 -
// returns:	an int64 [-2^63, 2^63-1].
func (source *Source) Int63() int64 {
	return source.mt.genrand64_int63()
}

// RNG64_Real2 -
/* generates a random number on [0,1)-real-interval */
func (source *Source) RNG64_Real2() float64 {
	return float64(source.mt.genrand64_int64()>>11) * (1.0 / 9007199254740992.0)
}

// implement Mersenne Twister 64bit for go
// References: http://www.math.sci.hiroshima-u.ac.jp/m-mat/MT/VERSIONS/C-LANG/mt19937-64.c
/*
   A C-program for MT19937-64 (2004/9/29 version).
   Coded by Takuji Nishimura and Makoto Matsumoto.

   This is a 64-bit version of Mersenne Twister pseudorandom number
   generator.

   Before using, initialize the state by using init_genrand64(seed)
   or init_by_array64(init_key, key_length).

   Copyright (C) 2004, Makoto Matsumoto and Takuji Nishimura,
   All rights reserved.

   Redistribution and use in source and binary forms, with or without
   modification, are permitted provided that the following conditions
   are met:

     1. Redistributions of source code must retain the above copyright
        notice, this list of conditions and the following disclaimer.

     2. Redistributions in binary form must reproduce the above copyright
        notice, this list of conditions and the following disclaimer in the
        documentation and/or other materials provided with the distribution.

     3. The names of its contributors may not be used to endorse or promote
        products derived from this software without specific prior written
        permission.

   THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
   "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
   LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
   A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE COPYRIGHT OWNER OR
   CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
   EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
   PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
   PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
   LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
   NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
   SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

   References:
   T. Nishimura, ``Tables of 64-bit Mersenne Twisters''
     ACM Transactions on Modeling and
     Computer Simulation 10. (2000) 348--357.
   M. Matsumoto and T. Nishimura,
     ``Mersenne Twister: a 623-dimensionally equidistributed
       uniform pseudorandom number generator''
     ACM Transactions on Modeling and
     Computer Simulation 8. (Jan. 1998) 3--30.

   Any feedback is very welcome.
   http://www.math.hiroshima-u.ac.jp/~m-mat/MT/emt.html
   email: m-mat @ math.sci.hiroshima-u.ac.jp (remove spaces)
*/

const (
	_NN      = 312
	_MM      = 156
	_MatrixA = 0xB5026F5AA96619E9 // constant vector a
	_UM      = 0xFFFFFFFF80000000 // Most significant 33 bits
	_LM      = 0x7FFFFFFF         // Least significant 31 bits
)

type mt19937 struct {
	state []uint64
	index int
}

// new_mt19937 -
// return a new instance of the 64bit Mersenne Twister.
func new_mt19937() *mt19937 {
	return &mt19937{
		state: make([]uint64, _NN),
		index: _NN + 1,
	}
}

// init_genrand64 -
/* initializes mt[NN] with a seed */
func (mt *mt19937) init_genrand64(seed uint64) {
	state := mt.state
	state[0] = seed
	for i := uint64(1); i < _NN; i++ {
		state[i] = 6364136223846793005*(state[i-1]^(state[i-1]>>62)) + i
	}
	mt.index = _NN
}

// init_by_array64 -
/* initialize by an array with array-length */
/* init_key is the array for initializing keys */
func (mt *mt19937) init_by_array64(init_key []uint64) {
	mt.init_genrand64(19650218)

	var i, j uint64
	i = 1
	j = 0
	k := len(init_key)
	state := mt.state

	if _NN > k {
		k = _NN
	}
	for ; k > 0; k-- {
		state[i] = (state[i] ^ ((state[i-1] ^ (state[i-1] >> 62)) * 3935559000370003845) + init_key[j] + j)
		i++
		if i >= _NN {
			state[0] = state[_NN-1]
			i = 1
		}
		j++
		if j >= uint64(len(init_key)) {
			j = 0
		}
	}

	for k = _NN - 1; k > 0; k-- {
		state[i] = state[i] ^ ((state[i-1] ^ (state[i-1] >> 62)) * 2862933555777941757) - i
		i++
		if i >= _NN {
			state[0] = state[_NN-1]
			i = 1
		}
	}

	state[0] = 1 << 63
}

// genrand64_int64 -
/* generates a random number on [0, 2^64-1]-interval */
func (mt *mt19937) genrand64_int64() uint64 {
	var x uint64
	state := mt.state[:]

	/* generate NN words at one time */
	if mt.index >= _NN {

		mag01 := []uint64{0, _MatrixA}

		/* if init_genrand64() has not been called, */
		/* a default initial seed is used     */
		if mt.index == _NN+1 {
			mt.init_genrand64(5489)
		}

		var i uint64
		for i = 0; i < _NN-_MM; i++ {
			x = (state[i] & _UM) | (state[i+1] & _LM)
			state[i] = state[i+_MM] ^ (x >> 1) ^ mag01[(x&1)]
		}
		for i = _NN - _MM; i < _NN-1; i++ {
			x = (state[i] & _UM) | (state[i+1] & _LM)
			state[i] = state[int(i)+(_MM-_NN)] ^ (x >> 1) ^ mag01[(x&1)]
		}

		x = (state[_NN-1] & _UM) | (state[0] & _LM)
		state[_NN-1] = state[_MM-1] ^ (x >> 1) ^ mag01[(x&1)]
		mt.index = 0
	}

	x = state[mt.index]
	mt.index++

	x ^= (x >> 29) & 0x5555555555555555
	x ^= (x << 17) & 0x71D67FFFEDA60000
	x ^= (x << 37) & 0xFFF7EEE000000000
	x ^= (x >> 43)

	return x
}

// genrand64_int63 -
/* generates a random number on [-2^63, 2^63-1]-interval */
func (mt *mt19937) genrand64_int63() int64 {
	return int64(mt.genrand64_int64() >> 1)
}
