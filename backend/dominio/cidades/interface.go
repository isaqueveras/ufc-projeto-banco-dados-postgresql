package cidades

// ICidades define uma interface para os metodos de acesso à camada de dados
type ICidades interface {
	CadastrarCidade(dados *DadosCidade) error
	ExcluirCidade(id *int64) error
	EditarCidade(dados *DadosCidade) error
	ListarCidades() (*ListaCidades, error)
	ListarCidade(id *int64) (*DadosCidade, error)

	MelhoresCidades() (*DadosCidades, error)
}
