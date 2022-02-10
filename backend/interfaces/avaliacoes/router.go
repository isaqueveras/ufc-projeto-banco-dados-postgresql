package avaliacoes

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("", listarAvaliacaos)    // Listar todas as avaliações
	r.POST("", cadastrarAvaliacao) // Cadastrar uma avaliação

	r.GET("/ultimas", ultimasAvaliacoes) // Listar as ultimas avaliações
}

func RouterWithID(r *gin.RouterGroup) {
	r.GET(":id", listarAvaliacao)     // Listar todas as avaliações de uma empresa
	r.PUT(":id", editarAvaliacao)     // Editar uma avaliação
	r.DELETE(":id", excluirAvaliacao) // Excluir uma avaliação
}
