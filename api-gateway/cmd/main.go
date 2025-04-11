// Пример простого HTTP‑обработчика через Gin для авторизации.
package main

import (
	"github.com/Ostap00034/siproject-beercut-backend/api-gateway/internal/handlers"
	authpb "github.com/Ostap00034/siproject-beercut-backend/auth-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// authClient — глобальный или инициализируется в main.
var authClient authpb.AuthServiceClient

func main() {
	// Создаем подключение к gRPC-серверу.
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	authClient := authpb.NewAuthServiceClient(conn)

	authHandler := handlers.NewAuthHandler(authClient)

	// Инициализация Gin
	router := gin.Default()

	// Обработчик для Login
	router.POST("/auth/login", authHandler.LoginHandler)

	// Обработчик для Register
	router.POST("/auth/register", authHandler.RegisterHandler)

	// Обработчик для ValidateToken
	router.POST("/auth/token/validate", authHandler.ValidateTokenHandler)

	// Запускаем HTTP-сервер на порту 8080
	router.Run(":8080")
}
