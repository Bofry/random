# Random

`Random` is a pseudo-random number generator(PRNG). its functions come from standard Go functions from package `math/rand` and copy the some functions from [gosl/rnd](https://github.com/cpmech/gosl/tree/main/rnd).

## Benckmark

```console

Running tool: go test -benchmem -bench .

goos: darwin
goarch: amd64
pkg: random
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_Go_Seed-12                    133672          7807 ns/op      0 B/op        0 allocs/op
Benchmark_Go_Int63-12                   374727135       3.227 ns/op     0 B/op        0 allocs/op
Benchmark_Go_Uint64-12                  322754622       3.706 ns/op     0 B/op        0 allocs/op
Benchmark_Go_Float64-12                 277221798       4.316 ns/op     0 B/op        0 allocs/op
Benchmark_Go_Safe_Seed-12               155568          7775 ns/op      0 B/op        0 allocs/op
Benchmark_Go_Safe_Int63-12              83138834        14.58 ns/op     0 B/op        0 allocs/op
Benchmark_Go_Safe_Uint64-12             79397599        15.32 ns/op     0 B/op        0 allocs/op
Benchmark_Go_Safe_Float64-12            78361706        14.87 ns/op     0 B/op        0 allocs/op
Benchmark_MT19937_Seed-12               1690099         710.6 ns/op     0 B/op        0 allocs/op
Benchmark_MT19937_Int63-12              225984925       5.316 ns/op     0 B/op        0 allocs/op
Benchmark_MT19937_Uint64-12             194375786       6.178 ns/op     0 B/op        0 allocs/op
Benchmark_MT19937_Float64-12            178710583       6.718 ns/op     0 B/op        0 allocs/op
Benchmark_MT19937_Safe_Seed-12          1692319         720.5 ns/op     0 B/op        0 allocs/op
Benchmark_MT19937_Safe_Int63-12         76736385        15.83 ns/op     0 B/op        0 allocs/op
Benchmark_MT19937_Safe_Uint64-12        71111898        16.84 ns/op     0 B/op        0 allocs/op
Benchmark_MT19937_Safe_Float64-12       71303149        16.58 ns/op     0 B/op        0 allocs/op
Benchmark_Float64w-12                   57282123        20.13 ns/op     0 B/op        0 allocs/op
Benchmark_Float32w-12                   56282934        21.14 ns/op     0 B/op        0 allocs/op
Benchmark_Uint64w-12                    48131596        25.00 ns/op     0 B/op        0 allocs/op
Benchmark_Uint32w-12                    58037516        21.55 ns/op     0 B/op        0 allocs/op
Benchmark_Int64w-12                     37663840        32.38 ns/op     0 B/op        0 allocs/op
Benchmark_Int32w-12                     46440220        22.44 ns/op     0 B/op        0 allocs/op
Benchmark_Intw-12                       53368224        22.63 ns/op     0 B/op        0 allocs/op
Benchmark_Test_ThreadSafe_Seed-12       1607907         763.0 ns/op     0 B/op        0 allocs/op
Benchmark_Test_ThreadSafe_Int63-12      22007811        52.48 ns/op     0 B/op        0 allocs/op
Benchmark_Test_ThreadSafe_Uint64-12     22034865        54.91 ns/op     0 B/op        0 allocs/op
Benchmark_Test_ThreadSafe_Float64-12    21995600        53.93 ns/op     0 B/op        0 allocs/op
PASS
coverage: 29.1% of statements
ok   random 38.310s

```
