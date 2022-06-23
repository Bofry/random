package random_test

import "testing"

func Benchmark_Test_ThreadSafe_Seed(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rng.Seed(seed)
		}
	})
}

func Benchmark_Test_ThreadSafe_Int63(b *testing.B) {
	rng_safe.Seed(seed)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rng_safe.Int63()
		}
	})
}

func Benchmark_Test_ThreadSafe_Uint64(b *testing.B) {
	rng_safe.Seed(seed)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rng_safe.Uint64()
		}
	})
}

func Benchmark_Test_ThreadSafe_Float64(b *testing.B) {
	rng_safe.Seed(seed)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rng_safe.Float64()
		}
	})
}
