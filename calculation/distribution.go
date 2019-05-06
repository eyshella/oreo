package calculation

import (
	"math"
	"math/rand"
)

func Gamma(k float64) float64 {
	var ksi float64
	if k < 1 {
		c := 1 / k
		d := math.Pow(k, k/(1-k)) * (1 - k)
		Accept := false
		for !Accept {
			Z := -math.Log(rand.Float64())
			E := -math.Log(rand.Float64())
			ksi = math.Pow(Z, c)
			Accept = Z+E <= d+ksi
		}
	} else if k > 1 {
		b := k - 1
		c := 3*k - 0.75
		Accept := false
		for !Accept {
			U := rand.Float64()
			V := rand.Float64()
			W := U * (1 - U)
			Y := math.Pow(c/W, 0.5) * (U - 0.5)
			ksi = b + Y
			if ksi >= 0 {
				Z := 64 * math.Pow(W, 3) * V * V
				Accept = Z <= 1-2*Y*Y/ksi
				if !Accept {
					Accept = math.Log(Z) <= 2*(b*math.Log(ksi/b)-Y)
				}
			}
		}
	} else if k == 1 {
		ksi = -math.Log(rand.Float64())
	}
	return ksi
}
