package std

import (
	"math"

	"github.com/razor-87/mathx"
)

func VecAdd[T mathx.Vector[E], E mathx.Floaty](v, w T) {
	for i := range v {
		v[i] += w[i]
	}
}

func VecAddScaled[T mathx.Vector[E], E mathx.Floaty](v, w T, c E) {
	for i := range v {
		v[i] += w[i] * c
	}
}

func VecDotProd[T mathx.Vector[E], E mathx.Floaty](v, w T) (ret E) {
	for i := range v {
		ret += v[i] * w[i]
	}
	return ret
}

func VecLength[T mathx.Vector[E], E mathx.Floaty](v T) E {
	return E(math.Sqrt(float64(VecDotProd(v, v))))
}

func VecScale[T mathx.Vector[E], E mathx.Floaty](v T, c E) {
	for i := range v {
		v[i] *= c
	}
}

func VecSum[T mathx.Vector[E], E mathx.Floaty](v T) (ret E) {
	for i := range v {
		ret += v[i]
	}
	return ret
}

func VecUnit[T mathx.Vector[E], E mathx.Floaty](v T) {
	VecScale(v, 1/VecLength(v))
}
