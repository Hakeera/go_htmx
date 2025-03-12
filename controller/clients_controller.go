package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Hakeera/crud/service"
	"github.com/gin-gonic/gin"
)

// Função para criar um cliente
func CreateClient(c *gin.Context) {
    var client service.ClientDTO

    fmt.Println("🔵 Recebendo requisição para criar cliente...")

    if err := c.ShouldBind(&client); err != nil {
        fmt.Println("🛑 Erro ao processar a requisição:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao processar a requisição"})
        return
    }

    fmt.Println("✅ Cliente recebido:", client)
    
    _, err := service.CreateClientService(client)
    if err != nil {
        fmt.Println("🛑 Erro ao criar cliente:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar cliente"})
        return
    }

    clients, err := service.GetClientsService()
    if err != nil {
        fmt.Println("🛑 Erro ao buscar clientes:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar clientes"})
        return
    }

    c.HTML(http.StatusOK, "clientes-list.html", gin.H{"clients": clients})
}

// Função para obter todos os clientes
func UpdateClient(c *gin.Context) {
	id := c.Param("id")
	var client service.ClientDTO

	// Bind JSON do corpo da requisição para a struct client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao processar a requisição"})
		return
	}

	// Converter o id para inteiro
	clientID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Atribuir o ID ao cliente
	client.ID = clientID

	// Chamar o service para atualizar o cliente
	updatedClient, err := service.UpdateClientService(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar cliente"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cliente atualizado com sucesso", "client": updatedClient})
}

// Função para deletar um cliente
func DeleteClient(c *gin.Context) {
	id := c.Param("id")

	// Converter o id para inteiro
	clientID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Chamar o service para deletar o cliente
	err = service.DeleteClientService(clientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar cliente"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cliente deletado com sucesso"})
}

// Função para obter todos os clientes
func GetClients(c *gin.Context) {
	clients, err := service.GetClientsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar clientes"})
		return
	}

	// Retorna os clientes como resposta
	c.JSON(http.StatusOK, clients)
}
