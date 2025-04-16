package main

import (
	"log"
	"net"
	"os"

	"github.com/Ostap00034/siproject-beercut-backend/exhibition-service/internal/db"
	"github.com/Ostap00034/siproject-beercut-backend/exhibition-service/internal/logger"
	"github.com/Ostap00034/siproject-beercut-backend/exhibition-service/internal/server"
	exhibitionpb "github.com/Ostap00034/siproject-beercut-backend/exhibition-service/proto"
	picturepb "github.com/Ostap00034/siproject-beercut-backend/picture-service/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Инициализируем ent-клиент для user-service.
	client := db.NewClient("postgres://postgres:roottoor@localhost:5432/gallery-art-exhibition?sslmode=disable")
	defer client.Close()

	pictureServiceAddr := os.Getenv("PICTURE_SERVICE_ADDR")
	if pictureServiceAddr == "" {
		pictureServiceAddr = "localhost:50055"
	}

	pictureConn, err := grpc.NewClient(pictureServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Log.Fatal("Ошибка подключения к Picture Service", zap.Error(err))
	}
	defer pictureConn.Close()
	pictureClient := picturepb.NewPictureServiceClient(pictureConn)

	// Создаем gRPC-сервер.
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	exhibitionSrv := server.NewExhibitionServer(client, pictureClient)
	exhibitionpb.RegisterExhibitionServiceServer(grpcServer, exhibitionSrv)

	lis, err := net.Listen("tcp", ":50056")
	if err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
	log.Printf("Exhibition Service запущен на %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка при работе сервера: %v", err)
	}
}
