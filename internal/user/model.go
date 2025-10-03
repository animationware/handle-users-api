package user

import "go.mongodb.org/mongo-driver/bson/primitive"

// User representa la entidad de usuario de la aplicación.
// Contiene información básica como nombre y correo, junto con campos de auditoría.
type User struct {
	// ID único generado por MongoDB
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	// Nombre del usuario
	Name string `json:"name" bson:"name"`
	// Correo electrónico del usuario
	Email string `json:"email" bson:"email"`
	// Timestamps de creación y actualización
	CreatedAt int64 `json:"createdAt" bson:"createdAt"`
	// Timestamp de la última actualización
	UpdatedAt int64 `json:"updatedAt" bson:"updatedAt"`
}
