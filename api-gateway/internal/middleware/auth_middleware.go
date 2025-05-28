package middleware

import (
	"context"
	"net/http"
	"time"

	authpb "github.com/Ostap00034/siproject-beercut-backend/auth-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// AuthClientHolder содержит глобальный gRPC-клиент для Auth Service.
var AuthClient authpb.AuthServiceClient

// InitAuthClient инициализирует gRPC-клиента для Auth Service.
// addr - адрес Auth Service (например, "localhost:50051")
func InitAuthClient(addr string) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Ошибка подключения к Auth Service: " + err.Error())
	}
	AuthClient = authpb.NewAuthServiceClient(conn)
}

// tokenFromCookie пытается получить токен из куки "sso".
func tokenFromCookie(c *gin.Context) (string, error) {
	token, err := c.Cookie("sso")
	if err != nil {
		return "", err
	}
	return token, nil
}

// AdminMiddleware проверяет, что в куке "sso" лежит валидный токен с ролью ADMIN.
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из куки
		tokenStr, err := tokenFromCookie(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Отсутствует или недействительная кука sso"})
			return
		}

		// Контекст с таймаутом для проверки токена
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Валидация токена через gRPC
		validateReq := &authpb.ValidateTokenRequest{Token: tokenStr}
		resp, err := AuthClient.ValidateToken(ctx, validateReq)
		if err != nil || !resp.GetValid() {
			st, _ := status.FromError(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Неавторизовано: " + st.Message()})
			return
		}

		// Проверяем роль
		if resp.GetRole() != "ADMIN" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Доступ запрещён: требуется роль ADMIN"})
			return
		}

		// Сохраняем данные пользователя в контекст
		c.Set("user_id", resp.GetUserId())
		c.Set("caller_role", resp.GetRole())

		c.Next()
	}
}

// EmployeeMiddleware проверяет, что в куке "sso" лежит валидный токен с ролью EMPLOYEE или ADMIN.
func EmployeeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из куки
		tokenStr, err := tokenFromCookie(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Отсутствует или недействительная кука sso"})
			return
		}

		// Контекст с таймаутом для проверки токена
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Валидация токена через gRPC
		validateReq := &authpb.ValidateTokenRequest{Token: tokenStr}
		resp, err := AuthClient.ValidateToken(ctx, validateReq)
		if err != nil || !resp.GetValid() {
			st, _ := status.FromError(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Неавторизовано: " + st.Message()})
			return
		}

		// Проверяем роль
		role := resp.GetRole()
		if role != "EMPLOYEE" && role != "ADMIN" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Доступ запрещён: требуется роль EMPLOYEE"})
			return
		}

		// Сохраняем данные пользователя в контекст
		c.Set("user_id", resp.GetUserId())
		c.Set("caller_role", role)

		c.Next()
	}
}
