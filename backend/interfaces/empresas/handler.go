package empresas

import (
	"github.com/gin-gonic/gin"
	app "github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/aplicacao/empresas"
)

func cadastrarEmpresa(c *gin.Context) {
	var (
		erro  error
		dados app.DadosEmpresaReq
	)

	if erro = c.ShouldBindJSON(&dados); erro != nil {
		c.JSON(201, gin.H{"messagem": "Não foi possivel cadastrar a empresa: " + erro.Error()})
		return
	}

	if erro = app.CadastrarEmpresa(&dados); erro != nil {
		c.JSON(201, gin.H{"messagem": "Não foi possivel cadastrar a empresa: " + erro.Error()})
		return
	}

	c.JSON(201, gin.H{"messagem": "Empresa cadastrado com sucesso."})
}
