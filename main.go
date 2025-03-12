package main

import (
	"fmt"
	"log"

	"github.com/Hakeera/crud/routes"
)

func main() {
	// Configura as rotas
	r := routes.SetupRoutes()

	// Serve arquivos estáticos a partir do diretório "view"
	r.Static("/assets", "./view")

	// Inicia o servidor na porta 8080
	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(r.Run(":8080")) // Utiliza r.Run para iniciar o servidor
}

