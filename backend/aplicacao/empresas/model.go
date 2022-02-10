package empresas

import "time"

// DadosEmpresaReq modela os dados para cadastrar uma empresa
type DadosEmpresa struct {
	ID          *int64     `json:"id,omitempty"`
	Nome        *string    `json:"nome,omitempty"`
	Telefone    *string    `json:"telefone,omitempty"`
	Cpf         *string    `json:"cpf,omitempty"`
	Cnpj        *string    `json:"cnpj,omitempty"`
	CidadeID    *int64     `json:"cidade_id,omitempty"`
	CategoriaID *int64     `json:"categoria_id,omitempty"`
	DataCriacao *time.Time `json:"data_criacao,omitempty"`
}

// ListaEmpresasRes modela uma lista de empresas
type ListaEmpresasRes struct {
	Total *int64         `json:"total_empresas,omitempty"`
	Dados []DadosEmpresa `json:"dados,omitempty"`
}

type DadosEmpresasRes struct {
	Dados []DadosAvaliacoesEmpresa `json:"dados,omitempty"`
}

type DadosAvaliacoesEmpresa struct {
	Nome          *string `json:"nome,omitempty"`
	MediaEstrelas *int64  `json:"media_estrelas,omitempty"`
	QtdAvaliacoes *int64  `json:"qtd_avaliacoes,omitempty"`
}
