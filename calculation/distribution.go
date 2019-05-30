package calculation

import (
	"math"
	"math/rand"
)

func Gamma(shape float64, scale float64) float64 {
	var ksi float64
	if shape == float64(int64(shape)) && shape <= 10.0 {
		for i := 0; int64(i) < int64(shape); i++ {
			ksi -= math.Log(rand.Float64())
		}
	} else if shape+0.5 == float64(int64(shape+0.5)) && shape <= 10.0 {
		a := make([]float64, int64(shape+1.5))
		for i := 0; i < len(a); i++ {
			a[i] = rand.Float64()
		}
		ksi = -math.Log(a[int64(shape+0.5)-1]) * math.Pow(math.Cos(math.Pi*2*a[int64(shape+1.5)-1]), 2)
		for i := 1; int64(i) <= int64(shape-0.5); i++ {
			ksi -= math.Log(a[i-1])
		}
	} else if shape < 1 {
		c := 1 / shape
		d := math.Pow(shape, shape/(1-shape)) * (1 - shape)
		NotAccept := true
		for NotAccept {
			Z := -math.Log(rand.Float64())
			E := -math.Log(rand.Float64())
			ksi = math.Pow(Z, c)
			NotAccept = Z+E <= d+ksi
		}
	} else if shape > 1 {
		b := shape - 1
		c := 3*shape - 0.75
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
	return ksi * scale
}
