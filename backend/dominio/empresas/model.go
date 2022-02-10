package empresas

import "time"

// DadosEmpresa modela os dados para cadastrar uma empresa
type DadosEmpresa struct {
	ID          *int64
	Nome        *string
	Telefone    *string
	Cpf         *string
	Cnpj        *string
	CidadeID    *int64
	CategoriaID *int64
	DataCriacao *time.Time
}

// ListaEmpresas modela uma lista de empresas
type ListaEmpresas struct {
	Total *int64
	Dados []DadosEmpresa
}

type DadosEmpresas struct {
	Dados []DadosAvaliacoesEmpresa
}

type DadosAvaliacoesEmpresa struct {
	Nome          *string
	MediaEstrelas *int64
	QtdAvaliacoes *int64
}
