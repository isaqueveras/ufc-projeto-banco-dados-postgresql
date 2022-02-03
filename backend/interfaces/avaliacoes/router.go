package avaliacoes

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("", cadastrarAvaliacao) // Cadastrar uma avaliação
}

func RouterWithID(r *gin.RouterGroup) {
	r.GET(":empresa_id", func(c *gin.Context) {})    // Listar todas as avaliações de uma empresa
	r.PUT(":empresa_id", func(c *gin.Context) {})    // Editar uma avaliação
	r.DELETE(":empresa_id", func(c *gin.Context) {}) // Excluir uma avaliação
}
