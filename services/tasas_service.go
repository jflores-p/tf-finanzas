package services

import (
	"github.com/jflores-p/finanzas_tool/domain"
	"math"
)

// TEfectivaAEfectiva transforma una tasa efectiva de un periodo 'n1' a una de periodo 'n2'
func TEfectivaAEfectiva(tasaEfec1 domain.TasaEfectiva, n2 float64) (tasaEfec2 domain.TasaEfectiva) {
	exponente := n2 / tasaEfec1.PlazoDias

	tasaEfec2.Tasa = (math.Pow(1+tasaEfec1.Tasa, exponente) - 1) * 100
	tasaEfec2.PlazoDias = n2

	return tasaEfec2
}

// TNominalAEfectiva transforma una tasa nominal de un periodo 'periodoNominal' a
// todos los periodos son en dias
func TNominalAEfectiva(tasaNom domain.TasaNominal, periodoEfectivo float64) (tasaEfec domain.TasaEfectiva) {
	m := tasaNom.Plazo / tasaNom.Capitalizacion
	n := periodoEfectivo / tasaNom.Capitalizacion

	tasaEfec.Tasa = math.Pow(1+(tasaNom.Tasa/m), n) - 1
	tasaEfec.PlazoDias = periodoEfectivo

	return tasaEfec
}
