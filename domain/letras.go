package domain

import "time"

type Letra struct {
	FechaGiro        time.Time
	FechaVencimiento time.Time
	ValorNominal     float64
	NombrePagador    string
	Retencion        float64
}

type LetraDescontada struct {
	ValoresLetra    Letra
	CostosIniciales []Costos
	CostosFinales   []Costos
	TasaDescuento   float64
	ValorNeto       float64
	ValorRecibido   float64
	ValorEntregado  float64
	TCEA            float64
}

type Costos struct {
	Tipo     string
	Cantidad float64
}
