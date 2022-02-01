package avaliacoes

import (
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/infraestrutura/persistence/avaliacoes/postgres"
)

type repositorio struct {
	pg *postgres.PGAvaliacoes
}

// Repositorio inicializa um repositório
func Repositorio(banco *database.DBTransacao) *repositorio {
	return &repositorio{pg: &postgres.PGAvaliacoes{DB: banco}}
}
