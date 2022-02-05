package cidades

import (
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	dominio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/dominio/cidades"
	repositorio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/infraestrutura/persistence/cidades"
)

// CadastrarCidade contem a logica de negocio para cadastrar uma cidade
func CadastrarCidade(dados *DadosCidade) (erro error) {
	transacao, erro := database.AbrirTransacao()
	if erro != nil {
		return
	}

	defer transacao.Rollback()

	repo := repositorio.Novo(transacao)
	if erro = repo.CadastrarCidade(&dominio.DadosCidade{
		Nome: dados.Nome,
		UF:   dados.UF,
	}); erro != nil {
		return
	}

	if erro = transacao.Commit(); erro != nil {
		return erro
	}

	return nil
}

// ExcluirCidade contem a logica de negocio para excluir uma cidade
func ExcluirCidade(id *int64) (erro error) {
	transacao, erro := database.AbrirTransacao()
	if erro != nil {
		return
	}

	defer transacao.Rollback()

	repo := repositorio.Novo(transacao)
	if erro = repo.ExcluirCidade(id); erro != nil {
		return
	}

	if erro = transacao.Commit(); erro != nil {
		return erro
	}

	return nil
}

// EditarCidade contem a logica de negocio para editar uma cidade
func EditarCidade(dados *DadosCidade) (erro error) {
	transacao, erro := database.AbrirTransacao()
	if erro != nil {
		return
	}
	defer transacao.Rollback()

	repo := repositorio.Novo(transacao)
	if erro = repo.EditarCidade(&dominio.DadosCidade{
		ID:   dados.ID,
		Nome: dados.Nome,
		UF:   dados.UF,
	}); erro != nil {
		return
	}

	if erro = transacao.Commit(); erro != nil {
		return erro
	}

	return
}

// ListarCidades contem a logica de negocio para listar as cidades
func ListarCidades() (listaRes *ListaCidadesRes, erro error) {
	listaRes = new(ListaCidadesRes)

	var (
		dados     *dominio.ListaCidades
		transacao *database.DBTransacao
	)

	if transacao, erro = database.AbrirTransacao(); erro != nil {
		return nil, erro
	}

	defer transacao.Rollback()

	repo := repositorio.Novo(transacao)
	if dados, erro = repo.ListarCidades(); erro != nil {
		return nil, erro
	}

	for i := range dados.Dados {
		listaRes.Dados = append(listaRes.Dados, DadosCidade{
			ID:          dados.Dados[i].ID,
			Nome:        dados.Dados[i].Nome,
			UF:          dados.Dados[i].UF,
			DataCriacao: dados.Dados[i].DataCriacao,
		})
	}

	// Atribuindo o total de cidades na lista
	listaRes.TotalCidades = dados.TotalCidades

	if erro = transacao.Commit(); erro != nil {
		return nil, erro
	}

	return
}

// ListarCidade contem a logica de negocio para listar os dados de uma cidade
func ListarCidade(id *int64) (cidadeRes *DadosCidade, erro error) {
	cidadeRes = new(DadosCidade)

	var (
		dados     *dominio.DadosCidade
		transacao *database.DBTransacao
	)

	if transacao, erro = database.AbrirTransacao(); erro != nil {
		return nil, erro
	}

	defer transacao.Rollback()

	repo := repositorio.Novo(transacao)
	if dados, erro = repo.ListarCidade(id); erro != nil {
		return nil, erro
	}

	cidadeRes = &DadosCidade{
		ID:          dados.ID,
		Nome:        dados.Nome,
		UF:          dados.UF,
		DataCriacao: dados.DataCriacao,
	}

	if erro = transacao.Commit(); erro != nil {
		return nil, erro
	}

	return
}
