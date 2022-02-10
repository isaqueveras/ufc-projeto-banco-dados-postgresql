package cidades

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("", listarCidades)    // Listar todas as cidades
	r.POST("", cadastrarCidade) // Cadastrar uma cidade

	r.GET("/melhores", melhoresCidades) // Lista as melhores cidades
}

func RouterWithID(r *gin.RouterGroup) {
	r.GET(":id", listarCidade)     // Listar dados de uma cidade
	r.PUT(":id", editarCidade)     // Editar uma cidade
	r.DELETE(":id", excluirCidade) // Excluir uma cidade
}
