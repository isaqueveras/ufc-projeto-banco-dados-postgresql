package categorias

// ICategorias define uma interface para os metodos de acesso à camada de dados
type ICategorias interface {
	CadastrarCategoria(dados *DadosCategoria) error
	ExcluirCategoria(id *int64) error
	EditarCategoria(dados *DadosCategoria) error
	ListarCategorias() (*ListaCategorias, error)
	ListarCategoria(id *int64) (*DadosCategoria, error)
}
