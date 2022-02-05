package postgres

import (
	"database/sql"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	dominio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/dominio/avaliacoes"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/utils"
)

type PGAvaliacoes struct {
	DB *database.DBTransacao
}

// CadastrarAvaliacao cadastra uma avaliação no bando de dados
func (pg *PGAvaliacoes) CadastrarAvaliacao(dados *dominio.DadosAvaliacao) (erro error) {
	valores := map[string]interface{}{
		"titulo":      dados.Titulo,
		"descricao":   dados.Descricao,
		"qtd_estrela": dados.QtdEstrela,
		"nome_pessoa": dados.NomePessoa,
		"empresa_id":  dados.EmpresaID,
	}

	if erro = pg.DB.Builder.
		Insert("public.t_avaliacoes").
		SetMap(valores).
		Suffix(`RETURNING "id"`).
		Scan(new(int64)); erro != nil {
		return erro
	}

	return
}

// ExcluirAvaliacao excluir uma avaliação no banco de dados
func (pg *PGAvaliacoes) ExcluirAvaliacao(id *int64) (erro error) {
	var (
		resultado sql.Result
		linhas    int64
	)

	if resultado, erro = pg.DB.Builder.
		Delete("public.t_avaliacoes").
		Where(squirrel.Eq{"id": id}).
		Exec(); erro != nil {
		return errors.New("erro ao excluir a avaliação")
	}

	if linhas, erro = resultado.RowsAffected(); erro != nil {
		return errors.New("erro ao excluir a avaliação")
	} else if linhas == 0 {
		return errors.New("nenhuma avaliação foi encontrada")
	}

	return
}

// EditarAvaliacao edita uma avaliação no banco de dados
func (pg *PGAvaliacoes) EditarAvaliacao(dados *dominio.DadosAvaliacao) (erro error) {
	valores := map[string]interface{}{
		"titulo":      dados.Titulo,
		"descricao":   dados.Descricao,
		"qtd_estrela": dados.QtdEstrela,
		"nome_pessoa": dados.NomePessoa,
		"empresa_id":  dados.EmpresaID,
	}

	if erro = pg.DB.Builder.
		Update("public.t_avaliacoes").
		SetMap(valores).
		Where(squirrel.Eq{"id": dados.ID}).
		Suffix(`RETURNING "id"`).
		Scan(new(int64)); erro != nil {
		return erro
	}

	return
}

// ListarAvaliacoes lista todas as avaliações no banco de dados
func (pg *PGAvaliacoes) ListarAvaliacoes() (avaliacaos *dominio.ListaAvaliacaos, erro error) {
	avaliacaos = new(dominio.ListaAvaliacaos)

	consulta := pg.DB.Builder.
		Select(`
			TA.id, 
			TA.titulo, 
			TA.descricao, 
			TA.qtd_estrela, 
			TA.nome_pessoa, 
			TE.id, 
			TE.nome, 
			TA.data_criacao`,
		).
		From("public.t_avaliacoes TA").
		Join("public.t_empresas TE ON TE.id = TA.empresa_id")

	linhas, erro := consulta.Query()
	if erro != nil {
		return nil, erro
	}

	for linhas.Next() {
		var avaliacao dominio.DadosAvaliacao
		avaliacao.Empresa = new(dominio.Empresa)

		if erro = linhas.Scan(
			&avaliacao.ID,
			&avaliacao.Titulo,
			&avaliacao.Descricao,
			&avaliacao.QtdEstrela,
			&avaliacao.NomePessoa,
			&avaliacao.Empresa.ID,
			&avaliacao.Empresa.Nome,
			&avaliacao.DataCriacao,
		); erro != nil {
			return nil, erro
		}

		avaliacaos.Dados = append(avaliacaos.Dados, avaliacao)
	}

	// Atribuindo o total de avaliações
	avaliacaos.TotalAvaliacoes = utils.PonteiroInt64(int64(len(avaliacaos.Dados)))

	return
}

// ListarAvaliacao lista os dados de uma avaliação no banco de dados
func (pg *PGAvaliacoes) ListarAvaliacao(id *int64) (avaliacao *dominio.DadosAvaliacao, erro error) {
	avaliacao = new(dominio.DadosAvaliacao)
	avaliacao.Empresa = new(dominio.Empresa)

	if erro = pg.DB.Builder.
		Select("TA.id, TA.titulo, TA.descricao, TA.qtd_estrela, TA.nome_pessoa, TE.id, TE.nome, TA.data_criacao").
		From("public.t_avaliacoes TA").
		Join("public.t_empresas TE ON TE.id = TA.empresa_id").
		Where(squirrel.Eq{"TA.id": id}).
		Limit(1).
		Scan(
			&avaliacao.ID,
			&avaliacao.Titulo,
			&avaliacao.Descricao,
			&avaliacao.QtdEstrela,
			&avaliacao.NomePessoa,
			&avaliacao.Empresa.ID,
			&avaliacao.Empresa.Nome,
			&avaliacao.DataCriacao,
		); erro != nil {
		if sql.ErrNoRows == erro {
			return nil, errors.New("não foi encontrado nenhuma avaliação com esse id")
		}

		return nil, erro
	}

	return
}
