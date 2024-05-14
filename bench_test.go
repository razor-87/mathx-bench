package bench

import (
	"testing"
	"time"

	"github.com/razor-87/mathx/vec"
)

var (
	bytes = []string{"~16KB", "~64KB", "~256KB", "~1MB", "~4MB", "~16MB", "~64MB", "~128MB", "~256MB"}
	sizes = []int{2047, 8191, 32_767, 131_071, 524_287, 2_097_151, 8_388_607, 16_777_215, 33_554_431}
	vecs  = make(map[int][]float64, len(sizes))
)

var prepared bool = func() bool {
	for _, size := range sizes {
		vecs[size] = vec.Rand[[]float64](size)
	}
	return true
}()

var (
	sink    float64
	sinkVec []float64
)

func BenchmarkVecSum(b *testing.B) {
	for !prepared {
		time.Sleep(2 * time.Second)
	}
	var (
		bs int64
		xs []float64
	)
	for i, size := range sizes {
		bs = int64(size * 8)
		xs = vecs[size]
		b.Run(bytes[i], func(b *testing.B) {
			b.SetBytes(bs)
			var local float64
			for i := 0; i < b.N; i++ {
				local = vec.Sum(xs)
			}
			sink = local
		})
	}
}

func BenchmarkVecDotProd(b *testing.B) {
	for !prepared {
		time.Sleep(2 * time.Second)
	}
	var (
		bs     int64
		xs, ys []float64
	)
	for i, size := range sizes {
		bs = int64(size * 8)
		xs = vecs[size]
		ys = vec.Copy(xs)
		b.Run(bytes[i], func(b *testing.B) {
			b.SetBytes(bs)
			var local float64
			for i := 0; i < b.N; i++ {
				local = vec.DotProd(xs, ys)
			}
			sink = local
		})
	}
}

func BenchmarkVecLength(b *testing.B) {
	for !prepared {
		time.Sleep(2 * time.Second)
	}
	var (
		bs int64
		xs []float64
	)
	for i, size := range sizes {
		bs = int64(size * 8)
		xs = vecs[size]
		b.Run(bytes[i], func(b *testing.B) {
			b.SetBytes(bs)
			var local float64
			for i := 0; i < b.N; i++ {
				local = vec.Length(xs)
			}
			sink = local
		})
	}
}

func BenchmarkVecAdd(b *testing.B) {
	for !prepared {
		time.Sleep(2 * time.Second)
	}
	var (
		bs     int64
		xs, ys []float64
	)
	for i, size := range sizes {
		bs = int64(size * 8)
		xs = vecs[size]
		ys = vec.Copy(xs)
		b.Run(bytes[i], func(b *testing.B) {
			b.SetBytes(bs)
			var local []float64
			for i := 0; i < b.N; i++ {
				vec.Add(xs, ys)
				local = xs
			}
			sinkVec = local
		})
	}
}

func BenchmarkVecAddScaled(b *testing.B) {
	for !prepared {
		time.Sleep(2 * time.Second)
	}
	var (
		bs     int64
		xs, ys []float64
	)
	for i, size := range sizes {
		bs = int64(size * 8)
		xs = vecs[size]
		ys = vec.Copy(xs)
		b.Run(bytes[i], func(b *testing.B) {
			b.SetBytes(bs)
			var local []float64
			for i := 0; i < b.N; i++ {
				vec.AddScaled(xs, ys, 0.999)
				local = xs
			}
			sinkVec = local
		})
	}
}

func BenchmarkVecScale(b *testing.B) {
	for !prepared {
		time.Sleep(2 * time.Second)
	}
	var (
		bs int64
		xs []float64
	)
	for i, size := range sizes {
		bs = int64(size * 8)
		xs = vecs[size]
		b.Run(bytes[i], func(b *testing.B) {
			b.SetBytes(bs)
			var local []float64
			for i := 0; i < b.N; i++ {
				vec.Scale(xs, 0.999)
				local = xs
			}
			sinkVec = local
		})
	}
}

func BenchmarkVecUnit(b *testing.B) {
	for !prepared {
		time.Sleep(2 * time.Second)
	}
	var (
		bs int64
		xs []float64
	)
	for i, size := range sizes {
		bs = int64(size * 8)
		xs = vecs[size]
		b.Run(bytes[i], func(b *testing.B) {
			b.SetBytes(bs)
			var local []float64
			for i := 0; i < b.N; i++ {
				vec.Unit(xs)
				local = xs
			}
			sinkVec = local
		})
	}
}
