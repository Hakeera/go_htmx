package model

import "github.com/Hakeera/crud/config"

// Estrutura do Usuário
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" form:"name"`
}

// Criar usuário no banco
func CreateUser(name string) (*User, error) {
	db := config.ConnectDB() // Mantemos ConnectDB()
	var user User

	err := db.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id, name", name).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Buscar todos os usuários
func GetUsers() ([]User, error) {
	db := config.ConnectDB()
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}

	return users, nil
}

// Atualizar usuário pelo ID
func UpdateUser(id int, name string) error {
	db := config.ConnectDB()
	_, err := db.Exec("UPDATE users SET name = $1 WHERE id = $2", name, id)
	return err
}

// Deletar usuário pelo ID
func DeleteUser(id int) error {
	db := config.ConnectDB()
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

