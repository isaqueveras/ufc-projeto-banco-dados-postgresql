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
	DataCriacao **time.Time
}
