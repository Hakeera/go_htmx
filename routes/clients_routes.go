package routes

import (
	"github.com/Hakeera/crud/controller"
	"github.com/gin-gonic/gin"
)

// ClientRoutes configura as rotas relacionadas a clientes
func ClientRoutes(r *gin.Engine) {
	// Rota para criar um cliente
	r.POST("/clients", controller.CreateClient)

	// Rota para listar clientes
	r.GET("/clients", controller.GetClients)

	// Rota para atualizar um cliente
	r.PUT("/clients/:id", controller.UpdateClient)

	// Rota para excluir um cliente
	r.DELETE("/clients/:id", controller.DeleteClient)
}

