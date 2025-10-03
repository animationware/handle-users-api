# User API en Go + MongoDB

Mini API RESTful para gestión de usuarios, desarrollada como prueba técnica.  
Implementada en **Golang**, usando **Gin** para el servidor HTTP y **MongoDB** como base de datos.

---

## Características

- Crear usuario (POST)
- Listar todos los usuarios (GET)
- Consultar usuario por ID (GET)
- Actualizar usuario (PUT)
- Eliminar usuario (DELETE)
- Arquitectura en capas: `handler → service → repository`
- Proyecto estructurado con **Project Layout de Go**
- Despliegue con **Docker Compose**

---

## Estructura del proyecto

/handle-users-api
├── cmd/
│ └── main.go # Punto de entrada
├── internal/
│ ├── database/ # Conexión Mongo
│ │ └── mongo.go
│ ├── user/ # Dominio User
│ │ ├── model.go
│ │ ├── repository.go
│ │ ├── service.go
│ │ └── handler.go
├── go.mod
├── go.sum
├── README.md
└── docker-compose.yml
└── Dockerfile

## Base de Datos

La API utiliza **MongoDB** como almacén de datos.  
Gracias a la integración con **Docker Compose**, la base de datos se levanta automáticamente junto con la API y no requiere scripts adicionales de creación.

- **Nombre de la base de datos:** `handle-users-db`  
- **Colección principal:** `users`  
- **Conexión:** definida en la variable de entorno `MONGO_URI` en `docker-compose.yml`  
  ```text
  mongodb://mongo:27017/handle-users-db


## Ejecución

## Requisitos
- Go >= 1.20
- Docker >= 20.x
- Docker Compose (v1 o v2)

Verifica las instalaciones:
    go version
    docker --version
    docker compose version

## Docker Compose (API + MongoDB en contenedores)
1. Clonar el repo:
   ```bash
   git clone https://github.com/animationware/handle-users-api.git
   cd handle-users-api
   ```
2.  Realiza la instalación de dependencias necesarias y creación de archivo go.sum para realizar verificación checksum de dependencias:
   go mod tidy

2. Levantar MongoDB (local o en docker):
    docker compose up --build

## Información Adicional y Pruebas funcionales:
1. La API estará disponible en http://localhost:3000
2. MongoDB estará disponible en mongodb://mongo:27017/handle-users-db
3. Creación de Usuario:
   curl -X POST http://localhost:3000/users/    -H "Content-Type: application/json"    -d '{"name":"Michael Romero","email":"michael@example.com"}'
4. Obtener todos los Usuarios:
   curl http://localhost:3000/users/
5. Actualizar usuario:
   curl -X PUT http://localhost:3000/users/<ID-Usuario> -H "Content-Type: application/json" -d '{"name":"Nombre Actualizado","email":"mail.updated@example.com"}'
6. Eliminar usuario:
   curl -X DELETE http://localhost:3000/users/<ID-Usuario>
7. Obtener usuario por ID:
   curl http://localhost:3000/users/<ID-Usuario>
   
## Arquitectura y Patrones

Se utilizó el patrón Repository-Service-Handler (arquitectura en capas):

- **Handler:** recibe y valida solicitudes HTTP.
- **Service:** contiene la lógica de negocio y reglas.
- **Repository:** maneja la persistencia en MongoDB.

Este patrón permite separación de responsabilidades, facilita pruebas unitarias y hace que la API sea mantenible y escalable.

## Autor

**Michael Romero**  
Desarrollador Backend Jr.  
[GitHub](https://github.com/animationware/DevelopWare)  
[Email](mailto:michaelromeroortega@gmail.com)
