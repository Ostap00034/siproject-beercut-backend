package main

import (
	"net"
	"os"

	"github.com/Ostap00034/siproject-beercut-backend/genre-service/internal/db"
	"github.com/Ostap00034/siproject-beercut-backend/genre-service/internal/logger"
	"github.com/Ostap00034/siproject-beercut-backend/genre-service/internal/server"
	genrepb "github.com/Ostap00034/siproject-beercut-backend/genre-service/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Инициализация Genre-сервиса")

	dsn := os.Getenv("GENRE_DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:roottoor@localhost:5432/gallery-art-genre?sslmode=disable"
	}
	client := db.NewClient(dsn)
	defer client.Close()

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		logger.Log.Fatal("Ошибка создания TCP-листенера", zap.Error(err))
	}

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	genreSrv := server.NewGenreServer(client)
	genrepb.RegisterGenreServiceServer(grpcServer, genreSrv)

	reflection.Register(grpcServer)

	logger.Log.Info("Genre-сервис запущен на порту 50053")
	if err := grpcServer.Serve(lis); err != nil {
		logger.Log.Fatal("Ошибка работы gRPC-сервера", zap.Error(err))
	}
}
