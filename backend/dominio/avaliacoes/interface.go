package avaliacoes

// IAvaliacoes define uma interface para os metodos de acesso à camada de dados
type IAvaliacoes interface {
	CadastrarAvaliacao(dados *DadosAvaliacao) error
}
