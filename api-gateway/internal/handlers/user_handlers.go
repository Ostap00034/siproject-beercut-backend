package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	userpb "github.com/Ostap00034/siproject-beercut-backend/user-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

type UserHandler struct {
	Client userpb.UserServiceClient
}

// NewAuthHandler создаёт новый экземпляр AuthHandler.
func NewUserHandler(client userpb.UserServiceClient) *UserHandler {
	return &UserHandler{Client: client}
}

func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var req userpb.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Неверный запрос",
			"errors":  gin.H{"request": err.Error()},
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.CreateUser(ctx, &req)
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
			if er, ok := d.(*userpb.ErrorResponse); ok {
				errorResponse["errors"] = er.Errors
			}
		}
		c.JSON(httpStatus, errorResponse)
		return
	}

	// Используем кастомный маршалер для сериализации ответа с пустыми полями
	marshaler := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	jsonData, err := marshaler.Marshal(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Отправляем сформированный JSON напрямую
	c.Data(http.StatusOK, "application/json", jsonData)
}

func (h *UserHandler) GetAllUsersHandler(c *gin.Context) {
	// Извлекаем параметры пагинации из строки запроса с значениями по умолчанию.
	pageNumberStr := c.DefaultQuery("page_number", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil || pageNumber <= 0 {
		c.JSON(400, gin.H{"message": "Неверное значение параметра page_number"})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		c.JSON(400, gin.H{"message": "Неверное значение параметра page_size"})
		return
	}

	req := &userpb.GetAllUsersRequest{
		PageNumber: int32(pageNumber),
		PageSize:   int32(pageSize),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.GetAllUsers(ctx, req)
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
			if er, ok := d.(*userpb.ErrorResponse); ok {
				errorResponse["errors"] = er.Errors
			}
		}
		c.JSON(httpStatus, errorResponse)
		return
	}

	// Используем кастомный маршалер для сериализации ответа с пустыми полями
	marshaler := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	jsonData, err := marshaler.Marshal(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Отправляем сформированный JSON напрямую
	c.Data(http.StatusOK, "application/json", jsonData)
}

func (h *UserHandler) GetUserHandler(c *gin.Context) {
	var req userpb.GetUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Неверный запрос",
			"errors":  gin.H{"request": err.Error()},
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.GetUser(ctx, &req)
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
			if er, ok := d.(*userpb.ErrorResponse); ok {
				errorResponse["errors"] = er.Errors
			}
		}
		c.JSON(httpStatus, errorResponse)
		return
	}

	// Используем кастомный маршалер для сериализации ответа с пустыми полями
	marshaler := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	jsonData, err := marshaler.Marshal(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Отправляем сформированный JSON напрямую
	c.Data(http.StatusOK, "application/json", jsonData)
}

func (h *UserHandler) UpdateUserHandler(c *gin.Context) {
	var req userpb.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Неверный запрос",
			"errors":  gin.H{"request": err.Error()},
		})
		return
	}

	req.UserId = c.Param("user_id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.UpdateUser(ctx, &req)
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
			if er, ok := d.(*userpb.ErrorResponse); ok {
				errorResponse["errors"] = er.Errors
			}
		}
		c.JSON(httpStatus, errorResponse)
		return
	}

	// Используем кастомный маршалер для сериализации ответа с пустыми полями
	marshaler := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	jsonData, err := marshaler.Marshal(resp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Отправляем сформированный JSON напрямую
	c.Data(http.StatusOK, "application/json", jsonData)
}
