package empresas

import "time"

// DadosEmpresaReq modela os dados para cadastrar uma empresa
type DadosEmpresaReq struct {
	ID          *int64     `json:"id,omitempty"`
	Nome        *string    `json:"nome,omitempty"`
	Telefone    *string    `json:"telefone,omitempty"`
	Cpf         *string    `json:"cpf,omitempty"`
	Cnpj        *string    `json:"cnpj,omitempty"`
	CidadeID    *int64     `json:"cidade_id,omitempty"`
	CategoriaID *int64     `json:"categoria_id,omitempty"`
	DataCriacao *time.Time `json:"data_criacao,omitempty"`
}
