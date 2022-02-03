package avaliacoes

import (
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao/database"
	dominio "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/dominio/avaliacoes"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/infraestrutura/persistence/avaliacoes"
)

// CadastrarAvaliacao contem a logica de negocio para cadastrar uma avaliação
func CadastrarAvaliacao(dados *DadosAvaliacaoReq) (erro error) {
	transacao, erro := database.AbrirTransacao()
	if erro != nil {
		return
	}

	defer transacao.Rollback()

	repo := avaliacoes.Novo(transacao)
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
