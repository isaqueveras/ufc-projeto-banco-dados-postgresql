package empresas

import (
	"strconv"

	"github.com/gin-gonic/gin"
	app "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/aplicacao/empresas"
)

func cadastrarEmpresa(c *gin.Context) {
	var (
		erro  error
		dados app.DadosEmpresa
	)

	if erro = c.ShouldBindJSON(&dados); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel cadastrar a empresa: " + erro.Error()})
		return
	}

	if erro = app.CadastrarEmpresa(&dados); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel cadastrar a empresa: " + erro.Error()})
		return
	}

	c.JSON(201, gin.H{"messagem": "Empresa cadastrado com sucesso."})
}

func excluirEmpresa(c *gin.Context) {
	var (
		erro error
		id   int64
	)

	if id, erro = strconv.ParseInt(c.Param("id"), 10, 64); erro != nil {
		c.JSON(400, gin.H{"messagem": "Informe o id da empresa: " + erro.Error()})
	}

	if erro = app.ExcluirEmpresa(&id); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel excluir a empresa: " + erro.Error()})
		return
	}

	c.JSON(204, nil)
}

func editarEmpresa(c *gin.Context) {
	var (
		erro  error
		id    int64
		dados app.DadosEmpresa
	)

	if erro = c.ShouldBindJSON(&dados); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel converter os dados: " + erro.Error()})
		return
	}

	if id, erro = strconv.ParseInt(c.Param("id"), 10, 64); erro != nil {
		c.JSON(400, gin.H{"messagem": "Informe o id da empresa: " + erro.Error()})
	}

	dados.ID = &id
	if erro = app.EditarEmpresa(&dados); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel cadastrar a empresa: " + erro.Error()})
		return
	}

	c.JSON(200, gin.H{
		"messagem": "Empresa editado com sucesso.",
	})
}

func listarEmpresas(c *gin.Context) {
	var (
		erro  error
		dados *app.ListaEmpresasRes
	)

	if dados, erro = app.ListarEmpresas(); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel listar as empresas: " + erro.Error()})
		return
	}

	c.JSON(200, dados)
}

func listarEmpresa(c *gin.Context) {
	var (
		erro  error
		dados *app.DadosEmpresa
		id    int64
	)

	if id, erro = strconv.ParseInt(c.Param("id"), 10, 64); erro != nil {
		c.JSON(400, gin.H{"messagem": "Informe o id da empresa: " + erro.Error()})
	}

	if dados, erro = app.ListarEmpresa(&id); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel listar os dados da empresa: " + erro.Error()})
		return
	}

	c.JSON(200, dados)
}
