package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Hakeera/crud/model"
	"github.com/Hakeera/crud/service"
	"github.com/gin-gonic/gin"
)

// Função para criar um cliente
func CreateClient(c *gin.Context) {
    var client model.Client

    // Decodifica os dados do formulário 
    if err := c.ShouldBind(&client); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
        return
    }

    fmt.Printf("✅ Cliente recebido: %+v\n", client)

    // Criar o cliente no banco de dados
    _, err := service.CreateClientService(service.ClientDTO{
        Name:    client.Name,
        Email:   client.Email,
        Phone:   client.Phone,
        Address: client.Address,
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Obter a lista completa de clientes
    clients, _ := service.GetClientsService()

    // Retornar o HTML atualizado com a lista de clientes
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
func GetClients(c *gin.Context) {
    clients, err := service.GetClientsService()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar clientes"})
        return
    }

    fmt.Println("🚀 Enviando para o template:", clients) // Debug

    // Teste renderizando diretamente um HTML básico
    c.HTML(http.StatusOK, "clientes-list.html", gin.H{
        "clients": clients,
    })
}
