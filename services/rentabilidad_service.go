package services

import "math"

func VA(nPeriodos int, cok, flujoCaja float64) (sumatoria float64) {
	for i := 1; i <= nPeriodos; i++ {
		var denominador = math.Pow(1+cok, float64(i))
		sumatoria += flujoCaja / denominador
	}
	return sumatoria
}

func VAN(nPeriodos int, inversion, cok, flujoCaja float64) float64{
	return VA(nPeriodos, cok, flujoCaja) - inversion
}
