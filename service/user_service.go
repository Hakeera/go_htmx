package service

import (
	"log"

	"github.com/Hakeera/crud/config"
	"github.com/Hakeera/crud/model"
)

// Função para criar um usuário no banco
func CreateUser(user model.User) error {
	db := config.ConnectDB()
	defer db.Close()

	sqlStatement := `INSERT INTO test_crud (name) VALUES ($1) RETURNING id`
	err := db.QueryRow(sqlStatement, user.Name).Scan(&user.ID)
	if err != nil {
		log.Println("Erro ao inserir usuário:", err)
		return err
	}

	log.Println("Usuário inserido com sucesso! ID:", user.ID)
	return nil
}

// Função para listar todos os usuários
func GetUsers() ([]model.User, error) {
	db := config.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM test_crud")
	if err != nil {
		log.Println("Erro ao buscar usuários:", err)
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Println("Erro ao escanear usuário:", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Função para atualizar um usuário
func UpdateUser(id int, name string) error {
	db := config.ConnectDB()
	defer db.Close()

	sqlStatement := `UPDATE test_crud SET name=$1 WHERE id=$2`
	_, err := db.Exec(sqlStatement, name, id)
	if err != nil {
		log.Println("Erro ao atualizar usuário:", err)
		return err
	}

	log.Println("Usuário atualizado com sucesso! ID:", id)
	return nil
}

// Função para deletar um usuário
func DeleteUser(id int) error {
	db := config.ConnectDB()
	defer db.Close()

	sqlStatement := `DELETE FROM test_crud WHERE id=$1`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Println("Erro ao deletar usuário:", err)
		return err
	}

	log.Println("Usuário deletado com sucesso! ID:", id)
	return nil
}

