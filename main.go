package main
import (
	"fmt"
	"log"
	"net/http"

	"github.com/Hakeera/crud/routes"
)

func main() {
	r := routes.SetupRoutes()

	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

