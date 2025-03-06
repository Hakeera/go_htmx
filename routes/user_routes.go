package routes

import (
	"github.com/Hakeera/crud/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
    // Rota para criar um usuário
    r.POST("/users", controller.CreateUser)

    // Rota para listar usuários
    r.GET("/users", controller.GetUsers)

    // Rota para atualizar um usuário
    r.PUT("/users/:id", controller.UpdateUser)

    // Rota para excluir um usuário
    r.DELETE("/users/:id", controller.DeleteUser)
}

