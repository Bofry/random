package main

import (
	"fmt"

	"github.com/bofry/random"
	"github.com/bofry/random/mt19937"
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
