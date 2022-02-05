package avaliacoes

import (
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	dominio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/dominio/avaliacoes"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/infraestrutura/persistence/avaliacoes/postgres"
)

type repositorio struct {
	pg *postgres.PGAvaliacoes
}

// Novo inicializa um repositório
func Novo(banco *database.DBTransacao) dominio.IAvaliacoes {
	return &repositorio{pg: &postgres.PGAvaliacoes{DB: banco}}
}

func (r *repositorio) CadastrarAvaliacao(dados *dominio.DadosAvaliacao) error {
	return r.pg.CadastrarAvaliacao(dados)
}

func (r *repositorio) ExcluirAvaliacao(id *int64) error {
	return r.pg.ExcluirAvaliacao(id)
}

func (r *repositorio) EditarAvaliacao(dados *dominio.DadosAvaliacao) error {
	return r.pg.EditarAvaliacao(dados)
}

func (r *repositorio) ListarAvaliacoes() (avaliacoes *dominio.ListaAvaliacaos, erro error) {
	return r.pg.ListarAvaliacoes()
}

func (r *repositorio) ListarAvaliacao(id *int64) (avaliacao *dominio.DadosAvaliacao, erro error) {
	return r.pg.ListarAvaliacao(id)
}
