package avaliacoes

import (
	"github.com/gin-gonic/gin"
	app "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/aplicacao/avaliacoes"
)

func cadastrarAvaliacao(c *gin.Context) {
	var (
		erro  error
		dados app.DadosAvaliacaoReq
	)

	if erro = c.ShouldBindJSON(&dados); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel cadastrar a avaliação: " + erro.Error()})
		return
	}

	if erro = app.CadastrarAvaliacao(&dados); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel cadastrar a avaliação: " + erro.Error()})
		return
	}

	c.JSON(201, gin.H{"messagem": "Avaliação cadastrado com sucesso."})
}
