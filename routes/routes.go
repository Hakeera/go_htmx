package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas as rotas do servidor
func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Configura rotas específicas para os recursos (usuários, clientes)
	UserRoutes(r)
	ClientRoutes(r)

	return r
}

