package calculation

import (
	"math"
	"math/rand"
)

func Gamma(k float64) float64 {
	var ksi float64
	if k == float64(int64(k)) && k <= 10.0 {
		for i := 0; int64(i) < int64(k); i++ {
			ksi -= math.Log(rand.Float64())
		}
	} else if k+0.5 == float64(int64(k+0.5)) && k <= 10.0 {
		a := make([]float64, int64(k+1.5))
		for i := 0; i < len(a); i++ {
			a[i] = rand.Float64()
		}
		ksi = -math.Log(a[int64(k+0.5)-1]) * math.Pow(math.Cos(math.Pi*2*a[int64(k+1.5)-1]), 2)
		for i := 1; int64(i) <= int64(k-0.5); i++ {
			ksi -= math.Log(a[i-1])
		}
	} else if k < 1 {
		c := 1 / k
		d := math.Pow(k, k/(1-k)) * (1 - k)
		NotAccept := true
		for NotAccept {
			Z := -math.Log(rand.Float64())
			E := -math.Log(rand.Float64())
			ksi = math.Pow(Z, c)
			NotAccept = Z+E <= d+ksi
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
	}
	return ksi
}
