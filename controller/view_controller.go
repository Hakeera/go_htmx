package controller

import (
	"net/http"

	"github.com/Hakeera/crud/service"
	"github.com/gin-gonic/gin"
)

func RenderClientsPage(c *gin.Context) {
    clients, err := service.GetClientsService()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar clientes"})
        return
    }

    // Aqui, você garante que a lista de clientes será renderizada na página
    c.HTML(http.StatusOK, "clientes.html", gin.H{
        "clients": clients,
    })
}


// Renderiza a página inicial
func RenderIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
