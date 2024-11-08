# Build stage
FROM golang:1.22-alpine AS builder

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers go.mod et go.sum et installer les dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le reste de l'application
COPY . .

# Construire l'application
RUN go build -o app

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 5004

CMD ["./app"]
