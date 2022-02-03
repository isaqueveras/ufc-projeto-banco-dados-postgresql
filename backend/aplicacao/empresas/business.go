package empresas

import (
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	dominio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/dominio/empresas"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/infraestrutura/persistence/empresas"
)

// CadastrarEmpresa contem a logica de negocio para cadastrar uma empresa
func CadastrarEmpresa(dados *DadosEmpresaReq) (erro error) {
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
