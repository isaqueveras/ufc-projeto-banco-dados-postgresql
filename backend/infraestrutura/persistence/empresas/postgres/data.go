package postgres

import (
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/dominio/empresas"
)

type PGEmpresas struct {
	DB *database.DBTransacao
}

// CadastrarEmpresa cadastra uma empresa no banco de dados
func (pg *PGEmpresas) CadastrarEmpresa(dados *empresas.DadosEmpresa) (erro error) {
	if erro = pg.DB.Builder.
		Insert("public.t_empresas").
		Columns("nome, telefone, cpf, cnpj, cidade_id, categoria_id").
		Values(
			&dados.Nome,
			&dados.Telefone,
			&dados.Cpf,
			&dados.Cnpj,
			&dados.CidadeID,
			&dados.CategoriaID).
		Suffix(`RETURNING "id"`).
		Scan(new(int64)); erro != nil {
		return erro
	}

	return
}
