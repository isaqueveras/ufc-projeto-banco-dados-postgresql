package cidades

import "time"

// DadosCidade modela os dados de uma cidade
type DadosCidade struct {
	ID          *int64     `json:"id,omitempty"`
	Nome        *string    `json:"nome,omitempty"`
	UF          *string    `json:"uf,omitempty"`
	DataCriacao *time.Time `json:"data_criacao,omitempty"`
}

// ListaCidades modela uma lista de cidades
type ListaCidadesRes struct {
	TotalCidades *int64        `json:"total_cidades,omitempty"`
	Dados        []DadosCidade `json:"dados,omitempty"`
}
