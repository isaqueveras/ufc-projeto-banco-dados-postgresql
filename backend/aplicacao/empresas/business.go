package empresas

import (
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	dominio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/dominio/empresas"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/infraestrutura/persistence/empresas"
)

// CadastrarEmpresa contem a logica de negocio para cadastrar uma empresa
func CadastrarEmpresa(dados *DadosEmpresa) (erro error) {
	transacao, erro := database.AbrirTransacao()
	if erro != nil {
		return
	}

	defer transacao.Rollback()

	repo := empresas.Novo(transacao)
	if erro = repo.CadastrarEmpresa(&dominio.DadosEmpresa{
		Nome:        dados.Nome,
		Telefone:    dados.Telefone,
		Cpf:         dados.Cpf,
		Cnpj:        dados.Cnpj,
		CidadeID:    dados.CidadeID,
		CategoriaID: dados.CategoriaID,
	}); erro != nil {
		return
	}

	if erro = transacao.Commit(); erro != nil {
		return erro
	}

	return nil
}

// ExcluirEmpresa contem a logica de negocio para excluir uma empresa
func ExcluirEmpresa(id *int64) (erro error) {
	transacao, erro := database.AbrirTransacao()
	if erro != nil {
		return
	}

	defer transacao.Rollback()

	repo := empresas.Novo(transacao)
	if erro = repo.ExcluirEmpresa(id); erro != nil {
		return
	}

	if erro = transacao.Commit(); erro != nil {
		return erro
	}

	return nil
}

// EditarEmpresa contem a logica de negocio para editar uma empresa
func EditarEmpresa(dados *DadosEmpresa) (erro error) {
	transacao, erro := database.AbrirTransacao()
	if erro != nil {
		return
	}

	defer transacao.Rollback()

	repo := empresas.Novo(transacao)
	if erro = repo.EditarEmpresa(&dominio.DadosEmpresa{
		ID:          dados.ID,
		Nome:        dados.Nome,
		Telefone:    dados.Telefone,
		Cpf:         dados.Cpf,
		Cnpj:        dados.Cnpj,
		CidadeID:    dados.CidadeID,
		CategoriaID: dados.CategoriaID,
	}); erro != nil {
		return
	}

	if erro = transacao.Commit(); erro != nil {
		return erro
	}

	return
}

// ListarEmpresas contem a logica de negocio para listar as empresas
func ListarEmpresas() (listaRes *ListaEmpresasRes, erro error) {
	listaRes = new(ListaEmpresasRes)

	var (
		dados     *dominio.ListaEmpresas
		transacao *database.DBTransacao
	)

	if transacao, erro = database.AbrirTransacao(); erro != nil {
		return nil, erro
	}

	defer transacao.Rollback()

	repo := empresas.Novo(transacao)
	if dados, erro = repo.ListarEmpresas(); erro != nil {
		return nil, erro
	}

	for i := range dados.Dados {
		listaRes.Dados = append(listaRes.Dados, DadosEmpresa{
			ID:          dados.Dados[i].ID,
			Nome:        dados.Dados[i].Nome,
			Telefone:    dados.Dados[i].Telefone,
			Cpf:         dados.Dados[i].Cpf,
			Cnpj:        dados.Dados[i].Cnpj,
			CidadeID:    dados.Dados[i].CidadeID,
			CategoriaID: dados.Dados[i].CategoriaID,
			DataCriacao: dados.Dados[i].DataCriacao,
		})
	}

	// Atribuindo o total de empresas na lista
	listaRes.Total = dados.Total

	if erro = transacao.Commit(); erro != nil {
		return nil, erro
	}

	return
}

// ListarEmpresa contem a logica de negocio para listar os dados de uma empresa
func ListarEmpresa(id *int64) (empresaRes *DadosEmpresa, erro error) {
	empresaRes = new(DadosEmpresa)

	var (
		dados     *dominio.DadosEmpresa
		transacao *database.DBTransacao
	)

	if transacao, erro = database.AbrirTransacao(); erro != nil {
		return nil, erro
	}

	defer transacao.Rollback()

	repo := empresas.Novo(transacao)
	if dados, erro = repo.ListarEmpresa(id); erro != nil {
		return nil, erro
	}

	empresaRes = &DadosEmpresa{
		ID:          dados.ID,
		Nome:        dados.Nome,
		Telefone:    dados.Telefone,
		Cpf:         dados.Cpf,
		Cnpj:        dados.Cnpj,
		CidadeID:    dados.CidadeID,
		CategoriaID: dados.CategoriaID,
		DataCriacao: dados.DataCriacao,
	}

	if erro = transacao.Commit(); erro != nil {
		return nil, erro
	}

	return
}

// MelhoresEmpresas contem a logica de negocio para listar as melhores empresas
func MelhoresEmpresas() (empresasRes *DadosEmpresasRes, erro error) {
	empresasRes = new(DadosEmpresasRes)

	var (
		dados     *dominio.DadosEmpresas
		transacao *database.DBTransacao
	)

	if transacao, erro = database.AbrirTransacao(); erro != nil {
		return nil, erro
	}

	defer transacao.Rollback()

	repo := empresas.Novo(transacao)
	if dados, erro = repo.MelhoresEmpresas(); erro != nil {
		return nil, erro
	}

	for i := range dados.Dados {
		empresasRes.Dados = append(empresasRes.Dados, DadosAvaliacoesEmpresa{
			Nome:          dados.Dados[i].Nome,
			MediaEstrelas: dados.Dados[i].MediaEstrelas,
			QtdAvaliacoes: dados.Dados[i].QtdAvaliacoes,
		})
	}

	return
}

// PioresEmpresas contem a logica de negocio para listar as piores empresas
func PioresEmpresas() (empresasRes *DadosEmpresasRes, erro error) {
	empresasRes = new(DadosEmpresasRes)

	var (
		dados     *dominio.DadosEmpresas
		transacao *database.DBTransacao
	)

	if transacao, erro = database.AbrirTransacao(); erro != nil {
		return nil, erro
	}

	defer transacao.Rollback()

	repo := empresas.Novo(transacao)
	if dados, erro = repo.PioresEmpresas(); erro != nil {
		return nil, erro
	}

	for i := range dados.Dados {
		empresasRes.Dados = append(empresasRes.Dados, DadosAvaliacoesEmpresa{
			Nome:          dados.Dados[i].Nome,
			MediaEstrelas: dados.Dados[i].MediaEstrelas,
			QtdAvaliacoes: dados.Dados[i].QtdAvaliacoes,
		})
	}

	return
}
