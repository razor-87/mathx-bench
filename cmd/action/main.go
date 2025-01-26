package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/razor-87/mathx/vec"
)

const (
	size = 131_072 // 1MB
	n    = 777
)

var (
	measure = flag.Bool("measure", false, "run with measure")
	sink    float64
)

func main() {
	flag.Parse()
	var local float64
	if *measure {
		local = runMeasure()
	} else {
		local = run()
	}
	sink = local
	_ = sink
}

func runMeasure() (ret float64) {
	var (
		nums, nums2, nums3 []float64
		start              time.Time
		elapsed            time.Duration
	)

	start = time.Now()
	for i := 0; i < n; i++ {
		nums2 = vec.Rand[[]float64](size)
	}
	elapsed = time.Since(start)
	fmt.Printf("[vec.Rand] elapsed time: %s, ns/op: %s\n", elapsed, time.Duration(float64(elapsed)/float64(n)))

	start = time.Now()
	for i := 0; i < n; i++ {
		nums = vec.Zeros[[]float64](len(nums2))
	}
	elapsed = time.Since(start)
	fmt.Printf("[vec.Zeros] elapsed time: %s, ns/op: %s\n", elapsed, time.Duration(float64(elapsed)/float64(n)))

	start = time.Now()
	for i := 0; i < n; i++ {
		nums = vec.Ones[[]float64](len(nums))
	}
	elapsed = time.Since(start)
	fmt.Printf("[vec.Ones] elapsed time: %s, ns/op: %s\n", elapsed, time.Duration(float64(elapsed)/float64(n)))

	start = time.Now()
	for i := 0; i < n; i++ {
		nums = vec.Inc[[]float64](len(nums))
	}
	elapsed = time.Since(start)
	fmt.Printf("[vec.Inc] elapsed time: %s, ns/op: %s\n", elapsed, time.Duration(float64(elapsed)/float64(n)))

	start = time.Now()
	for i := 0; i < n; i++ {
		nums = vec.IncBy[[]float64](len(nums), 0.01)
	}
	elapsed = time.Since(start)
	fmt.Printf("[vec.IncBy] elapsed time: %s, ns/op: %s\n", elapsed, time.Duration(float64(elapsed)/float64(n)))

	start = time.Now()
	for i := 0; i < n; i++ {
		nums = vec.Rep[[]float64](len(nums), 0.01)
	}
	elapsed = time.Since(start)
	fmt.Printf("[vec.Rep] elapsed time: %s, ns/op: %s\n", elapsed, time.Duration(float64(elapsed)/float64(n)))

	start = time.Now()
	for i := 0; i < n; i++ {
		vec.Add(nums, nums2)
	}
	elapsed = time.Since(start)
	fmt.Printf("[vec.Add] elapsed time: %s, ns/op: %s\n", elapsed, time.Duration(float64(elapsed)/float64(n)))

	start = time.Now()
	for i := 0; i < n; i++ {
		vec.AddScaled(nums, nums2, 0.999)
	}
	elapsed = time.Since(start)
	fmt.Printf("[vec.AddScaled] elapsed time: %s, ns/op: %s\n", elapsed, time.Duration(float64(elapsed)/float64(n)))

	start = time.Now()
	for i := 0; i < n; i++ {
		vec.Addition(nums, 0.111)
	}
	elapsed = time.Since(start)
	fmt.Printf("[vec.Addition] elapsed time: %s, ns/op: %s\n", elapsed, time.Duration(float64(elapsed)/float64(n)))

	start = time.Now()
	for i := 0; i < n; i++ {
		vec.Scale(nums, 0.999)
	}
	elapsed = time.Since(start)
	fmt.Printf("[vec.Scale] elapsed time: %s, ns/op: %s\n", elapsed, time.Duration(float64(elapsed)/float64(n)))

	start = time.Now()
	for i := 0; i < n; i++ {
		vec.Unit(nums)
	}
	elapsed = time.Since(start)
	fmt.Printf("[vec.Unit] elapsed time: %s, ns/op: %s\n", elapsed, time.Duration(float64(elapsed)/float64(n)))

	start = time.Now()
	for i := 0; i < n; i++ {
		nums3 = vec.Copy(nums2)
	}
	elapsed = time.Since(start)
	fmt.Printf("[vec.Copy] elapsed time: %s, ns/op: %s\n", elapsed, time.Duration(float64(elapsed)/float64(n)))

	start = time.Now()
	for i := 0; i < n; i++ {
		ret = vec.CosSim(nums, nums3)
	}
	elapsed = time.Since(start)
	fmt.Printf("[vec.CosSim] elapsed time: %s, ns/op: %s\n", elapsed, time.Duration(float64(elapsed)/float64(n)))

	start = time.Now()
	for i := 0; i < n; i++ {
		ret = vec.DotProd(nums, nums3)
	}
	elapsed = time.Since(start)
	fmt.Printf("[vec.DotProd] elapsed time: %s, ns/op: %s\n", elapsed, time.Duration(float64(elapsed)/float64(n)))

	start = time.Now()
	for i := 0; i < n; i++ {
		ret = vec.EucDist(nums, nums3)
	}
	elapsed = time.Since(start)
	fmt.Printf("[vec.EucDist] elapsed time: %s, ns/op: %s\n", elapsed, time.Duration(float64(elapsed)/float64(n)))

	start = time.Now()
	for i := 0; i < n; i++ {
		ret = vec.Length(nums)
	}
	elapsed = time.Since(start)
	fmt.Printf("[vec.Length] elapsed time: %s, ns/op: %s\n", elapsed, time.Duration(float64(elapsed)/float64(n)))

	start = time.Now()
	for i := 0; i < n; i++ {
		ret = vec.Sum(nums)
	}
	elapsed = time.Since(start)
	fmt.Printf("[vec.Sum] elapsed time: %s, ns/op: %s\n", elapsed, time.Duration(float64(elapsed)/float64(n)))

	return ret
}

func run() (ret float64) {
	var nums []float64
	nums = vec.Zeros[[]float64](size)
	nums = vec.Ones[[]float64](len(nums))
	nums = vec.Inc[[]float64](len(nums))
	nums = vec.IncBy[[]float64](len(nums), 0.1)
	nums = vec.Rep[[]float64](len(nums), 0.1)
	nums2 := vec.Rand[[]float64](len(nums))
	vec.Add(nums, nums2)
	vec.Addition(nums, 0.1)
	vec.Scale(nums, 0.5)
	vec.Unit(nums)
	nums3 := vec.Copy(nums2)
	ret = vec.CosSim(nums, nums3)
	ret += vec.DotProd(nums, nums3)
	ret += vec.EucDist(nums, nums3)
	ret += vec.Length(nums)
	ret += vec.Sum(nums)

	return ret
}
