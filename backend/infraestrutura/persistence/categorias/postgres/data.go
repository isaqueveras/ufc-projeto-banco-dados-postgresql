package postgres

import (
	"database/sql"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	dominio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/dominio/categorias"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/utils"
)

type PGCategorias struct {
	DB *database.DBTransacao
}

// CadastrarCategoria cadastra uma categoria no banco de dados
func (pg *PGCategorias) CadastrarCategoria(dados *dominio.DadosCategoria) (erro error) {
	if erro = pg.DB.Builder.
		Insert("public.t_categorias").
		Columns("nome, descricao").
		Values(&dados.Nome, &dados.Descricao).
		Suffix(`RETURNING "id"`).
		Scan(new(int64)); erro != nil {
		return erro
	}

	return
}

// ExcluirCategoria exclui uma categoria no banco de dados
func (pg *PGCategorias) ExcluirCategoria(id *int64) (erro error) {
	var (
		resultado sql.Result
		linhas    int64
	)

	if resultado, erro = pg.DB.Builder.
		Delete("public.t_categorias").
		Where(squirrel.Eq{"id": id}).
		Exec(); erro != nil {
		return errors.New("erro ao excluir a categoria")
	}

	if linhas, erro = resultado.RowsAffected(); erro != nil {
		return errors.New("erro ao excluir a categoria")
	} else if linhas == 0 {
		return errors.New("nenhuma categoria foi encontrada")
	}

	return
}

// EditarCategoria edita uma categoria no banco de dados
func (pg *PGCategorias) EditarCategoria(dados *dominio.DadosCategoria) (erro error) {
	valores := map[string]interface{}{
		"nome":      dados.Nome,
		"descricao": dados.Descricao,
	}

	if erro = pg.DB.Builder.
		Update("public.t_categorias").
		SetMap(valores).
		Where(squirrel.Eq{"id": dados.ID}).
		Suffix(`RETURNING "id"`).
		Scan(new(int64)); erro != nil {
		return errors.New("não foi possivel editar a categoria")
	}

	return
}

// ListarCategorias lista todas as categorias no banco de dados
func (pg *PGCategorias) ListarCategorias() (categorias *dominio.ListaCategorias, erro error) {
	categorias = new(dominio.ListaCategorias)

	consulta := pg.DB.Builder.
		Select("id, nome, descricao, data_criacao").
		From("public.t_categorias")

	linhas, erro := consulta.Query()
	if erro != nil {
		return nil, erro
	}

	for linhas.Next() {
		var categoria dominio.DadosCategoria
		if erro = linhas.Scan(&categoria.ID, &categoria.Nome, &categoria.Descricao, &categoria.DataCriacao); erro != nil {
			return nil, erro
		}
		categorias.Dados = append(categorias.Dados, categoria)
	}

	// Atribuindo o total de categorias
	categorias.TotalCategorias = utils.PonteiroInt64(int64(len(categorias.Dados)))

	return
}

// ListarCategoria lista os dados de uma categoria no banco de dados
func (pg *PGCategorias) ListarCategoria(id *int64) (categoria *dominio.DadosCategoria, erro error) {
	categoria = new(dominio.DadosCategoria)

	if erro = pg.DB.Builder.
		Select("id, nome, descricao, data_criacao").
		From("public.t_categorias").
		Where(squirrel.Eq{"id": id}).
		Scan(&categoria.ID, &categoria.Nome, &categoria.Descricao, &categoria.DataCriacao); erro != nil {
		if sql.ErrNoRows == erro {
			return nil, errors.New("não foi encontrado nenhuma categoria com esse id")
		}

		return nil, erro
	}

	return
}
