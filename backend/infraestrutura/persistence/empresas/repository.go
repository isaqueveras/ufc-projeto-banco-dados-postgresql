package empresas

import (
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	dominio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/dominio/empresas"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/infraestrutura/persistence/empresas/postgres"
)

type repositorio struct {
	pg *postgres.PGEmpresas
}

// Novo inicializa um repositório
func Novo(banco *database.DBTransacao) dominio.IEmpresas {
	return &repositorio{pg: &postgres.PGEmpresas{DB: banco}}
}

func (r *repositorio) CadastrarEmpresa(dados *dominio.DadosEmpresa) error {
	return r.pg.CadastrarEmpresa(dados)
}

func (r *repositorio) ExcluirEmpresa(id *int64) error {
	return r.pg.ExcluirEmpresa(id)
}

func (r *repositorio) EditarEmpresa(dados *dominio.DadosEmpresa) error {
	return r.pg.EditarEmpresa(dados)
}

func (r *repositorio) ListarEmpresas() (empresas *dominio.ListaEmpresas, erro error) {
	return r.pg.ListarEmpresas()
}

func (r *repositorio) ListarEmpresa(id *int64) (empresa *dominio.DadosEmpresa, erro error) {
	return r.pg.ListarEmpresa(id)
}
