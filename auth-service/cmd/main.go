package main

import (
	"net"
	"os"

	"github.com/Ostap00034/siproject-beercut-backend/auth-service/internal/db"
	"github.com/Ostap00034/siproject-beercut-backend/auth-service/internal/logger"
	"github.com/Ostap00034/siproject-beercut-backend/auth-service/internal/server"
	authpb "github.com/Ostap00034/siproject-beercut-backend/auth-service/proto"
	userpb "github.com/Ostap00034/siproject-beercut-backend/user-service/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Инициализация Auth-сервиса")

	// Чтение DSN для базы данных Auth Service из переменной окружения или использование значения по умолчанию.
	dsn := os.Getenv("AUTH_DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:roottoor@localhost:5432/gallery-art-auth?sslmode=disable"
	}
	authDBClient := db.NewClient(dsn)
	defer authDBClient.Close()

	// Чтение адреса User Service из переменной окружения или значение по умолчанию.
	userServiceAddr := os.Getenv("USER_SERVICE_ADDR")
	if userServiceAddr == "" {
		userServiceAddr = "localhost:50052"
	}
	// Создаем gRPC-соединение с User Service.
	userConn, err := grpc.NewClient(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Log.Fatal("Ошибка подключения к User Service", zap.Error(err))
	}
	defer userConn.Close()
	userClient := userpb.NewUserServiceClient(userConn)

	// Создаем TCP-листенер на порту 50051 для Auth Service.
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log.Fatal("Ошибка создания TCP-листенера", zap.Error(err))
	}

	// Создаем новый gRPC-сервер с insecure credentials (для локальной разработки).
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	// Создаем экземпляр Auth-сервиса, передавая ent-клиент Auth Service и gRPC-клиент User Service.
	authSrv := server.NewAuthServer(authDBClient, userClient)
	authpb.RegisterAuthServiceServer(grpcServer, authSrv)

	// Регистрируем reflection для возможности запроса описания сервисов (например, через grpcurl).
	reflection.Register(grpcServer)

	logger.Log.Info("Auth-сервис запущен на порту 50051")
	if err := grpcServer.Serve(lis); err != nil {
		logger.Log.Fatal("Ошибка работы gRPC-сервера", zap.Error(err))
	}
}
