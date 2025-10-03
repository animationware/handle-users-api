package user

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository maneja la persistencia de datos de usuarios en MongoDB.
// Se encarga de todas las operaciones CRUD sobre la colección "users".
type Repository struct {
	collection *mongo.Collection
}

// NewRepository crea un nuevo Repository conectado a la colección "users".
func NewRepository(db *mongo.Database) *Repository {
	return &Repository{collection: db.Collection("users")}
}

// Create inserta un nuevo usuario en la colección.
// Se registran timestamps de creación y actualización.
func (r *Repository) Create(ctx context.Context, user User) (*mongo.InsertOneResult, error) {
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()
	return r.collection.InsertOne(ctx, user)
}

// FindAll obtiene todos los usuarios almacenados en la colección.
func (r *Repository) FindAll(ctx context.Context) ([]User, error) {
	cur, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var users []User
	for cur.Next(ctx) {
		var u User
		if err := cur.Decode(&u); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// FindByID obtiene un usuario por su ID único.
func (r *Repository) FindByID(ctx context.Context, id string) (User, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	var user User
	err := r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user)
	return user, err
}

// Update actualiza un usuario existente por su ID y devuelve el documento actualizado.
func (r *Repository) Update(ctx context.Context, id string, u User) (User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return User{}, err
	}

	// Usamos FindOneAndUpdate para devolver el documento actualizado
	var updated User
	err = r.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": u},
	).Decode(&updated)
	if err != nil {
		return User{}, err
	}

	updated.ID = objID
	return updated, nil
}

// Delete elimina un usuario por su ID de la colección.
func (r *Repository) Delete(ctx context.Context, id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}
