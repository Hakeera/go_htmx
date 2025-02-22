package controller

import (
	"net/http"

	"github.com/Hakeera/crud/model"
	"github.com/gin-gonic/gin"
)

// Renderizar a lista de usuários na página inicial
func RenderUserList(c *gin.Context) {
	users, err := model.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuários"})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{"Users": users})
}

// Renderiza a página inicial
func RenderIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
