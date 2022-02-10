package cidades

import (
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	dominio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/dominio/cidades"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/infraestrutura/persistence/cidades/postgres"
)

type repositorio struct {
	pg *postgres.PGCidades
}

// Novo inicializa um repositório
func Novo(banco *database.DBTransacao) dominio.ICidades {
	return &repositorio{pg: &postgres.PGCidades{DB: banco}}
}

func (r *repositorio) CadastrarCidade(dados *dominio.DadosCidade) error {
	return r.pg.CadastrarCidade(dados)
}

func (r *repositorio) ExcluirCidade(id *int64) error {
	return r.pg.ExcluirCidade(id)
}

func (r *repositorio) EditarCidade(dados *dominio.DadosCidade) error {
	return r.pg.EditarCidade(dados)
}

func (r *repositorio) ListarCidades() (cidades *dominio.ListaCidades, erro error) {
	return r.pg.ListarCidades()
}

func (r *repositorio) ListarCidade(id *int64) (cidade *dominio.DadosCidade, erro error) {
	return r.pg.ListarCidade(id)
}

func (r *repositorio) MelhoresCidades() (*dominio.DadosCidades, error) {
	return r.pg.MelhoresCidades()
}
