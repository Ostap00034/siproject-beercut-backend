package main

import (
	"log"
	"net"

	"github.com/Ostap00034/siproject-beercut-backend/user-service/internal/db"
	"github.com/Ostap00034/siproject-beercut-backend/user-service/internal/server"
	userpb "github.com/Ostap00034/siproject-beercut-backend/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Инициализируем ent-клиент для user-service.
	client := db.NewClient("postgres://postgres:roottoor@localhost:5432/gallery-art-user?sslmode=disable")
	defer client.Close()

	// Создаем gRPC-сервер.
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	userSrv := server.NewUserServer(client)
	userpb.RegisterUserServiceServer(grpcServer, userSrv)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
	log.Printf("User Service запущен на %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка при работе сервера: %v", err)
	}
}
