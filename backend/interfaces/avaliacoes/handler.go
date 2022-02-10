package avaliacoes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	app "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/aplicacao/avaliacoes"
)

func cadastrarAvaliacao(c *gin.Context) {
	var (
		erro  error
		dados app.DadosAvaliacao
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

func excluirAvaliacao(c *gin.Context) {
	var (
		erro error
		id   int64
	)

	if id, erro = strconv.ParseInt(c.Param("id"), 10, 64); erro != nil {
		c.JSON(400, gin.H{"messagem": "Informe o id da avaliação: " + erro.Error()})
	}

	if erro = app.ExcluirAvaliacao(&id); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel excluir a avaliação: " + erro.Error()})
		return
	}

	c.JSON(204, nil)
}

func editarAvaliacao(c *gin.Context) {
	var (
		erro  error
		id    int64
		dados app.DadosAvaliacao
	)

	if erro = c.ShouldBindJSON(&dados); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel converter os dados: " + erro.Error()})
		return
	}

	if id, erro = strconv.ParseInt(c.Param("id"), 10, 64); erro != nil {
		c.JSON(400, gin.H{"messagem": "Informe o id da avaliação: " + erro.Error()})
	}

	dados.ID = &id
	if erro = app.EditarAvaliacao(&dados); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel editar a avaliação: " + erro.Error()})
		return
	}

	c.JSON(200, gin.H{
		"messagem": "Avaliação editado com sucesso.",
	})
}

func listarAvaliacaos(c *gin.Context) {
	var (
		erro  error
		dados *app.ListaAvaliacaosRes
	)

	if dados, erro = app.ListarAvaliacoes(); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel listar as avaliações: " + erro.Error()})
		return
	}

	c.JSON(200, dados)
}

func listarAvaliacao(c *gin.Context) {
	var (
		erro  error
		dados *app.DadosAvaliacao
		id    int64
	)

	if id, erro = strconv.ParseInt(c.Param("id"), 10, 64); erro != nil {
		c.JSON(400, gin.H{"messagem": "Informe o id da avaliação: " + erro.Error()})
	}

	if dados, erro = app.ListarAvaliacao(&id); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel listar os dados da avaliação: " + erro.Error()})
		return
	}

	c.JSON(200, dados)
}

func ultimasAvaliacoes(c *gin.Context) {
	var (
		dados *app.DadosUltimasAvaliacoesRes
		erro  error
	)

	if dados, erro = app.UltimasAvaliacoesInicio(); erro != nil {
		c.JSON(500, gin.H{"messagem": "Não foi possivel listar as ultimas avaliações: " + erro.Error()})
		return
	}

	c.JSON(200, dados)
}
