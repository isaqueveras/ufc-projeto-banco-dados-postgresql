package empresas

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("", func(c *gin.Context) {})  // Lista todas as empresas
	r.POST("", func(c *gin.Context) {}) // Cadastrar uma empresa
}

func RouterWithID(r *gin.RouterGroup) {
	r.GET(":id", func(c *gin.Context) {})    // Listar dados de uma empresa
	r.PUT(":id", func(c *gin.Context) {})    // Editar uma empresa
	r.DELETE(":id", func(c *gin.Context) {}) // Excluir uma empresa
}
