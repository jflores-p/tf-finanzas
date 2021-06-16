package services

import (
	"github.com/jflores-p/finanzas_tool/domain"
	"math"
	"time"
)

func sumaCostos(costos []domain.Costos, valorNominal float64) (monto float64) {
	for _, costo := range costos {
		if costo.Tipo == "E" {
			monto += costo.Cantidad
		} else {
			monto += valorNominal * costo.Cantidad
		}
	}
	return monto
}

func Descuento(diasXAno float64, letra domain.Letra, costosIniciales, costosFinales []domain.Costos, tasaEfec domain.TasaEfectiva, fechaDescuento time.Time) (letraDesc domain.LetraDescontada) {
	letraDesc.ValoresLetra = letra
	letraDesc.CostosIniciales = costosIniciales
	letraDesc.CostosFinales = costosFinales

	// calculando tasas del periodo de descuento
	diasDescuento := fechaDescuento.Sub(letra.FechaVencimiento).Hours() / 24
	tasaEfecPerDesc := TEfectivaAEfectiva(tasaEfec, diasDescuento)

	letraDesc.TasaDescuento = tasaEfecPerDesc.Tasa / (1 + tasaEfecPerDesc.Tasa)

	letraDesc.ValorNeto = letra.ValorNominal * (1 - letraDesc.TasaDescuento) // VNominal - Descuento
	letraDesc.ValorRecibido = letraDesc.ValorNeto - sumaCostos(costosIniciales, letra.ValorNominal) - letraDesc.ValoresLetra.Retencion
	letraDesc.ValorEntregado = letraDesc.ValorNeto + sumaCostos(costosFinales, letra.ValorNominal) - letraDesc.ValoresLetra.Retencion

	letraDesc.TCEA = math.Pow(letraDesc.ValorEntregado/letraDesc.ValorRecibido, diasXAno/diasDescuento) - 1

	return letraDesc
}
