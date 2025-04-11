package main

import (
	"net"

	"github.com/Ostap00034/siproject-beercut-backend/auth-service/internal/db"
	"github.com/Ostap00034/siproject-beercut-backend/auth-service/internal/logger"
	"github.com/Ostap00034/siproject-beercut-backend/auth-service/internal/server"
	proto "github.com/Ostap00034/siproject-beercut-backend/auth-service/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Инициализация auth-сервиса")

	// Получаем строку подключения непосредственно (или из переменной окружения)
	dsn := "postgres://postgres:roottoor@localhost:5432/gallery-art-auth?sslmode=disable"
	client := db.NewClient(dsn)
	defer client.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log.Fatal("failed to listen", zap.Error(err))
	}

	grpcServer := grpc.NewServer()
	authSrv := server.NewAuthServer(client)
	proto.RegisterAuthServiceServer(grpcServer, authSrv)

	// Регистрируем сервис reflection, чтобы grpcurl мог запрашивать описание сервисов
	reflection.Register(grpcServer)

	logger.Log.Info("Auth-сервис запущен на порту 50051")
	if err := grpcServer.Serve(lis); err != nil {
		logger.Log.Fatal("failed to serve", zap.Error(err))
	}
}
