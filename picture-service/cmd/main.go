package main

import (
	"log"
	"net"
	"os"

	authorpb "github.com/Ostap00034/siproject-beercut-backend/author-service/proto"
	exhibitionpb "github.com/Ostap00034/siproject-beercut-backend/exhibition-service/proto"
	genrepb "github.com/Ostap00034/siproject-beercut-backend/genre-service/proto"
	movementhistorypb "github.com/Ostap00034/siproject-beercut-backend/movementhistory-service/proto"
	"github.com/Ostap00034/siproject-beercut-backend/picture-service/internal/db"
	"github.com/Ostap00034/siproject-beercut-backend/picture-service/internal/logger"
	"github.com/Ostap00034/siproject-beercut-backend/picture-service/internal/server"
	picturepb "github.com/Ostap00034/siproject-beercut-backend/picture-service/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Инициализируем ent-клиент для user-service.
	client := db.NewClient("postgres://postgres:roottoor@localhost:5432/gallery-art-picture?sslmode=disable")
	defer client.Close()

	genreServiceAddr := os.Getenv("GENRE_SERVICE_ADDR")
	if genreServiceAddr == "" {
		genreServiceAddr = "localhost:50053"
	}

	genreConn, err := grpc.NewClient(genreServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Log.Fatal("Ошибка подключения к Genre Service", zap.Error(err))
	}
	defer genreConn.Close()
	genreClient := genrepb.NewGenreServiceClient(genreConn)

	authorServiceAddr := os.Getenv("AUTHOR_SERVICE_ADDR")
	if authorServiceAddr == "" {
		authorServiceAddr = "localhost:50054"
	}

	authorConn, err := grpc.NewClient(authorServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Log.Fatal("Ошибка подключения к Author Service", zap.Error(err))
	}
	defer authorConn.Close()
	authorClient := authorpb.NewAuthorServiceClient(authorConn)

	exhibitionServiceAddr := os.Getenv("EXHIBITION_SERVICE_ADDR")
	if exhibitionServiceAddr == "" {
		exhibitionServiceAddr = "localhost:50056"
	}

	exhibitionConn, err := grpc.NewClient(exhibitionServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Log.Fatal("Ошибка подключения к Exhibition Service", zap.Error(err))
	}
	defer authorConn.Close()
	exhibitionClient := exhibitionpb.NewExhibitionServiceClient(exhibitionConn)

	movementhistoryServiceAddr := os.Getenv("MOVEMENT_HISTORY_SERVICE_ADDR")
	if movementhistoryServiceAddr == "" {
		movementhistoryServiceAddr = "localhost:50057"
	}

	movementhistoryConn, err := grpc.NewClient(movementhistoryServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Log.Fatal("Ошибка подключения к Exhibition Service", zap.Error(err))
	}
	defer authorConn.Close()
	movementhistoryClient := movementhistorypb.NewMovementHistoryServiceClient(movementhistoryConn)

	// Создаем gRPC-сервер.
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	pictureSrv := server.NewPictureServer(client, authorClient, genreClient, exhibitionClient, movementhistoryClient)
	picturepb.RegisterPictureServiceServer(grpcServer, pictureSrv)

	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
	log.Printf("Picture Service запущен на %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка при работе сервера: %v", err)
	}
}
