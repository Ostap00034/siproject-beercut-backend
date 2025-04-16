package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	authpb "github.com/Ostap00034/siproject-beercut-backend/auth-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// AuthClientHolder содержит глобальный gRPC-клиент для Auth Service,
// который используется для валидации токенов.
var AuthClient authpb.AuthServiceClient

// InitAuthClient инициализирует gRPC-клиента для Auth Service.
// addr - адрес Auth Service (например, "localhost:50051")
func InitAuthClient(addr string) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Ошибка подключения к Auth Service: " + err.Error())
	}
	// Можно сохранить соединение в глобальную переменную, если это необходимо.
	AuthClient = authpb.NewAuthServiceClient(conn)
}

// AdminMiddleware проверяет, что в заголовке Authorization передан токен типа "Bearer <token>",
// и что этот токен валиден и принадлежит пользователю с ролью ADMIN.
// Если проверка не проходит, запрос прерывается с соответствующим HTTP статусом.
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлекаем заголовок Authorization.
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Отсутствует токен авторизации"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Неверный формат токена"})
			return
		}
		tokenStr := parts[1]

		// Создаем контекст с таймаутом для проверки токена.
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Вызываем метод ValidateToken через gRPC Auth Service.
		validateReq := &authpb.ValidateTokenRequest{Token: tokenStr}
		resp, err := AuthClient.ValidateToken(ctx, validateReq)
		if err != nil || !resp.GetValid() {
			st, _ := status.FromError(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Неавторизовано: " + st.Message()})
			return
		}

		// Проверяем, что роль пользователя равна "ADMIN".
		if resp.GetRole() != "ADMIN" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Доступ запрещён: требуется роль ADMIN"})
			return
		}

		// Если проверка пройдена, можно сохранить информацию в контексте для дальнейшего использования.
		c.Set("user_id", resp.GetUserId())
		c.Set("caller_role", resp.GetRole())

		c.Next()
	}
}

func EmployeeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлекаем заголовок Authorization.
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Отсутствует токен авторизации"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Неверный формат токена"})
			return
		}
		tokenStr := parts[1]

		// Создаем контекст с таймаутом для проверки токена.
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Вызываем метод ValidateToken через gRPC Auth Service.
		validateReq := &authpb.ValidateTokenRequest{Token: tokenStr}
		resp, err := AuthClient.ValidateToken(ctx, validateReq)
		if err != nil || !resp.GetValid() {
			st, _ := status.FromError(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Неавторизовано: " + st.Message()})
			return
		}

		// Проверяем, что роль пользователя равна "ADMIN".
		if resp.GetRole() != "EMPLOYEE" && resp.GetRole() != "ADMIN" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Доступ запрещён: требуется роль EMPLOYEE"})
			return
		}

		// Если проверка пройдена, можно сохранить информацию в контексте для дальнейшего использования.
		c.Set("user_id", resp.GetUserId())
		c.Set("caller_role", resp.GetRole())

		c.Next()
	}
}
