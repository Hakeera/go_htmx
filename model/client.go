package model

import (
	"fmt"

	"github.com/Hakeera/crud/config"
)

type Client struct {
	ID      int    `json:"id"`
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

// Função para criar um novo cliente no banco
func CreateClient(client Client) (Client, error) {
	db := config.ConnectDB()
	query := `
		INSERT INTO clients (name, email, phone, address)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, email, phone, address
	`

	var newClient Client
	err := db.QueryRow(query, client.Name, client.Email, client.Phone, client.Address).Scan(
		&newClient.ID, &newClient.Name, &newClient.Email, &newClient.Phone, &newClient.Address,
	)
	if err != nil {
		return Client{}, fmt.Errorf("erro ao criar cliente: %v", err)
	}

	return newClient, nil
}

// Função para obter todos os clientes
func GetClients() ([]Client, error) {
	db := config.ConnectDB()

	rows, err := db.Query("SELECT id, name, email, phone, address FROM clients")
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar clientes: %v", err)
	}
	defer rows.Close()

	var clients []Client
	for rows.Next() {
		var client Client
		if err := rows.Scan(&client.ID, &client.Name, &client.Email, &client.Phone, &client.Address); err != nil {
			return nil, fmt.Errorf("erro ao ler dados dos clientes: %v", err)
		}
		clients = append(clients, client)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar sobre resultados: %v", err)
	}

	// Debug para verificar se os clientes foram buscados corretamente
	fmt.Println("📢 Clientes encontrados no banco de dados:", clients)

	return clients, nil
}

// Função para atualizar um cliente no banco de dados
func UpdateClient(client Client) (Client, error) {
	db := config.ConnectDB()
	query := `
		UPDATE clients
		SET name = $1, email = $2, phone = $3, address = $4
		WHERE id = $5
		RETURNING id, name, email, phone, address
	`

	var updatedClient Client
	err := db.QueryRow(query, client.Name, client.Email, client.Phone, client.Address, client.ID).Scan(
		&updatedClient.ID, &updatedClient.Name, &updatedClient.Email, &updatedClient.Phone, &updatedClient.Address,
	)
	if err != nil {
		return Client{}, fmt.Errorf("erro ao atualizar cliente: %v", err)
	}

	return updatedClient, nil
}

// Função para deletar um cliente no banco de dados
func DeleteClient(id int) error {
	db := config.ConnectDB()
	query := `
		DELETE FROM clients
		WHERE id = $1
	`

	// Deletar o cliente
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar cliente: %v", err)
	}

	return nil
}

