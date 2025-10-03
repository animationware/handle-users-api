package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/animationware/handle-users-api/internal/database"
	"github.com/animationware/handle-users-api/internal/user"
	"github.com/gin-gonic/gin"
)

func main() {
	// Configuración de la conexión hacia la base de datos MongoDB:
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}
	// Conexión a Mongo:
	client := database.ConnectMongo(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer client.Disconnect(ctx)

	db := client.Database("handle-users-db")

	// Inyección de dependencias en orden
	repo := user.NewRepository(db)
	service := user.NewService(repo)    // <- capa de negocio
	handler := user.NewHandler(service) // <- handler consume service

	// Inicializar router
	r := gin.Default()
	handler.RegisterRoutes(r)

	// Puerto configurable
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("API RESTFull ejecutandose en http://localhost:%s", port)
	r.Run(":" + port)
}
