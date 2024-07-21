package baseline

import (
	"math"

	"github.com/razor-87/mathx"
)

func vecAdd[T mathx.Vector[E], E mathx.Floaty](v, w T) {
	for i := range v {
		v[i] += w[i]
	}
}

func vecAddScaled[T mathx.Vector[E], E mathx.Floaty](v, w T, c E) {
	for i := range v {
		v[i] += w[i] * c
	}
}

func vecDotProd[T mathx.Vector[E], E mathx.Floaty](v, w T) (ret E) {
	for i := range v {
		ret += v[i] * w[i]
	}
	return ret
}

func vecLength[T mathx.Vector[E], E mathx.Floaty](v T) E {
	return E(math.Sqrt(float64(vecDotProd(v, v))))
}

func vecScale[T mathx.Vector[E], E mathx.Floaty](v T, c E) {
	for i := range v {
		v[i] *= c
	}
}

func vecSum[T mathx.Vector[E], E mathx.Floaty](v T) (ret E) {
	for i := range v {
		ret += v[i]
	}
	return ret
}

func vecUnit[T mathx.Vector[E], E mathx.Floaty](v T) {
	vecScale(v, 1/vecLength(v))
}
