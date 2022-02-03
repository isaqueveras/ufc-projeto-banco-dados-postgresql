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
