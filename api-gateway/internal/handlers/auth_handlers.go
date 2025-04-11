package handlers

import (
	"context"
	"net/http"
	"time"

	authpb "github.com/Ostap00034/siproject-beercut-backend/auth-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		st, ok := status.FromError(err)
		httpStatus := http.StatusInternalServerError
		if ok {
			if st.Code() == codes.InvalidArgument {
				httpStatus = http.StatusBadRequest
			} else if st.Code() == codes.NotFound {
				httpStatus = http.StatusNotFound
			}
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

	c.JSON(http.StatusOK, resp)
}

// LoginHandler обрабатывает запрос на вход (авторизацию) пользователя.
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
		st, ok := status.FromError(err)
		httpStatus := http.StatusInternalServerError
		if ok {
			if st.Code() == codes.Unauthenticated {
				httpStatus = http.StatusUnauthorized
			} else if st.Code() == codes.InvalidArgument {
				httpStatus = http.StatusBadRequest
			} else if st.Code() == codes.NotFound {
				httpStatus = http.StatusNotFound
			}
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

	c.JSON(http.StatusOK, resp)
}

func (h *AuthHandler) ValidateTokenHandler(c *gin.Context) {
	var req authpb.ValidateTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Неверный запрос",
			"errors":  gin.H{"request": err.Error()},
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.ValidateToken(ctx, &req)
	if err != nil {
		st, ok := status.FromError(err)
		httpStatus := http.StatusInternalServerError
		if ok {
			if st.Code() == codes.Unauthenticated {
				httpStatus = http.StatusUnauthorized
			} else if st.Code() == codes.InvalidArgument {
				httpStatus = http.StatusBadRequest
			} else if st.Code() == codes.NotFound {
				httpStatus = http.StatusNotFound
			}
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

	c.JSON(http.StatusOK, resp)
}
