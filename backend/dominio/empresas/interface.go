package empresas

// IEmpresas define uma interface para os metodos de acesso à camada de dados
type IEmpresas interface {
	CadastrarEmpresa(dados *DadosEmpresa) error
	ExcluirEmpresa(id *int64) error
	EditarEmpresa(dados *DadosEmpresa) error
	ListarEmpresas() (*ListaEmpresas, error)
	ListarEmpresa(id *int64) (*DadosEmpresa, error)
}
