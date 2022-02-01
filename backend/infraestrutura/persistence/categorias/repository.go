package categorias

import (
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/infraestrutura/persistence/categorias/postgres"
)

type repositorio struct {
	pg *postgres.PGCategorias
}

// Repositorio inicializa um repositório
func Repositorio(banco *database.DBTransacao) *repositorio {
	return &repositorio{pg: &postgres.PGCategorias{DB: banco}}
}
