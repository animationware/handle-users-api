# Usa la imagen oficial de Go
FROM golang:1.25-alpine AS builder

# Establecer directorio de trabajo
WORKDIR /app

# Copiar los archivos de dependencias primero (mejora caché en builds)
COPY go.mod go.sum ./
RUN go mod download

# Copiar todo el código
COPY . .

# Compilar la API
RUN go build -o handle-users-api ./cmd/main.go

# Imagen final (más liviana, solo ejecutable)
FROM debian:bookworm-slim

WORKDIR /app

# Copiar binario desde el builder
COPY --from=builder /app/handle-users-api .

# Exponer el puerto de la API
EXPOSE 3000

# Comando por defecto al iniciar el contenedor
CMD ["./handle-users-api"]
