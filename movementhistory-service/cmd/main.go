package main

import (
	"net"
	"os"

	"github.com/Ostap00034/siproject-beercut-backend/movementhistory-service/internal/db"
	"github.com/Ostap00034/siproject-beercut-backend/movementhistory-service/internal/logger"
	"github.com/Ostap00034/siproject-beercut-backend/movementhistory-service/internal/server"
	movementhistorypb "github.com/Ostap00034/siproject-beercut-backend/movementhistory-service/proto"
	picturepb "github.com/Ostap00034/siproject-beercut-backend/picture-service/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("Инициализация Movement-history-сервиса")

	dsn := os.Getenv("MOVEMENT_HISTORY_DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:roottoor@localhost:5432/gallery-art-movement-history?sslmode=disable"
	}
	client := db.NewClient(dsn)
	defer client.Close()

	lis, err := net.Listen("tcp", ":50057")
	if err != nil {
		logger.Log.Fatal("Ошибка создания TCP-листенера", zap.Error(err))
	}

	pictureServiceAddr := os.Getenv("PICTURE_SERVICE_ADDR")
	if pictureServiceAddr == "" {
		pictureServiceAddr = "localhost:50055"
	}

	pictureConn, err := grpc.NewClient(pictureServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Log.Fatal("Ошибка подключения к Exhibition Service", zap.Error(err))
	}
	defer pictureConn.Close()
	pictureClient := picturepb.NewPictureServiceClient(pictureConn)

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	movementhistorySrv := server.NewMovementHistoryServer(client, pictureClient)
	movementhistorypb.RegisterMovementHistoryServiceServer(grpcServer, movementhistorySrv)

	reflection.Register(grpcServer)

	logger.Log.Info("Movement-history-сервис запущен на порту 50057")
	if err := grpcServer.Serve(lis); err != nil {
		logger.Log.Fatal("Ошибка работы gRPC-сервера", zap.Error(err))
	}
}
