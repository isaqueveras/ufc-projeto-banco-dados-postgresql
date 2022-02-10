package cidades

import (
	"strconv"

	"github.com/gin-gonic/gin"
	app "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/aplicacao/cidades"
)

func cadastrarCidade(c *gin.Context) {
	var (
		erro  error
		dados app.DadosCidade
	)

	if erro = c.ShouldBindJSON(&dados); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel cadastrar a cidade: " + erro.Error()})
		return
	}

	if erro = app.CadastrarCidade(&dados); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel cadastrar a idade: " + erro.Error()})
		return
	}

	c.JSON(201, gin.H{"messagem": "Cidade cadastrado com sucesso."})
}

func excluirCidade(c *gin.Context) {
	var (
		erro error
		id   int64
	)

	if id, erro = strconv.ParseInt(c.Param("id"), 10, 64); erro != nil {
		c.JSON(400, gin.H{"messagem": "Informe o id da cidade: " + erro.Error()})
	}

	if erro = app.ExcluirCidade(&id); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel excluir a cidade: " + erro.Error()})
		return
	}

	c.JSON(204, nil)
}

func editarCidade(c *gin.Context) {
	var (
		erro  error
		id    int64
		dados app.DadosCidade
	)

	if erro = c.ShouldBindJSON(&dados); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel converter os dados: " + erro.Error()})
		return
	}

	if id, erro = strconv.ParseInt(c.Param("id"), 10, 64); erro != nil {
		c.JSON(400, gin.H{"messagem": "Informe o id da cidade: " + erro.Error()})
	}

	dados.ID = &id
	if erro = app.EditarCidade(&dados); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel editar a cidade: " + erro.Error()})
		return
	}

	c.JSON(200, gin.H{
		"messagem": "Cidade editado com sucesso.",
	})
}

func listarCidades(c *gin.Context) {
	var (
		erro  error
		dados *app.ListaCidadesRes
	)

	if dados, erro = app.ListarCidades(); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel listar as cidades: " + erro.Error()})
		return
	}

	c.JSON(200, dados)
}

func listarCidade(c *gin.Context) {
	var (
		erro  error
		dados *app.DadosCidade
		id    int64
	)

	if id, erro = strconv.ParseInt(c.Param("id"), 10, 64); erro != nil {
		c.JSON(400, gin.H{"messagem": "Informe o id da cidade: " + erro.Error()})
	}

	if dados, erro = app.ListarCidade(&id); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel listar os dados da cidade: " + erro.Error()})
		return
	}

	c.JSON(200, dados)
}

func melhoresCidades(c *gin.Context) {
	var (
		dados *app.DadosCidadesRes
		erro  error
	)

	if dados, erro = app.MelhoresCidade(); erro != nil {
		c.JSON(500, gin.H{"messagem": "Não foi possivel listar as melhores empresa: " + erro.Error()})
		return
	}

	c.JSON(200, dados)
}
