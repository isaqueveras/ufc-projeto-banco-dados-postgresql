package avaliacoes

import (
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	dominio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/dominio/avaliacoes"
	repositorio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/infraestrutura/persistence/avaliacoes"
)

// CadastrarAvaliacao contem a logica de negocio para cadastrar uma avaliação
func CadastrarAvaliacao(dados *DadosAvaliacao) (erro error) {
	transacao, erro := database.AbrirTransacao()
	if erro != nil {
		return
	}

	defer transacao.Rollback()

	repo := repositorio.Novo(transacao)
	if erro = repo.CadastrarAvaliacao(&dominio.DadosAvaliacao{
		Titulo:     dados.Titulo,
		Descricao:  dados.Descricao,
		QtdEstrela: dados.QtdEstrela,
		NomePessoa: dados.NomePessoa,
		EmpresaID:  dados.EmpresaID,
	}); erro != nil {
		return
	}

	if erro = transacao.Commit(); erro != nil {
		return erro
	}

	return nil
}

// EditarAvaliacao contem a logica de negocio para editar uma avaliação
func EditarAvaliacao(dados *DadosAvaliacao) (erro error) {
	transacao, erro := database.AbrirTransacao()
	if erro != nil {
		return
	}

	defer transacao.Rollback()

	repo := repositorio.Novo(transacao)
	if erro = repo.EditarAvaliacao(&dominio.DadosAvaliacao{
		ID:         dados.ID,
		Titulo:     dados.Titulo,
		Descricao:  dados.Descricao,
		QtdEstrela: dados.QtdEstrela,
		NomePessoa: dados.NomePessoa,
		EmpresaID:  dados.EmpresaID,
	}); erro != nil {
		return
	}

	if erro = transacao.Commit(); erro != nil {
		return erro
	}

	return nil
}

// ExcluirAvaliacao contem a logica de negocio para excluir uma avaliação
func ExcluirAvaliacao(id *int64) (erro error) {
	transacao, erro := database.AbrirTransacao()
	if erro != nil {
		return
	}

	defer transacao.Rollback()

	repo := repositorio.Novo(transacao)
	if erro = repo.ExcluirAvaliacao(id); erro != nil {
		return
	}

	if erro = transacao.Commit(); erro != nil {
		return erro
	}

	return nil
}

// ListarAvaliacoes contem a logica de negocio para listar as avaliações
func ListarAvaliacoes() (listaRes *ListaAvaliacaosRes, erro error) {
	listaRes = new(ListaAvaliacaosRes)

	var (
		dados     *dominio.ListaAvaliacaos
		transacao *database.DBTransacao
	)

	if transacao, erro = database.AbrirTransacao(); erro != nil {
		return nil, erro
	}

	defer transacao.Rollback()

	repo := repositorio.Novo(transacao)
	if dados, erro = repo.ListarAvaliacoes(); erro != nil {
		return nil, erro
	}

	for i := range dados.Dados {
		listaRes.Dados = append(listaRes.Dados, DadosAvaliacao{
			ID:         dados.Dados[i].ID,
			Titulo:     dados.Dados[i].Titulo,
			Descricao:  dados.Dados[i].Descricao,
			QtdEstrela: dados.Dados[i].QtdEstrela,
			NomePessoa: dados.Dados[i].NomePessoa,
			Empresa: &empresa{
				ID:   dados.Dados[i].Empresa.ID,
				Nome: dados.Dados[i].Empresa.Nome,
			},
			DataCriacao: dados.Dados[i].DataCriacao,
		})
	}

	// Atribuindo o total de avaliações na lista
	listaRes.TotalAvaliacoes = dados.TotalAvaliacoes

	if erro = transacao.Commit(); erro != nil {
		return nil, erro
	}

	return
}

// ListarAvaliacao contem a logica de negocio para listar os dados de uma avaliação
func ListarAvaliacao(id *int64) (avaliacaoRes *DadosAvaliacao, erro error) {
	avaliacaoRes = new(DadosAvaliacao)

	var (
		dados     *dominio.DadosAvaliacao
		transacao *database.DBTransacao
	)

	if transacao, erro = database.AbrirTransacao(); erro != nil {
		return nil, erro
	}

	defer transacao.Rollback()

	repo := repositorio.Novo(transacao)
	if dados, erro = repo.ListarAvaliacao(id); erro != nil {
		return nil, erro
	}

	avaliacaoRes = &DadosAvaliacao{
		ID:         dados.ID,
		Titulo:     dados.Titulo,
		Descricao:  dados.Descricao,
		QtdEstrela: dados.QtdEstrela,
		NomePessoa: dados.NomePessoa,
		Empresa: &empresa{
			ID:   dados.Empresa.ID,
			Nome: dados.Empresa.Nome,
		},
		DataCriacao: dados.DataCriacao,
	}

	if erro = transacao.Commit(); erro != nil {
		return nil, erro
	}

	return
}
