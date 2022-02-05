package avaliacoes

// IAvaliacoes define uma interface para os metodos de acesso à camada de dados
type IAvaliacoes interface {
	CadastrarAvaliacao(dados *DadosAvaliacao) error
	ExcluirAvaliacao(id *int64) error
	EditarAvaliacao(dados *DadosAvaliacao) error
	ListarAvaliacoes() (*ListaAvaliacaos, error)
	ListarAvaliacao(id *int64) (*DadosAvaliacao, error)
}
