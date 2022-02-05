package cidades

import "time"

// DadosCidade modela os dados de uma cidade
type DadosCidade struct {
	ID          *int64
	Nome        *string
	UF          *string
	DataCriacao *time.Time
}

// ListaCidades modela uma lista de cidades
type ListaCidades struct {
	TotalCidades *int64
	Dados        []DadosCidade
}
