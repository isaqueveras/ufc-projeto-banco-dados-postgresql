package categorias

import (
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	dominio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/dominio/categorias"
	repositorio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/infraestrutura/persistence/categorias"
)

// CadastrarCategoria contem a logica de negocio para cadastrar uma categoria
func CadastrarCategoria(dados *DadosCategoria) (erro error) {
	transacao, erro := database.AbrirTransacao()
	if erro != nil {
		return
	}

	defer transacao.Rollback()

	repo := repositorio.Novo(transacao)
	if erro = repo.CadastrarCategoria(&dominio.DadosCategoria{
		Nome:      dados.Nome,
		Descricao: dados.Descricao,
	}); erro != nil {
		return
	}

	if erro = transacao.Commit(); erro != nil {
		return erro
	}

	return nil
}

// ExcluirCategoria contem a logica de negocio para excluir uma categoria
func ExcluirCategoria(id *int64) (erro error) {
	transacao, erro := database.AbrirTransacao()
	if erro != nil {
		return
	}

	defer transacao.Rollback()

	repo := repositorio.Novo(transacao)
	if erro = repo.ExcluirCategoria(id); erro != nil {
		return
	}

	if erro = transacao.Commit(); erro != nil {
		return erro
	}

	return nil
}

// EditarCategoria contem a logica de negocio para editar uma categoria
func EditarCategoria(dados *DadosCategoria) (erro error) {
	transacao, erro := database.AbrirTransacao()
	if erro != nil {
		return
	}
	defer transacao.Rollback()

	repo := repositorio.Novo(transacao)
	if erro = repo.EditarCategoria(&dominio.DadosCategoria{
		ID:        dados.ID,
		Nome:      dados.Nome,
		Descricao: dados.Descricao,
	}); erro != nil {
		return
	}

	if erro = transacao.Commit(); erro != nil {
		return erro
	}

	return
}

// ListarCategorias contem a logica de negocio para listar as categorias
func ListarCategorias() (listaRes *ListaCategoriasRes, erro error) {
	listaRes = new(ListaCategoriasRes)

	var (
		dados     *dominio.ListaCategorias
		transacao *database.DBTransacao
	)

	if transacao, erro = database.AbrirTransacao(); erro != nil {
		return nil, erro
	}

	defer transacao.Rollback()

	repo := repositorio.Novo(transacao)
	if dados, erro = repo.ListarCategorias(); erro != nil {
		return nil, erro
	}

	for i := range dados.Dados {
		listaRes.Dados = append(listaRes.Dados, DadosCategoria{
			ID:          dados.Dados[i].ID,
			Nome:        dados.Dados[i].Nome,
			Descricao:   dados.Dados[i].Descricao,
			DataCriacao: dados.Dados[i].DataCriacao,
		})
	}

	// Atribuindo o total de categorias na lista
	listaRes.TotalCategorias = dados.TotalCategorias

	if erro = transacao.Commit(); erro != nil {
		return nil, erro
	}

	return
}

// ListarCategoria contem a logica de negocio para listar os dados de uma categoria
func ListarCategoria(id *int64) (categoriaRes *DadosCategoria, erro error) {
	categoriaRes = new(DadosCategoria)

	var (
		dados     *dominio.DadosCategoria
		transacao *database.DBTransacao
	)

	if transacao, erro = database.AbrirTransacao(); erro != nil {
		return nil, erro
	}

	defer transacao.Rollback()

	repo := repositorio.Novo(transacao)
	if dados, erro = repo.ListarCategoria(id); erro != nil {
		return nil, erro
	}

	categoriaRes = &DadosCategoria{
		ID:          dados.ID,
		Nome:        dados.Nome,
		Descricao:   dados.Descricao,
		DataCriacao: dados.DataCriacao,
	}

	if erro = transacao.Commit(); erro != nil {
		return nil, erro
	}

	return
}
