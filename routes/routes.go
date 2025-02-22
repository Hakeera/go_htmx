package routes

import (
	"github.com/Hakeera/crud/controller"
	"github.com/gin-gonic/gin"
)

// Configurar as rotas
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

	// Rotas do CRUD de usuários
	r.GET("/users", controller.GetUsers)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)

	return r
}

