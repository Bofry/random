package random_test

import (
	"testing"
)

func TestInt63r(t *testing.T) {
	rng.Seed(seed)
	for n := 0; n < 100; n++ {
		n := rng.Int63r(0, 2)
		if n < 0 || n > 2 {
			panic("The Range is not between [0, 2].")
		}
	}
}

func TestInt63s(t *testing.T) {
	rng.Seed(seed)
	n := make([]int64, 0)
	rng.Int63s(n, 0, 2)
	n = make([]int64, 5)
	for i := 0; i < 100; i++ {
		rng.Int63s(n, 0, 2)

		if len(n) != 5 {
			panic("The length is not 5.")
		}
		for _, v := range n {
			if v < 0 || v > 2 {
				panic("The Range is not between [0, 2].")
			}
		}
	}
}

func TestInt63Shuffle(t *testing.T) {
	rng.Seed(seed)
	n := []int64{0, 1, 2, 3, 4}
	rng.Int63Shuffle(n)
}

func TestUint32r(t *testing.T) {
	rng.Seed(seed)
	for n := 0; n < 100; n++ {
		n := rng.Uint32r(0, 2)
		if n < 0 || n > 2 {
			panic("The Range is not between [0, 2].")
		}
	}
}

func TestUint32s(t *testing.T) {
	rng.Seed(seed)
	n := make([]uint32, 0)
	rng.Uint32s(n, 0, 2)
	n = make([]uint32, 5)
	for i := 0; i < 100; i++ {
		rng.Uint32s(n, 0, 2)

		if len(n) != 5 {
			panic("The length is not 5.")
		}
		for _, v := range n {
			if v < 0 || v > 2 {
				panic("The Range is not between [0, 2].")
			}
		}
	}
}

func TestUint32Shuffle(t *testing.T) {
	rng.Seed(seed)
	n := []uint32{0, 1, 2, 3, 4}
	rng.Uint32Shuffle(n)
}

func TestUint64r(t *testing.T) {
	rng.Seed(seed)
	for n := 0; n < 100; n++ {
		n := rng.Uint64r(0, 2)
		if n < 0 || n > 2 {
			panic("The Range is not between [0, 2].")
		}
	}
}

func TestUint64s(t *testing.T) {
	rng.Seed(seed)
	n := make([]uint64, 0)
	rng.Uint64s(n, 0, 2)
	n = make([]uint64, 5)
	for i := 0; i < 100; i++ {
		rng.Uint64s(n, 0, 2)

		if len(n) != 5 {
			panic("The length is not 5.")
		}
		for _, v := range n {
			if v < 0 || v > 2 {
				panic("The Range is not between [0, 2].")
			}
		}
	}
}

func TestUint64Shuffle(t *testing.T) {
	rng.Seed(seed)
	n := []uint64{0, 1, 2, 3, 4}
	rng.Uint64Shuffle(n)
}

func TestInt31r(t *testing.T) {
	rng.Seed(seed)
	for n := 0; n < 100; n++ {
		n := rng.Int31r(0, 2)
		if n < 0 || n > 2 {
			panic("The Range is not between [0, 2].")
		}
	}
}

func TestInt31s(t *testing.T) {
	rng.Seed(seed)
	n := make([]int32, 0)
	rng.Int31s(n, 0, 2)
	n = make([]int32, 5)
	for i := 0; i < 100; i++ {
		rng.Int31s(n, 0, 2)

		if len(n) != 5 {
			panic("The length is not 5.")
		}
		for _, v := range n {
			if v < 0 || v > 2 {
				panic("The Range is not between [0, 2].")
			}
		}
	}
}

func TestInt31Shuffle(t *testing.T) {
	rng.Seed(seed)
	n := []int32{0, 1, 2, 3, 4}
	rng.Int31Shuffle(n)
}

func TestIntr(t *testing.T) {
	rng.Seed(seed)
	for n := 0; n < 100; n++ {
		n := rng.Intr(0, 2)
		if n < 0 || n > 2 {
			panic("The Range is not between [0, 2].")
		}
	}
}

func TestInts(t *testing.T) {
	rng.Seed(seed)
	n := make([]int, 0)
	rng.Ints(n, 0, 2)
	n = make([]int, 5)
	for i := 0; i < 100; i++ {
		rng.Ints(n, 0, 2)

		if len(n) != 5 {
			panic("The length is not 5.")
		}
		for _, v := range n {
			if v < 0 || v > 2 {
				panic("The Range is not between [0, 2].")
			}
		}
	}
}

func TestIntShuffle(t *testing.T) {
	rng.Seed(seed)
	n := []int{0, 1, 2, 3, 4}
	rng.IntShuffle(n)
}

func TestFloat64r(t *testing.T) {
	rng.Seed(seed)
	for n := 0; n < 100; n++ {
		n := rng.Float64r(0, 2)
		if n < 0 || n > 2 {
			panic("The Range is not between [0, 2].")
		}
	}
}

func TestFloat64s(t *testing.T) {
	rng.Seed(seed)
	n := make([]float64, 0)
	rng.Float64s(n, 0, 2)
	n = make([]float64, 5)
	for i := 0; i < 100; i++ {
		rng.Float64s(n, 0, 2)

		if len(n) != 5 {
			panic("The length is not 5.")
		}
		for _, v := range n {
			if v < 0 || v > 2 {
				panic("The Range is not between [0, 2].")
			}
		}
	}
}

func TestFloat64Shuffle(t *testing.T) {
	rng.Seed(seed)
	n := []float64{0, .1, .2, .3, .4}
	rng.Float64Shuffle(n)
}

func TestFloat32r(t *testing.T) {
	rng.Seed(seed)
	for n := 0; n < 100; n++ {
		n := rng.Float32r(0, 2)
		if n < 0 || n > 2 {
			panic("The Range is not between [0, 2].")
		}
	}
}

func TestFloat32s(t *testing.T) {
	rng.Seed(seed)
	n := make([]float32, 0)
	rng.Float32s(n, 0, 2)
	n = make([]float32, 5)
	for i := 0; i < 100; i++ {
		rng.Float32s(n, 0, 2)

		if len(n) != 5 {
			panic("The length is not 5.")
		}
		for _, v := range n {
			if v < 0 || v > 2 {
				panic("The Range is not between [0, 2].")
			}
		}
	}
}

func TestFloat32Shuffle(t *testing.T) {
	rng.Seed(seed)
	n := []float32{0, .1, .2, .3, .4}
	rng.Float32Shuffle(n)
}

func TestFlipCoin(t *testing.T) {
	rng.Seed(seed)
	for i := 0; i < 100; i++ {
		rng.FlipCoin(.5)
	}
	if rng.FlipCoin(0) {
		panic("p == 0 is not false")
	}
	if !rng.FlipCoin(1) {
		panic("p == 1 is not true")
	}
}
