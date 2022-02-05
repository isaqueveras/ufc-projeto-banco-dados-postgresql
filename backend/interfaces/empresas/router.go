package empresas

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("", listarEmpresas)    // Lista todas as empresas
	r.POST("", cadastrarEmpresa) // Cadastrar uma empresa
}

func RouterWithID(r *gin.RouterGroup) {
	r.GET(":id", listarEmpresa)     // Listar dados de uma empresa
	r.PUT(":id", editarEmpresa)     // Editar uma empresa
	r.DELETE(":id", excluirEmpresa) // Excluir uma empresa
}
