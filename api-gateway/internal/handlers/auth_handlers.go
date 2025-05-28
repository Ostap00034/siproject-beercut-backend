package handlers

import (
	"context"
	"net/http"
	"time"

	authpb "github.com/Ostap00034/siproject-beercut-backend/auth-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// AuthHandler содержит gRPC-клиент для сервиса аутентификации.
type AuthHandler struct {
	Client authpb.AuthServiceClient
}

// NewAuthHandler создаёт новый экземпляр AuthHandler.
func NewAuthHandler(client authpb.AuthServiceClient) *AuthHandler {
	return &AuthHandler{Client: client}
}

// RegisterHandler обрабатывает запрос на регистрацию пользователя.
func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	var req authpb.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Неверный запрос",
			"errors":  gin.H{"request": err.Error()},
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.RegisterUser(ctx, &req)
	if err != nil {
		st, _ := status.FromError(err)
		httpStatus := http.StatusInternalServerError
		switch st.Code() {
		case codes.InvalidArgument:
			httpStatus = http.StatusBadRequest
		case codes.NotFound:
			httpStatus = http.StatusNotFound
		}
		errorResponse := gin.H{"message": st.Message()}
		for _, d := range st.Details() {
			if er, ok := d.(*authpb.ErrorResponse); ok {
				errorResponse["errors"] = er.Errors
			}
		}
		c.JSON(httpStatus, errorResponse)
		return
	}

	// Отправляем JSON с пустыми полями
	marshaler := protojson.MarshalOptions{EmitUnpopulated: true}
	jsonData, err := marshaler.Marshal(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", jsonData)
}

// LoginHandler обрабатывает запрос на вход пользователя и сохраняет токен в HTTP-Only cookie.
func (h *AuthHandler) LoginHandler(c *gin.Context) {
	var req authpb.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Неверный запрос",
			"errors":  gin.H{"request": err.Error()},
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.Login(ctx, &req)
	if err != nil {
		st, _ := status.FromError(err)
		httpStatus := http.StatusInternalServerError
		switch st.Code() {
		case codes.Unauthenticated:
			httpStatus = http.StatusUnauthorized
		case codes.InvalidArgument:
			httpStatus = http.StatusBadRequest
		case codes.NotFound:
			httpStatus = http.StatusNotFound
		}
		errorResponse := gin.H{"message": st.Message()}
		for _, d := range st.Details() {
			if er, ok := d.(*authpb.ErrorResponse); ok {
				errorResponse["errors"] = er.Errors
			}
		}
		c.JSON(httpStatus, errorResponse)
		return
	}

	// Извлекаем токен из ответа и сохраняем в HTTP-Only cookie
	token := resp.GetToken() // предполагается, что в LoginResponse есть поле Token
	// maxAge=0 → сессионная кука, path="/", secure=false (для prod ставьте true), httpOnly=true
	c.SetCookie("sso", token, 0, "/", "", false, true)

	// Если нужно вернуть тело ответа, можно всё так же отдать JSON:
	marshaler := protojson.MarshalOptions{EmitUnpopulated: true}
	jsonData, err := marshaler.Marshal(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", jsonData)
}

// LogoutHandler стирает куку с ключом sso.
func (h *AuthHandler) LogoutHandler(c *gin.Context) {
	// maxAge < 0 → удаляет куку, остальные параметры должны совпадать с SetCookie
	c.SetCookie("sso", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Успешный выход"})
}

// ValidateTokenHandler обрабатывает запрос на валидацию токена.
// Здесь, при желании, можно читать токен из cookie, но оставлю текущую логику по JSON.
// ValidateTokenHandler обрабатывает запрос на валидацию токена,
// извлекая токен из HTTP-Only cookie "sso".
func (h *AuthHandler) ValidateTokenHandler(c *gin.Context) {
	// Пытаемся получить токен из куки
	token, err := c.Cookie("sso")
	if err != nil || token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Отсутствует или недействительная кука sso"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Вызов gRPC-метода ValidateToken
	resp, err := h.Client.ValidateToken(ctx, &authpb.ValidateTokenRequest{Token: token})
	if err != nil {
		st, _ := status.FromError(err)
		httpStatus := http.StatusInternalServerError
		switch st.Code() {
		case codes.Unauthenticated:
			httpStatus = http.StatusUnauthorized
		case codes.InvalidArgument:
			httpStatus = http.StatusBadRequest
		case codes.NotFound:
			httpStatus = http.StatusNotFound
		}
		// Собираем детали ошибки, если есть
		errorResponse := gin.H{"message": st.Message()}
		for _, d := range st.Details() {
			if er, ok := d.(*authpb.ErrorResponse); ok {
				errorResponse["errors"] = er.Errors
			}
		}
		c.JSON(httpStatus, errorResponse)
		return
	}

	// Серилизация ответа с пустыми полями
	marshaler := protojson.MarshalOptions{EmitUnpopulated: true}
	jsonData, err := marshaler.Marshal(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Отправляем ответ
	c.Data(http.StatusOK, "application/json", jsonData)
}
