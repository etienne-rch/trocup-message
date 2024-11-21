# Trocup Message Service

## Description

This is a message management microservice built with Go, Fiber, and MongoDB.

## Setup

1. Install Go
2. Clone the repository
3. Run `go mod tidy`
4. Set up MongoDB Atlas and update the URI in `config/config.go`
5. Run `copy .env.example to .env`
6. Run the service: `go run main.go`

go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

go get -u github.com/swaggo/fiber-swagger
go get -u github.com/swaggo/swag/cmd/swag
go get github.com/gofiber/swagger