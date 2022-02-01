package categorias

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("", func(c *gin.Context) {})  // Listar todas as categorias
	r.POST("", func(c *gin.Context) {}) // Cadastrar uma categoria
}

func RouterWithID(r *gin.RouterGroup) {
	r.GET(":id", func(c *gin.Context) {})    // Listar dados de uma categoria
	r.PUT(":id", func(c *gin.Context) {})    // Editar uma categoria
	r.DELETE(":id", func(c *gin.Context) {}) // Excluir uma categoria
}
