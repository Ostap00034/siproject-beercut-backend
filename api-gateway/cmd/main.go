package main

import (
	"time"

	"github.com/Ostap00034/siproject-beercut-backend/api-gateway/internal/handlers"
	"github.com/Ostap00034/siproject-beercut-backend/api-gateway/internal/middleware"
	authpb "github.com/Ostap00034/siproject-beercut-backend/auth-service/proto"
	authorpb "github.com/Ostap00034/siproject-beercut-backend/author-service/proto"
	exhibitionpb "github.com/Ostap00034/siproject-beercut-backend/exhibition-service/proto"
	genrepb "github.com/Ostap00034/siproject-beercut-backend/genre-service/proto"
	movementhistorypb "github.com/Ostap00034/siproject-beercut-backend/movementhistory-service/proto"
	picturepb "github.com/Ostap00034/siproject-beercut-backend/picture-service/proto"
	userpb "github.com/Ostap00034/siproject-beercut-backend/user-service/proto"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	authConn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer authConn.Close()

	userConn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer userConn.Close()

	genreConn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer genreConn.Close()

	authorConn, err := grpc.NewClient("localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer authorConn.Close()

	pictureConn, err := grpc.NewClient("localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer pictureConn.Close()

	exhibitionConn, err := grpc.NewClient("localhost:50056", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer exhibitionConn.Close()

	movementhistoryConn, err := grpc.NewClient("localhost:50057", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer movementhistoryConn.Close()

	// Clients
	authClient := authpb.NewAuthServiceClient(authConn)
	userClient := userpb.NewUserServiceClient(userConn)
	genreClient := genrepb.NewGenreServiceClient(genreConn)
	authorClient := authorpb.NewAuthorServiceClient(authorConn)
	pictureClient := picturepb.NewPictureServiceClient(pictureConn)
	exhibitionClient := exhibitionpb.NewExhibitionServiceClient(exhibitionConn)
	movementhistoryClient := movementhistorypb.NewMovementHistoryServiceClient(movementhistoryConn)

	// Handlers
	authHandler := handlers.NewAuthHandler(authClient)
	userHandler := handlers.NewUserHandler(userClient)
	genreHandler := handlers.NewGenreHandler(genreClient)
	authorHandler := handlers.NewAuthorHandler(authorClient)
	pictureHandler := handlers.NewPictureHandler(pictureClient)
	exhibitionHandler := handlers.NewExhibitionHandler(exhibitionClient)
	movementhistoryHandler := handlers.NewMovementHistoryServer(movementhistoryClient)

	// Инициализация Gin
	router := gin.Default()

	// Подключаем CORS middleware для разрешения запросов с localhost:5173
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	middleware.InitAuthClient("localhost:50051")

	// Создаем группу с префиксом /api
	api := router.Group("/api")
	{
		// Группа для ADMIN
		adminGroup := api.Group("/admin")
		adminGroup.Use(middleware.AdminMiddleware())
		{
			adminGroup.GET("/users/", userHandler.GetAllUsersHandler)
			adminGroup.POST("/users/create", authHandler.RegisterHandler)
			adminGroup.PUT("/users/:user_id", userHandler.UpdateUserHandler)
			// Genres
			adminGroup.DELETE("/genres/:genre_id", genreHandler.DeleteGenreHandler)
			// Authors
			adminGroup.DELETE("/authors/:author_id", authorHandler.DeleteAuthorHandler)
			// Pictures
			adminGroup.DELETE("/pictures/:picture_id", pictureHandler.DeletePictureHandler)
			// Exhibitions
			adminGroup.DELETE("/exhibitions/:exhibition_id", exhibitionHandler.DeleteExhibitionHandler)
			// MovementHistories
			adminGroup.DELETE("/movementhistory/:movementhistory_id", movementhistoryHandler.DeleteMovementHistoryHandler)
		}

		// Группа для сотрудника или других ролей (например, EMPLOYEE)
		employeeGroup := api.Group("/")
		employeeGroup.Use(middleware.EmployeeMiddleware())
		{
			employeeGroup.POST("/genres/", genreHandler.CreateGenreHandler)
			employeeGroup.PUT("/genres/:genre_id", genreHandler.UpdateGenreHandler)
			employeeGroup.GET("/genres/", genreHandler.GetAllGenresHandler)
			employeeGroup.GET("/genres/:genre_id", genreHandler.GetGenreHandler)

			employeeGroup.GET("/authors/", authorHandler.GetAllAuthorsHandler)
			employeeGroup.GET("/authors/:author_id", authorHandler.GetAuthorHandler)
			employeeGroup.POST("/authors/", authorHandler.CreateAuthorHandler)
			employeeGroup.PUT("/authors/:author_id", authorHandler.UpdateAuthorHandler)

			employeeGroup.POST("/pictures/", pictureHandler.CreatePictureHandler)
			employeeGroup.GET("/pictures/", pictureHandler.GetAllPicturesHandler)
			employeeGroup.GET("/pictures/:picture_id", pictureHandler.GetPictureHandler)
			employeeGroup.PUT("/pictures/:picture_id", pictureHandler.UpdatePictureHandler)

			employeeGroup.POST("/exhibitions/", exhibitionHandler.CreateExhibitionHandler)
			employeeGroup.GET("/exhibitions/", exhibitionHandler.GetAllExhibitionsHandler)
			employeeGroup.GET("/exhibitions/:exhibition_id", exhibitionHandler.GetExhibitionHandler)
			employeeGroup.PUT("/exhibitions/:exhibition_id", exhibitionHandler.UpdateExhibitionHandler)

			employeeGroup.POST("/movementhistory/", movementhistoryHandler.CreateMovementHistoryHandler)
			employeeGroup.GET("/movementhistory/", movementhistoryHandler.GetAllMovementHistoryHandler)
			employeeGroup.GET("/movementhistory/:movementhistory_id", movementhistoryHandler.GetMovementHistoryHandler)
			employeeGroup.GET("/movementhistory/picture/:picture_id", movementhistoryHandler.GetMovementHistorysByPictureIdHandler)
		}

		// Обработчики, не требующие авторизации (например, login и validate)
		api.POST("/auth/login", authHandler.LoginHandler)
		api.POST("/auth/logout", authHandler.LogoutHandler)
		api.POST("/auth/token/validate", authHandler.ValidateTokenHandler)
	}

	// Запускаем HTTP-сервер на порту 8080
	router.Run(":8080")
}
