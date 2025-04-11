// cmd/main.go
package main

import (
	"net"
	"os"

	"github.com/Ostap00034/siproject-beercut-backend/auth-service/internal/db"
	"github.com/Ostap00034/siproject-beercut-backend/auth-service/internal/logger"
	"github.com/Ostap00034/siproject-beercut-backend/auth-service/internal/server"
	"github.com/Ostap00034/siproject-beercut-backend/auth-service/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	// Инициализация логгера
	logger.InitLogger()
	logger.Log.Info("Инициализация auth-сервиса")

	// Инициализация ent-клиента (строка подключения может быть передана через переменные окружения)
	dsn := os.Getenv("postgres://postgres:roottoor@localhost:5432/gallery-art-auth?sslmode=disable") // например: postgres://username:password@localhost:5432/dbname?sslmode=disable
	client := db.NewClient(dsn)
	defer client.Close()

	// Создание TCP-листенера на порту 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log.Fatal("failed to listen", zap.Error(err))
	}

	// Создание gRPC-сервера
	grpcServer := grpc.NewServer()

	// Создание экземпляра сервера авторизации с передачей ent-клиента
	authSrv := server.NewAuthServer(client)
	proto.RegisterAuthServiceServer(grpcServer, authSrv)

	logger.Log.Info("Auth-сервис запущен на порту 50051")
	if err := grpcServer.Serve(lis); err != nil {
		logger.Log.Fatal("failed to serve", zap.Error(err))
	}
}
