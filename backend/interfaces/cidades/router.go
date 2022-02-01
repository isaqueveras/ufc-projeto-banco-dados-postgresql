package cidades

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("", func(c *gin.Context) {})  // Listar todas as cidades
	r.POST("", func(c *gin.Context) {}) // Cadastrar uma cidade
}

func RouterWithID(r *gin.RouterGroup) {
	r.GET(":id", func(c *gin.Context) {})    // Listar dados de uma cidade
	r.PUT(":id", func(c *gin.Context) {})    // Editar uma cidade
	r.DELETE(":id", func(c *gin.Context) {}) // Excluir uma cidade
}
