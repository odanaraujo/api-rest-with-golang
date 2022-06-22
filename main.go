package main

import (
	"fmt"
	"github.com/odanaraujo/api-rest-with-golang/models"
	"github.com/odanaraujo/api-rest-with-golang/routes"
)

func main() {

	models.Personalitys = []models.Personality{
		{1, "Dan", "Teste 1"},
		{2, "Maria", "Teste 2"},
		{3, "Araújo", "Teste 3"},
	}

	fmt.Println("Iniciando o servidor")
	routes.LoadRoutes()
}
