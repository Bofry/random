# random

`random` is a pseudo-random number generator(PRNG). its functions come from standard Go functions from package `math/rand` and copy the some functions from [gosl/rnd](https://github.com/cpmech/gosl/tree/main/rnd).

---

## Install

```console
go get -u -v github.com/bofry/random
```

## Usage

Let's start with a trivial example:

```go
package main

import (
    "fmt"

    "[github.com/bofry/random](https://github.com/bofry/random)"
    "[github.com/bofry/random/mt19937](https://github.com/bofry/random/mt19937)"
)

func main() {
    rng := random.New(mt19937.New())

    // range [0,10](inclusive)
    println("Int63r(0,10)(inclusive) return", rng.Int63r(0, 10))

    // range [0,10)(not inclusive)
    println("Int63n(10)(not inclusive) return", rng.Int63n(10))

    // Random pick with weight
    weights := []int64{1, 2, 3, 4}
    stat := make([]float64, len(weights))
    round := 10000000
    for i := 1; i < round; i++ {
        stat[rng.Int64w(weights)]++
    }
    for _, v := range stat {
        fmt.Printf("%.4f \t", v/float64(round))
    }
}
```

## Output

```console
go run app.
Int63r(0,10)(inclusive) return 7
Int63n(10)(not inclusive) return 4
0.1002  0.2001  0.3000  0.3997  %    
```

## Benckmark

```console
Running tool: go test -benchmem -bench .

goos: darwin
goarch: arm64
pkg: [github.com/bofry/random](https://github.com/bofry/random)
Benchmark_Go_Seed                         107834             10987 ns/op               0 B/op          0 allocs/op
Benchmark_Go_Int63                      414903625                2.873 ns/op           0 B/op          0 allocs/op
Benchmark_Go_Uint64                     379880875                3.154 ns/op           0 B/op          0 allocs/op
Benchmark_Go_Float64                    351158229                3.432 ns/op           0 B/op          0 allocs/op
Benchmark_Go_Safe_Seed                    111523             10904 ns/op               0 B/op          0 allocs/op
Benchmark_Go_Safe_Int63                 83442944                13.88 ns/op            0 B/op          0 allocs/op
Benchmark_Go_Safe_Uint64                81863529                13.91 ns/op            0 B/op          0 allocs/op
Benchmark_Go_Safe_Float64               83060790                13.94 ns/op            0 B/op          0 allocs/op
Benchmark_MT19937_Seed                   1000000              1032 ns/op               0 B/op          0 allocs/op
Benchmark_MT19937_Int63                 318810306                3.752 ns/op           0 B/op          0 allocs/op
Benchmark_MT19937_Uint64                296913031                4.033 ns/op           0 B/op          0 allocs/op
Benchmark_MT19937_Float64               272669924                4.288 ns/op           0 B/op          0 allocs/op
Benchmark_MT19937_Safe_Seed              1000000              1072 ns/op               0 B/op          0 allocs/op
Benchmark_MT19937_Safe_Int63            74343490                15.08 ns/op            0 B/op          0 allocs/op
Benchmark_MT19937_Safe_Uint64           75081927                14.88 ns/op            0 B/op          0 allocs/op
Benchmark_MT19937_Safe_Float64          78464281                14.89 ns/op            0 B/op          0 allocs/op
Benchmark_Float64w                      60282195                17.16 ns/op            0 B/op          0 allocs/op
Benchmark_Float32w                      31604313                36.71 ns/op           32 B/op          1 allocs/op
Benchmark_Uint64w                       82728193                14.20 ns/op            0 B/op          0 allocs/op
Benchmark_Uint32w                       34264460                35.76 ns/op           32 B/op          1 allocs/op
Benchmark_Int64w                        30533314                38.30 ns/op           32 B/op          1 allocs/op
Benchmark_Int32w                        20278176                58.50 ns/op           64 B/op          2 allocs/op
Benchmark_Intw                          30103322                38.97 ns/op           32 B/op          1 allocs/op
Benchmark_Test_ThreadSafe_Seed            111788             10572 ns/op               0 B/op          0 allocs/op
Benchmark_Test_ThreadSafe_Int63         83918524                13.91 ns/op            0 B/op          0 allocs/op
Benchmark_Test_ThreadSafe_Uint64        83163926                13.86 ns/op            0 B/op          0 allocs/op
Benchmark_Test_ThreadSafe_Float64       83615436                13.90 ns/op            0 B/op          0 allocs/op
PASS
coverage: 20.3% of statements
ok      [github.com/bofry/random](https://github.com/bofry/random) 34.390s
```
