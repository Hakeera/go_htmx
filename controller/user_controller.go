package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Hakeera/crud/config"
	"github.com/Hakeera/crud/model"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
    db := config.ConnectDB()
    defer db.Close()

    var user model.User

    // Para um formulário tradicional, o Gin vai automaticamente 
    // preencher os dados no struct com ShouldBind
    if err := c.ShouldBind(&user); err != nil {
        fmt.Println("Erro ao vincular os dados:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao criar usuário"})
        return
    }

    // Inserindo o usuário no banco de dados
    query := "INSERT INTO users (name) VALUES ($1) RETURNING id"
    err := db.QueryRow(query, user.Name).Scan(&user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Usuário criado com sucesso!"})
}

func GetUsers(c *gin.Context) {
	db := config.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		fmt.Println("Erro ao buscar usuários:", err) // LOG DETALHADO
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuários"})
		return
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			fmt.Println("Erro ao ler os dados do usuário:", err) // LOG DETALHADO
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar usuários"})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}


// Atualizar usuário
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var request struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	err = model.UpdateUser(id, request.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário atualizado com sucesso!"})
}

// Deletar usuário
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = model.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar usuário"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso!"})
}

