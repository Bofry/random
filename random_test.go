package random_test

import (
	"math/rand"
	"random"
	"testing"
)

var (
	rng = random.New(rand.NewSource(5489))
)

// func Test_Float64_Test(t *testing.T) {
// 	for n := 0; n < 100; n++ {
// 		fmt.Printf("%10.8f\t", rng.Float64_Test())

// 		if (n % 5) == 4 {
// 			fmt.Println()
// 		}
// 	}
// }

// func Benchmark_Float64_Test(b *testing.B) {
// 	for n := b.N; n > 0; n-- {
// 		rng.Float64_Test()
// 	}
// }

func Benchmark_GoRand(b *testing.B) {
	for n := b.N; n > 0; n-- {
		rng.Int63()
	}
}

func Benchmark_Float64(b *testing.B) {
	for n := b.N; n > 0; n-- {
		rng.Float64()
	}
}

// func Benchmark_GoRand2(b *testing.B) {
// 	rng2 := rand.New(rand.NewSource(5489))
// 	b.RunParallel(func(pb *testing.PB) {

// 		for pb.Next() {
// 			rng2.Int63()
// 		}
// 	})
// }

// func Benchmark_GoRand(b *testing.B) {
// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			rng.Int63()
// 		}
// 	})
// }
