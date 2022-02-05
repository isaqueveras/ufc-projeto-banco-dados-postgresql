package categorias

import (
	"strconv"

	"github.com/gin-gonic/gin"
	app "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/aplicacao/categorias"
)

func cadastrarCategoria(c *gin.Context) {
	var (
		erro  error
		dados app.DadosCategoria
	)

	if erro = c.ShouldBindJSON(&dados); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel cadastrar a categoria: " + erro.Error()})
		return
	}

	if erro = app.CadastrarCategoria(&dados); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel cadastrar a idade: " + erro.Error()})
		return
	}

	c.JSON(201, gin.H{"messagem": "Categoria cadastrado com sucesso."})
}

func excluirCategoria(c *gin.Context) {
	var (
		erro error
		id   int64
	)

	if id, erro = strconv.ParseInt(c.Param("id"), 10, 64); erro != nil {
		c.JSON(400, gin.H{"messagem": "Informe o id da categoria: " + erro.Error()})
	}

	if erro = app.ExcluirCategoria(&id); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel excluir a categoria: " + erro.Error()})
		return
	}

	c.JSON(204, nil)
}

func editarCategoria(c *gin.Context) {
	var (
		erro  error
		id    int64
		dados app.DadosCategoria
	)

	if erro = c.ShouldBindJSON(&dados); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel converter os dados: " + erro.Error()})
		return
	}

	if id, erro = strconv.ParseInt(c.Param("id"), 10, 64); erro != nil {
		c.JSON(400, gin.H{"messagem": "Informe o id da categoria: " + erro.Error()})
	}

	dados.ID = &id
	if erro = app.EditarCategoria(&dados); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel editar a categoria: " + erro.Error()})
		return
	}

	c.JSON(200, gin.H{
		"messagem": "Categoria editado com sucesso.",
	})
}

func listarCategorias(c *gin.Context) {
	var (
		erro  error
		dados *app.ListaCategoriasRes
	)

	if dados, erro = app.ListarCategorias(); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel listar as categorias: " + erro.Error()})
		return
	}

	c.JSON(200, dados)
}

func listarCategoria(c *gin.Context) {
	var (
		erro  error
		dados *app.DadosCategoria
		id    int64
	)

	if id, erro = strconv.ParseInt(c.Param("id"), 10, 64); erro != nil {
		c.JSON(400, gin.H{"messagem": "Informe o id da categoria: " + erro.Error()})
	}

	if dados, erro = app.ListarCategoria(&id); erro != nil {
		c.JSON(400, gin.H{"messagem": "Não foi possivel listar os dados da categoria: " + erro.Error()})
		return
	}

	c.JSON(200, dados)
}
