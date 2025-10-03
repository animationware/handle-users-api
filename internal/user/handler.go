package user

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler expone los endpoints HTTP de la API de usuarios.
// Se comunica con la capa Service para ejecutar la lógica de negocio.
type Handler struct {
	service *Service
}

// NewHandler crea una nueva instancia del Handler con la capa Service inyectada.
func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

// CreateUser maneja la creación de un nuevo usuario.
// Recibe datos JSON en la solicitud y responde con el usuario creado.
func (h *Handler) CreateUser(c *gin.Context) {
	var u User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.service.CreateUser(context.Background(), u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

// Obtener todos los usuarios
func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.service.GetUsers(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Obtener un usuario por ID
func (h *Handler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetUserByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Actualizar usuario
func (h *Handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var u User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedUser, err := h.service.UpdateUser(context.Background(), id, u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

// Eliminar usuario
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteUser(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// Método para registrar rutas
func (h *Handler) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("/", h.CreateUser)
		users.GET("/", h.GetUsers)
		users.GET("/:id", h.GetUserByID)
		users.PUT("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)
	}
}
