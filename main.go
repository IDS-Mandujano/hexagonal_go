package main

import (
	"log"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
		equipamentInfra "gym-system/src/inventory/infraestructure/routes/Equipment"
)


func main () {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error al cargar el archivo.env: %v", err)
	}

	r := gin.Default()

	equipamentInfra.SetupRoutesEquipament(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v",err)
	}
}