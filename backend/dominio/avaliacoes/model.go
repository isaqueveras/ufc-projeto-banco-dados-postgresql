package avaliacoes

import "time"

// DadosAvaliacao modela os dados para cadastrar uma avaliação
type DadosAvaliacao struct {
	ID          *int64
	Titulo      *string
	Descricao   *string
	QtdEstrela  *int64
	NomePessoa  *string
	EmpresaID   *int64
	DataCriacao *time.Time
}
