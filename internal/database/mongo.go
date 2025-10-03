package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client es la instancia global de MongoDB que puede ser reutilizada en todo el proyecto
var Client *mongo.Client

// ConnectMongo establece y retorna una conexión con MongoDB usando la URI proporcionada.
// Además, hace un ping para asegurarse de que la conexión fue exitosa.
func ConnectMongo(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// Test de conexión hacia la Base de Datos
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("No fue posible establecer conexión hacia MongoDB:", err)
	}

	log.Println("Se realizó conexión exitosa hacia MongoDB")
	Client = client
	return client
}
