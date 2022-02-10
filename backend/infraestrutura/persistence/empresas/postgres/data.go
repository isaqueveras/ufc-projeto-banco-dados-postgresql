package postgres

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	dominio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/dominio/empresas"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/utils"
)

type PGEmpresas struct {
	DB *database.DBTransacao
}

// CadastrarEmpresa cadastra uma empresa no banco de dados
func (pg *PGEmpresas) CadastrarEmpresa(dados *dominio.DadosEmpresa) (erro error) {
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

// ExcluirEmpresa exclui uma empresa no banco de dados
func (pg *PGEmpresas) ExcluirEmpresa(id *int64) (erro error) {
	var (
		resultado sql.Result
		linhas    int64
	)

	if resultado, erro = pg.DB.Builder.
		Delete("public.t_empresas").
		Where(squirrel.Eq{"id": id}).
		Exec(); erro != nil {
		return errors.New("erro ao excluir empresa")
	}

	if linhas, erro = resultado.RowsAffected(); erro != nil {
		return errors.New("erro ao excluir empresa")
	} else if linhas == 0 {
		return errors.New("nenhuma empresa foi encontrada")
	}

	return
}

// EditarEmpresa edita uma empresa no banco de dados
func (pg *PGEmpresas) EditarEmpresa(dados *dominio.DadosEmpresa) (erro error) {
	valores := map[string]interface{}{
		"nome":         dados.Nome,
		"telefone":     dados.Telefone,
		"cpf":          dados.Cpf,
		"cnpj":         dados.Cnpj,
		"cidade_id":    dados.CidadeID,
		"categoria_id": dados.CategoriaID,
	}

	if erro = pg.DB.Builder.
		Update("public.t_empresas").
		SetMap(valores).
		Where(squirrel.Eq{"id": dados.ID}).
		Suffix(`RETURNING "id"`).
		Scan(new(int64)); erro != nil {
		return erro
	}

	return
}

// ListarEmpresas lista todas as empresas no banco de dados
func (pg *PGEmpresas) ListarEmpresas() (empresas *dominio.ListaEmpresas, erro error) {
	empresas = new(dominio.ListaEmpresas)

	consulta := pg.DB.Builder.
		Select("id, nome, telefone, cpf, cnpj, cidade_id, categoria_id, data_criacao").
		From("public.t_empresas")

	linhas, erro := consulta.Query()
	if erro != nil {
		return nil, erro
	}

	for linhas.Next() {
		var empresa dominio.DadosEmpresa

		if erro = linhas.Scan(
			&empresa.ID,
			&empresa.Nome,
			&empresa.Telefone,
			&empresa.Cpf,
			&empresa.Cnpj,
			&empresa.CidadeID,
			&empresa.CategoriaID,
			&empresa.DataCriacao,
		); erro != nil {
			return nil, erro
		}

		empresas.Dados = append(empresas.Dados, empresa)
	}

	// Atribuindo o total de empresas
	empresas.Total = utils.PonteiroInt64(int64(len(empresas.Dados)))

	return
}

func (pg *PGEmpresas) ListarEmpresa(id *int64) (empresa *dominio.DadosEmpresa, erro error) {
	empresa = new(dominio.DadosEmpresa)

	if erro = pg.DB.Builder.
		Select("id, nome, telefone, cpf, cnpj, cidade_id, categoria_id, data_criacao").
		From("public.t_empresas").
		Where(squirrel.Eq{
			"id": id,
		}).
		Scan(
			&empresa.ID,
			&empresa.Nome,
			&empresa.Telefone,
			&empresa.Cpf,
			&empresa.Cnpj,
			&empresa.CidadeID,
			&empresa.CategoriaID,
			&empresa.DataCriacao,
		); erro != nil {
		if sql.ErrNoRows == erro {
			return nil, errors.New("não foi encontrado nenhuma empresa com esse id")
		}
		return nil, erro
	}

	return
}

func (pg *PGEmpresas) MelhoresEmpresas() (empresas *dominio.DadosEmpresas, erro error) {
	empresas = new(dominio.DadosEmpresas)

	consulta := pg.DB.Builder.
		Select(`
			nome, 
			media_estrelas::INTEGER,
    	qtd_avaliacoes::INTEGER`).
		From("melhores_empresas")

	linhas, erro := consulta.Query()
	if erro != nil {
		return nil, erro
	}

	for linhas.Next() {
		var empresa dominio.DadosAvaliacoesEmpresa
		if erro = linhas.Scan(&empresa.Nome, &empresa.MediaEstrelas, &empresa.QtdAvaliacoes); erro != nil {
			return nil, erro
		}

		log.Println(empresa.Nome)
		empresas.Dados = append(empresas.Dados, empresa)
	}

	return
}

func (pg *PGEmpresas) PioresEmpresas() (empresas *dominio.DadosEmpresas, erro error) {
	empresas = new(dominio.DadosEmpresas)

	consulta := pg.DB.Builder.
		Select(`
			nome, 
			media_estrelas::INTEGER,
    	qtd_avaliacoes::INTEGER`).
		From("piores_empresas")

	linhas, erro := consulta.Query()
	if erro != nil {
		return nil, erro
	}

	for linhas.Next() {
		var empresa dominio.DadosAvaliacoesEmpresa
		if erro = linhas.Scan(&empresa.Nome, &empresa.MediaEstrelas, &empresa.QtdAvaliacoes); erro != nil {
			return nil, erro
		}

		empresas.Dados = append(empresas.Dados, empresa)
	}

	return
}
