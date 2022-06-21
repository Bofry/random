package random_test

import (
	"math/rand"
	"random/mt19937"
	"testing"

	"random"
)

var (
	seed        = int64(5489)
	rng         = random.New(rand.NewSource(seed))
	rng_mt19937 = random.New(mt19937.New())
)

func TestInt63(t *testing.T) {
	rng.Seed(seed)
	rng.Int63()
}

func TestUint64(t *testing.T) {
	rng.Seed(seed)
	rng.Uint64()
}

func TestUint32(t *testing.T) {
	rng.Seed(seed)
	rng.Uint32()
}

func TestInt31(t *testing.T) {
	rng.Seed(seed)
	rng.Int31()
}

func TestInt(t *testing.T) {
	rng.Seed(seed)
	rng.Int()
}

func TestInt63n(t *testing.T) {
	rng.Seed(seed)
	for n := 0; n < 100; n++ {
		n := rng.Int63n(2)
		if n < 0 || n >= 2 {
			panic("The Range is not between [0, 2).")
		}
	}
}

func TestInt63n_Panic(t *testing.T) {
	rng.Seed(seed)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()

	rng.Int63n(-32757)
}

func TestInt31n(t *testing.T) {
	rng.Seed(seed)
	for i := 0; i < 100; i++ {
		n := rng.Int31n(2)
		if n < 0 || n >= 2 {
			panic("The Range is not between [0, 2).")
		}
	}
}

func TestInt31n_Panic(t *testing.T) {
	rng.Seed(seed)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()

	rng.Int63n(-32757)
}

func TestIntn(t *testing.T) {
	rng.Seed(seed)
	for i := 0; i < 100; i++ {
		n := rng.Intn(2)
		if n < 0 || n >= 2 {
			panic("The Range is not between [0, 2).")
		}
	}
}

func TestIntn_Panic(t *testing.T) {
	rng.Seed(seed)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic.")
		}
	}()

	rng.Intn(-32757)
}

func TestFloat64(t *testing.T) {
	rng.Seed(seed)
	rng.Float64()
}

func TestFloat32(t *testing.T) {
	rng.Seed(seed)
	rng.Float32()
}

func TestPerm(t *testing.T) {
	rng.Seed(seed)
	for i := 0; i < 100; i++ {
		n := rng.Perm(5)

		if len(n) != 5 {
			panic("The length is not 5.")
		}
		for _, v := range n {
			if v < 0 || v >= 5 {
				panic("The Range is not between [0, 5).")
			}
		}
	}
}

func TestShuffle(t *testing.T) {
	rng.Seed(seed)
	n := rng.Perm(5)
	rng.Shuffle(len(n), func(i, j int) {
		n[i], n[j] = n[j], n[i]
	})
}

func TestRead(t *testing.T) {
	rng.Seed(seed)
	n := make([]byte, 5)
	_, err := rng.Read(n)
	if err != nil {
		panic(err)
	}
}
