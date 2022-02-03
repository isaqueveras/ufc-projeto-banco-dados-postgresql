package postgres

import (
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	dominio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/dominio/avaliacoes"
)

type PGAvaliacoes struct {
	DB *database.DBTransacao
}

// CadastrarAvaliacao cadastra uma avalição no bando de dados
func (pg *PGAvaliacoes) CadastrarAvaliacao(dados *dominio.DadosAvaliacao) (erro error) {
	if erro = pg.DB.Builder.
		Insert("public.t_avaliacoes").
		Columns("titulo, descricao, qtd_estrela, nome_pessoa, empresa_id").
		Values(
			&dados.Titulo,
			&dados.Descricao,
			&dados.QtdEstrela,
			&dados.NomePessoa,
			&dados.EmpresaID).
		Suffix(`RETURNING "id"`).
		Scan(new(int64)); erro != nil {
		return erro
	}

	return
}
