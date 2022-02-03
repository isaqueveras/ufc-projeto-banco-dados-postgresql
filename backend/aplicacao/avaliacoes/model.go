package avaliacoes

import "time"

// DadosAvaliacaoReq modela os dados para cadastrar uma avaliação
type DadosAvaliacaoReq struct {
	ID          *int64     `json:"id,omitempty"`
	Titulo      *string    `json:"titulo,omitempty"`
	Descricao   *string    `json:"descricao,omitempty"`
	QtdEstrela  *int64     `json:"qtd_estrela,omitempty"`
	NomePessoa  *string    `json:"nome_pessoa,omitempty"`
	EmpresaID   *int64     `json:"empresa_id,omitempty"`
	DataCriacao *time.Time `json:"data_criacao,omitempty"`
}
