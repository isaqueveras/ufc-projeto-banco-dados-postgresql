package cidades

import (
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/infraestrutura/persistence/cidades/postgres"
)

type repositorio struct {
	pg *postgres.PGCidades
}

// Repositorio inicializa um repositório
func Repositorio(banco *database.DBTransacao) *repositorio {
	return &repositorio{pg: &postgres.PGCidades{DB: banco}}
}
