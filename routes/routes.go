package routes

import (
	"github.com/Hakeera/crud/controller"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas as rotas do servidor
func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Servir arquivos estáticos (se necessário)
	r.Static("/static", "./static")

	// Carregar templates HTML
	r.LoadHTMLGlob("view/*.html")

	// Rota para a página inicial
	r.GET("/", func(c *gin.Context) {
		controller.RenderIndex(c)
	})

	// Configura rotas específicas para os recursos (usuários, clientes)
	UserRoutes(r)
	ClientRoutes(r)

	return r
}

