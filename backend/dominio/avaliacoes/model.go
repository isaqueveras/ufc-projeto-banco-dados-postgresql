package avaliacoes

import (
	"time"
)

// DadosAvaliacao modela os dados para cadastrar uma avaliação
type DadosAvaliacao struct {
	ID          *int64
	Titulo      *string
	Descricao   *string
	QtdEstrela  *int64
	NomePessoa  *string
	EmpresaID   *int64
	Empresa     *Empresa
	DataCriacao *time.Time
}

// ListaAvaliacaos modela uma lista de
type ListaAvaliacaos struct {
	TotalAvaliacoes *int64
	Dados           []DadosAvaliacao
}

type Empresa struct {
	ID   *int64
	Nome *string
}
