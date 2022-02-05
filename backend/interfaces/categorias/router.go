package categorias

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("", listarCategorias)    // Listar todas as categorias
	r.POST("", cadastrarCategoria) // Cadastrar uma categoria
}

func RouterWithID(r *gin.RouterGroup) {
	r.GET(":id", listarCategoria)     // Listar dados de uma categoria
	r.PUT(":id", editarCategoria)     // Editar uma categoria
	r.DELETE(":id", excluirCategoria) // Excluir uma categoria
}
