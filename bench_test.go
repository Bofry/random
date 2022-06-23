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

func Benchmark_Go_Safe_Seed(b *testing.B) {
	for n := b.N; n > 0; n-- {
		rng_safe.Seed(seed)
	}
}

func Benchmark_Go_Safe_Int63(b *testing.B) {
	rng_safe.Seed(seed)
	for n := b.N; n > 0; n-- {
		rng_safe.Int63()
	}
}

func Benchmark_Go_Safe_Uint64(b *testing.B) {
	rng_safe.Seed(seed)
	for n := b.N; n > 0; n-- {
		rng_safe.Uint64()
	}
}

func Benchmark_Go_Safe_Float64(b *testing.B) {
	rng_safe.Seed(seed)
	for n := b.N; n > 0; n-- {
		rng_safe.Float64()
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

func Benchmark_MT19937_Safe_Seed(b *testing.B) {
	for n := b.N; n > 0; n-- {
		rng_safe_mt19937.Seed(seed)
	}
}

func Benchmark_MT19937_Safe_Int63(b *testing.B) {
	rng_safe_mt19937.Seed(seed)
	for n := b.N; n > 0; n-- {
		rng_safe_mt19937.Int63()
	}
}

func Benchmark_MT19937_Safe_Uint64(b *testing.B) {
	rng_safe_mt19937.Seed(seed)
	for n := b.N; n > 0; n-- {
		rng_safe_mt19937.Uint64()
	}
}

func Benchmark_MT19937_Safe_Float64(b *testing.B) {
	rng_safe_mt19937.Seed(seed)
	for n := b.N; n > 0; n-- {
		rng_safe_mt19937.Float64()
	}
}

func Benchmark_Float64w(b *testing.B) {
	ws := []float64{2, 2, 4, 2}
	for n := b.N; n > 0; n-- {
		rng.Float64w(ws)
	}
}

func Benchmark_Float32w(b *testing.B) {
	ws := []float32{2, 2, 4, 2}
	for n := b.N; n > 0; n-- {
		rng.Float32w(ws)
	}
}

func Benchmark_Uint64w(b *testing.B) {
	ws := []uint64{2, 2, 4, 2}
	for n := b.N; n > 0; n-- {
		rng.Uint64w(ws)
	}
}

func Benchmark_Uint32w(b *testing.B) {
	ws := []uint32{2, 2, 4, 2}
	for n := b.N; n > 0; n-- {
		rng.Uint32w(ws)
	}
}

func Benchmark_Int64w(b *testing.B) {
	ws := []int64{2, 2, 4, 2}
	for n := b.N; n > 0; n-- {
		rng.Int64w(ws)
	}
}

func Benchmark_Int32w(b *testing.B) {
	ws := []int32{2, 2, 4, 2}
	for n := b.N; n > 0; n-- {
		rng.Int32w(ws)
	}
}

func Benchmark_Intw(b *testing.B) {
	ws := []int{2, 2, 4, 2}
	for n := b.N; n > 0; n-- {
		rng.Intw(ws)
	}
}
