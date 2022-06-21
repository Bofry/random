package random_test

import "testing"

func Benchmark_Go_Seed(b *testing.B) {
	for n := b.N; n > 0; n-- {
		rng.Seed(seed)
	}
}

func Benchmark_Go_Int63(b *testing.B) {
	rng.Seed(seed)
	for n := b.N; n > 0; n-- {
		rng.Int63()
	}
}

func Benchmark_Go_Uint64(b *testing.B) {
	rng.Seed(seed)
	for n := b.N; n > 0; n-- {
		rng.Uint64()
	}
}

func Benchmark_Go_Float64(b *testing.B) {
	rng.Seed(seed)
	for n := b.N; n > 0; n-- {
		rng.Float64()
	}
}

func Benchmark_MT19937_Seed(b *testing.B) {
	for n := b.N; n > 0; n-- {
		rng_mt19937.Seed(seed)
	}
}

func Benchmark_MT19937_Int63(b *testing.B) {
	rng_mt19937.Seed(seed)
	for n := b.N; n > 0; n-- {
		rng_mt19937.Int63()
	}
}

func Benchmark_MT19937_Uint64(b *testing.B) {
	rng_mt19937.Seed(seed)
	for n := b.N; n > 0; n-- {
		rng_mt19937.Uint64()
	}
}

func Benchmark_MT19937_Float64(b *testing.B) {
	rng_mt19937.Seed(seed)
	for n := b.N; n > 0; n-- {
		rng_mt19937.Float64()
	}
}
