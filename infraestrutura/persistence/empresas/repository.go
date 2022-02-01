package empresas

import (
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/infraestrutura/persistence/empresas/postgres"
)

type repositorio struct {
	pg *postgres.PGEmpresas
}

// Repositorio inicializa um repositório
func Repositorio(banco *database.DBTransacao) *repositorio {
	return &repositorio{pg: &postgres.PGEmpresas{DB: banco}}
}
