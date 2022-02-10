package avaliacoes

import (
	"time"
)

// DadosAvaliacao modela os dados para cadastrar uma avaliação
type DadosAvaliacao struct {
	ID          *int64     `json:"id,omitempty"`
	Titulo      *string    `json:"titulo,omitempty"`
	Descricao   *string    `json:"descricao,omitempty"`
	QtdEstrela  *int64     `json:"qtd_estrela,omitempty"`
	NomePessoa  *string    `json:"nome_pessoa,omitempty"`
	EmpresaID   *int64     `json:"empresa_id,omitempty"`
	Empresa     *empresa   `json:"empresa,omitempty"`
	DataCriacao *time.Time `json:"data_criacao,omitempty"`
}

// ListaAvaliacaosRes modela uma lista de avaliações
type ListaAvaliacaosRes struct {
	TotalAvaliacoes *int64           `json:"total_avaliacoes,omitempty"`
	Dados           []DadosAvaliacao `json:"dados,omitempty"`
}

type empresa struct {
	ID   *int64  `json:"id,omitempty"`
	Nome *string `json:"Nome,omitempty"`
}

type DadosUltimasAvaliacoesRes struct {
	Dados []UltimasAvaliacoes `json:"dados,omitempty"`
}

type UltimasAvaliacoes struct {
	ID          *string `json:"id,omitempty"`
	Titulo      *string `json:"titulo,omitempty"`
	Descricao   *string `json:"descricao,omitempty"`
	QtdEstrela  *int64  `json:"qtd_estrela,omitempty"`
	NomePessoa  *string `json:"nome_pessoa,omitempty"`
	NomeEmpresa *string `json:"nome_empresa,omitempty"`
	Categoria   *string `json:"categoria,omitempty"`
}
