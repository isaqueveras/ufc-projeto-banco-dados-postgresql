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

type DadosCidadesRes struct {
	Dados []DadosAvaliacoesCidade `json:"dados,omitempty"`
}

type DadosAvaliacoesCidade struct {
	Nome          *string `json:"nome,omitempty"`
	MediaEstrelas *int64  `json:"media_estrelas,omitempty"`
	QtdAvaliacoes *int64  `json:"qtd_avaliacoes,omitempty"`
}
