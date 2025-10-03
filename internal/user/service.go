package user

import (
	"context"
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Service representa la capa de negocio de la aplicación.
// Contiene la lógica de validación y coordinación entre la capa HTTP (Handler) y la capa de persistencia (Repository).
type Service struct {
	repo *Repository
}

// NewService crea una nueva instancia del Service con el Repository inyectado.
func NewService(repo *Repository) *Service {
	return &Service{repo}
}

// CreateUser valida y crea un nuevo usuario.
// Retorna el usuario creado con su ID generado por MongoDB.
func (s *Service) CreateUser(ctx context.Context, u User) (User, error) {
	// Validaciones de negocio
	if strings.TrimSpace(u.Name) == "" {
		return User{}, errors.New("el nombre es obligatorio")
	}
	if !strings.Contains(u.Email, "@") {
		return User{}, errors.New("email inválido")
	}

	res, err := s.repo.Create(ctx, u)
	if err != nil {
		return User{}, err
	}

	// Retornar el usuario creado con ID
	u.ID = res.InsertedID.(primitive.ObjectID)
	return u, nil
}

// GetUsers obtiene todos los usuarios disponibles.
func (s *Service) GetUsers(ctx context.Context) ([]User, error) {
	return s.repo.FindAll(ctx)
}

// GetUserByID obtiene un usuario específico por su ID.
func (s *Service) GetUserByID(ctx context.Context, id string) (User, error) {
	return s.repo.FindByID(ctx, id)
}

// UpdateUser valida y actualiza un usuario existente.
// Retorna el usuario actualizado.
func (s *Service) UpdateUser(ctx context.Context, id string, u User) (User, error) {
	// Validaciones simples
	if u.Email == "" {
		return User{}, errors.New("el email no puede estar vacío")
	}

	updated, err := s.repo.Update(ctx, id, u)
	if err != nil {
		return User{}, err
	}
	return updated, nil
}

// DeleteUser elimina un usuario por su ID.
func (s *Service) DeleteUser(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
