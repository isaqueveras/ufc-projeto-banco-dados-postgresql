package avaliacoes

import "github.com/gin-gonic/gin"

func RouterWithID(r *gin.RouterGroup) {
	r.GET(":empresa_id", func(c *gin.Context) {})    // Listar todas as avaliações de uma empresa
	r.POST(":empresa_id", func(c *gin.Context) {})   // Cadastrar uma avaliação
	r.PUT(":empresa_id", func(c *gin.Context) {})    // Editar uma avaliação
	r.DELETE(":empresa_id", func(c *gin.Context) {}) // Excluir uma avaliação
}
