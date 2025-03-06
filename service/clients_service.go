package service

import (
	"fmt"

	"github.com/Hakeera/crud/model"
)

// Struct DTO para Cliente (Data Transfer Object)
type ClientDTO struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// Função para criar um cliente
func CreateClientService(client ClientDTO) (ClientDTO, error) {
	// Converter o DTO para model.Client
	clientModel := model.Client{
		Name:    client.Name,
		Email:   client.Email,
		Phone:   client.Phone,
		Address: client.Address,
	}

	// Criar o cliente no banco de dados
	createdClient, err := model.CreateClient(clientModel)
	if err != nil {
		return ClientDTO{}, fmt.Errorf("erro ao criar cliente: %w", err)
	}

	// Converter o model.Client para ClientDTO
	clientDTO := ClientDTO{
		ID:      createdClient.ID,
		Name:    createdClient.Name,
		Email:   createdClient.Email,
		Phone:   createdClient.Phone,
		Address: createdClient.Address,
	}

	return clientDTO, nil
}

// Função para obter todos os clientes
func GetClientsService() ([]ClientDTO, error) {
	var clientDTOs []ClientDTO

	// Obter os clientes do banco de dados
	clients, err := model.GetClients()
	if err != nil {
		return nil, fmt.Errorf("erro ao obter clientes: %w", err)
	}

	// Converter de model.Client para ClientDTO
	for _, client := range clients {
		clientDTO := ClientDTO{
			ID:      client.ID,
			Name:    client.Name,
			Email:   client.Email,
			Phone:   client.Phone,
			Address: client.Address,
		}
		clientDTOs = append(clientDTOs, clientDTO)
	}

	return clientDTOs, nil
}

// Função para atualizar um cliente
func UpdateClientService(client ClientDTO) (ClientDTO, error) {
	// Converter o DTO para model.Client
	clientModel := model.Client{
		ID:      client.ID,
		Name:    client.Name,
		Email:   client.Email,
		Phone:   client.Phone,
		Address: client.Address,
	}

	// Atualizar o cliente no banco de dados
	updatedClient, err := model.UpdateClient(clientModel)
	if err != nil {
		return ClientDTO{}, fmt.Errorf("erro ao atualizar cliente: %w", err)
	}

	// Converter o model.Client para ClientDTO
	clientDTO := ClientDTO{
		ID:      updatedClient.ID,
		Name:    updatedClient.Name,
		Email:   updatedClient.Email,
		Phone:   updatedClient.Phone,
		Address: updatedClient.Address,
	}

	return clientDTO, nil
}

// Função para deletar um cliente
func DeleteClientService(clientID int) error {
	// Deletar o cliente do banco de dados
	err := model.DeleteClient(clientID)
	if err != nil {
		return fmt.Errorf("erro ao deletar cliente: %w", err)
	}

	return nil
}

