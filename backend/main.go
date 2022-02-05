package main

import (
	"log"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/configuracao"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/interfaces/avaliacoes"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/interfaces/categorias"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/interfaces/cidades"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/interfaces/empresas"
	"github.com/isaqueveras/ufc-projeto-banco-dados-postgresql/backend/middleware"
	_ "github.com/lib/pq"
	"golang.org/x/sync/errgroup"
)

func main() {
	// Carregando as configurações do sistema
	configuracao.Carregar()

	// Inicializando o gin
	rotas := gin.Default()
	// gin.SetMode(gin.)

	// Middlewares para todas as rotas
	rotas.Use(middleware.CORS())

	// Grupo para as rotas da primeira versão
	v1 := rotas.Group("v1")

	empresas.Router(v1.Group("empresas"))
	empresas.RouterWithID(v1.Group("empresa"))

	cidades.Router(v1.Group("cidades"))
	cidades.RouterWithID(v1.Group("cidade"))

	categorias.Router(v1.Group("categorias"))
	categorias.RouterWithID(v1.Group("categoria"))

	avaliacoes.Router(v1.Group("avaliacoes"))
	avaliacoes.RouterWithID(v1.Group("avaliacao"))

	grupoErro := errgroup.Group{}
	grupoErro.Go(func() error {
		return endless.ListenAndServe(configuracao.Obter().PortaServidor, rotas)
	})

	// Inicializando o servidor
	if erro := grupoErro.Wait(); erro != nil {
		log.Fatal(erro)
		return
	}
}
