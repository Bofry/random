package mt19937_test

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"

	"github.com/bofry/random/mt19937"
)

// http://www.math.sci.hiroshima-u.ac.jp/m-mat/MT/mt19937-64.out.txt

func TestOutputMT19937_64_1000(t *testing.T) {
	init := []uint64{0x12345, 0x23456, 0x34567, 0x45678}

	mt := mt19937.New()
	mt.SeedbyArray(init)
	output := fmt.Sprint("ref: http://www.math.sci.hiroshima-u.ac.jp/m-mat/MT/mt19937-64.out.txt\n\n")

	output += fmt.Sprint("1000 outputs of genrand64_int64()\n")
	i := 0
	for i = 0; i < 1000; i++ {
		output += fmt.Sprintf("%20d\t", mt.Uint64())
		if (i % 5) == 4 {
			output += fmt.Sprintln()
		}
	}

	output += fmt.Sprint("\n1000 outputs of genrand64_real2()\n")
	for i = 0; i < 1000; i++ {
		output += fmt.Sprintf("%10.8f\t", mt.RNG64_Real2())
		if (i % 5) == 4 {
			output += fmt.Sprintln()
		}
	}

	err := ioutil.WriteFile("./mt19937-64.out.txt", []byte(output), 0644)
	if err != nil {
		panic(err)
	}
	print(output)
}

func Benchmark_RNG_Int63(b *testing.B) {
	rng := rand.New(rand.NewSource(5489))

	for n := b.N; n > 0; n-- {
		rng.Int63()
	}
}

func Benchmark_RNG_Uint64(b *testing.B) {
	rng := rand.New(rand.NewSource(5489))

	for n := b.N; n > 0; n-- {
		rng.Uint64()
	}
}
func Benchmark_RNG_RuntimeUint64(b *testing.B) {
	rng := rand.New(rand.NewSource(5489))

	for n := b.N; n > 0; n-- {
		rng.Uint64()
	}
}

func Benchmark_RMT19937_New(b *testing.B) {
	for n := b.N; n > 0; n-- {
		rand.New(mt19937.New())
	}
}

func Benchmark_RMT19937_Seed(b *testing.B) {
	rng := rand.New(mt19937.New())

	for n := b.N; n > 0; n-- {
		rng.Seed(5489)
	}
}

func Benchmark_RMT19937_Int63(b *testing.B) {
	rng := rand.New(mt19937.New())

	for n := b.N; n > 0; n-- {
		rng.Int63()
	}
}

func Benchmark_RMT19937_Uint64(b *testing.B) {
	rng := rand.New(mt19937.New())

	for n := b.N; n > 0; n-- {
		rng.Uint64()
	}
}
