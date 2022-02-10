package postgres

import (
	"database/sql"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	dominio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/dominio/cidades"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/utils"
)

type PGCidades struct {
	DB *database.DBTransacao
}

// CadastrarCidade cadastra uma cidade no banco de dados
func (pg *PGCidades) CadastrarCidade(dados *dominio.DadosCidade) (erro error) {
	if erro = pg.DB.Builder.
		Insert("public.t_cidades").
		Columns("nome, uf").
		Values(&dados.Nome, &dados.UF).
		Suffix(`RETURNING "id"`).
		Scan(new(int64)); erro != nil {
		return erro
	}

	return
}

// ExcluirCidade exclui uma cidade no banco de dados
func (pg *PGCidades) ExcluirCidade(id *int64) (erro error) {
	var (
		resultado sql.Result
		linhas    int64
	)

	if resultado, erro = pg.DB.Builder.
		Delete("public.t_cidades").
		Where(squirrel.Eq{"id": id}).
		Exec(); erro != nil {
		return errors.New("erro ao excluir a cidade")
	}

	if linhas, erro = resultado.RowsAffected(); erro != nil {
		return errors.New("erro ao excluir a cidade")
	} else if linhas == 0 {
		return errors.New("nenhuma cidade foi encontrada")
	}

	return
}

// EditarCidade edita uma cidade no banco de dados
func (pg *PGCidades) EditarCidade(dados *dominio.DadosCidade) (erro error) {
	valores := map[string]interface{}{
		"nome": dados.Nome,
		"uf":   dados.UF,
	}

	if erro = pg.DB.Builder.
		Update("public.t_cidades").
		SetMap(valores).
		Where(squirrel.Eq{"id": dados.ID}).
		Suffix(`RETURNING "id"`).
		Scan(new(int64)); erro != nil {
		return erro
	}

	return
}

// ListarCidades lista todas as cidades no banco de dados
func (pg *PGCidades) ListarCidades() (cidades *dominio.ListaCidades, erro error) {
	cidades = new(dominio.ListaCidades)

	consulta := pg.DB.Builder.
		Select("id, nome, uf, data_criacao").
		From("public.t_cidades")

	linhas, erro := consulta.Query()
	if erro != nil {
		return nil, erro
	}

	for linhas.Next() {
		var empresa dominio.DadosCidade
		if erro = linhas.Scan(&empresa.ID, &empresa.Nome, &empresa.UF, &empresa.DataCriacao); erro != nil {
			return nil, erro
		}
		cidades.Dados = append(cidades.Dados, empresa)
	}

	// Atribuindo o total de cidades
	cidades.TotalCidades = utils.PonteiroInt64(int64(len(cidades.Dados)))

	return
}

// ListarCidade lista os dados de uma cidade no banco de dados
func (pg *PGCidades) ListarCidade(id *int64) (cidade *dominio.DadosCidade, erro error) {
	cidade = new(dominio.DadosCidade)

	if erro = pg.DB.Builder.
		Select("id, nome, uf, data_criacao").
		From("public.t_cidades").
		Where(squirrel.Eq{"id": id}).
		Scan(&cidade.ID, &cidade.Nome, &cidade.UF, &cidade.DataCriacao); erro != nil {
		if sql.ErrNoRows == erro {
			return nil, errors.New("não foi encontrado nenhuma cidade com esse id")
		}

		return nil, erro
	}

	return
}

func (pg *PGCidades) MelhoresCidades() (empresas *dominio.DadosCidades, erro error) {
	empresas = new(dominio.DadosCidades)

	consulta := pg.DB.Builder.
		Select(`
			nome, 
			media_estrelas::INTEGER,
    	qtd_avaliacoes::INTEGER`).
		From("melhores_cidades")

	linhas, erro := consulta.Query()
	if erro != nil {
		return nil, erro
	}

	for linhas.Next() {
		var empresa dominio.DadosAvaliacoesCidade
		if erro = linhas.Scan(&empresa.Nome, &empresa.MediaEstrelas, &empresa.QtdAvaliacoes); erro != nil {
			return nil, erro
		}

		empresas.Dados = append(empresas.Dados, empresa)
	}

	return
}
