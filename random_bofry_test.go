package random_test

import (
	"testing"
)

var (
	testRound = 100000000
)

func TestFloat64w(t *testing.T) {
	rng.Seed(seed)
	ws := []float64{2, 2, 2, 4}
	for n := 0; n < 100; n++ {
		rng.Float64w(ws)
	}
}

func TestFloat64w_zeroarray(t *testing.T) {
	rng.Seed(seed)
	ws := []float64{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()
	for n := 0; n < 100; n++ {
		rng.Float64w(ws)
	}
}

func TestFloat64w_negative(t *testing.T) {
	rng.Seed(seed)
	ws := []float64{-1, 2, 3, 4}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()
	for n := 0; n < 100; n++ {
		rng.Float64w(ws)
	}
}

func TestFloat32w(t *testing.T) {
	rng.Seed(seed)
	ws := []float32{2, 2, 2, 4}
	for n := 0; n < 100; n++ {
		rng.Float32w(ws)
	}
}

func TestFloat32w_zeroarray(t *testing.T) {
	rng.Seed(seed)
	ws := []float32{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()
	for n := 0; n < 100; n++ {
		rng.Float32w(ws)
	}
}

func TestFloat32w_negative(t *testing.T) {
	rng.Seed(seed)
	ws := []float32{-1, 2, 3, 4}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()
	for n := 0; n < 100; n++ {
		rng.Float32w(ws)
	}
}

func TestUint64w(t *testing.T) {
	rng.Seed(seed)
	ws := []uint64{1, 2, 3, 4}
	for n := 0; n < 100; n++ {
		rng.Uint64w(ws)
	}
}

func TestUint64w_zeroarray(t *testing.T) {
	rng.Seed(seed)
	ws := []uint64{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()
	for n := 0; n < 100; n++ {
		rng.Uint64w(ws)
	}
}

func TestUint32w(t *testing.T) {
	rng.Seed(seed)
	ws := []uint32{1, 2, 3, 4}
	for n := 0; n < 100; n++ {
		rng.Uint32w(ws)
	}
}

func TestUint32w_zeroarray(t *testing.T) {
	rng.Seed(seed)
	ws := []uint32{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()
	for n := 0; n < 100; n++ {
		rng.Uint32w(ws)
	}
}

func TestInt64w(t *testing.T) {
	rng.Seed(seed)
	ws := []int64{1, 2, 3, 4}
	for n := 0; n < 100; n++ {
		rng.Int64w(ws)
	}
}

func TestInt64w_zeroarray(t *testing.T) {
	rng.Seed(seed)
	ws := []int64{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()
	for n := 0; n < 100; n++ {
		rng.Int64w(ws)
	}
}

func TestInt64w_negative(t *testing.T) {
	rng.Seed(seed)
	ws := []int64{-1, 2, 3, 4}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()
	for n := 0; n < 100; n++ {
		rng.Int64w(ws)
	}
}

func TestInt32w(t *testing.T) {
	rng.Seed(seed)
	ws := []int32{1, 2, 3, 4}
	for n := 0; n < 100; n++ {
		rng.Int32w(ws)
	}
}

func TestInt32w_zeroarray(t *testing.T) {
	rng.Seed(seed)
	ws := []int32{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()
	for n := 0; n < 100; n++ {
		rng.Int32w(ws)
	}
}

func TestInt32w_negative(t *testing.T) {
	rng.Seed(seed)
	ws := []int32{-1, 2, 3, 4}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()
	for n := 0; n < 100; n++ {
		rng.Int32w(ws)
	}
}

func TestIntw(t *testing.T) {
	rng.Seed(seed)
	ws := []int{1, 2, 3, 4}
	for n := 0; n < 100; n++ {
		rng.Intw(ws)
	}
}

func TestIntw_zeroarray(t *testing.T) {
	rng.Seed(seed)
	ws := []int{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()
	for n := 0; n < 100; n++ {
		rng.Intw(ws)
	}
}

func TestIntw_negative(t *testing.T) {
	rng.Seed(seed)
	ws := []int{-1, 2, 3, 4}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()
	for n := 0; n < 100; n++ {
		rng.Intw(ws)
	}
}

func TestUint64n(t *testing.T) {
	rng.Seed(seed)
	for n := 0; n < 100; n++ {
		n := rng.Uint64n(2)
		if n < 0 || n > 2 {
			panic("The Range is not between [0, 2).")
		}
	}
}

func TestUint32n(t *testing.T) {
	rng.Seed(seed)
	for n := 0; n < 100; n++ {
		n := rng.Uint32n(2)
		if n < 0 || n > 2 {
			panic("The Range is not between [0, 2).")
		}
	}
}

func TestFloat64n(t *testing.T) {
	rng.Seed(seed)
	for n := 0; n < 100; n++ {
		n := rng.Float64n(2)
		if n < 0 || n > 2 {
			panic("The Range is not between [0, 2).")
		}
	}
}

func TestFloat32n(t *testing.T) {
	rng.Seed(seed)
	for n := 0; n < 100; n++ {
		n := rng.Float32n(2)
		if n < 0 || n > 2 {
			panic("The Range is not between [0, 2).")
		}
	}
}
