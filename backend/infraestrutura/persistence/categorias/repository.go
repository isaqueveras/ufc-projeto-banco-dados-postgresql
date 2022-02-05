package categorias

import (
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	dominio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/dominio/categorias"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/infraestrutura/persistence/categorias/postgres"
)

type repositorio struct {
	pg *postgres.PGCategorias
}

// Novo inicializa um repositório
func Novo(banco *database.DBTransacao) dominio.ICategorias {
	return &repositorio{pg: &postgres.PGCategorias{DB: banco}}
}

func (r *repositorio) CadastrarCategoria(dados *dominio.DadosCategoria) error {
	return r.pg.CadastrarCategoria(dados)
}

func (r *repositorio) ExcluirCategoria(id *int64) error {
	return r.pg.ExcluirCategoria(id)
}

func (r *repositorio) EditarCategoria(dados *dominio.DadosCategoria) error {
	return r.pg.EditarCategoria(dados)
}

func (r *repositorio) ListarCategorias() (categorias *dominio.ListaCategorias, erro error) {
	return r.pg.ListarCategorias()
}

func (r *repositorio) ListarCategoria(id *int64) (categoria *dominio.DadosCategoria, erro error) {
	return r.pg.ListarCategoria(id)
}
