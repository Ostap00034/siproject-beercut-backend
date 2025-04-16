package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	authorpb "github.com/Ostap00034/siproject-beercut-backend/author-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// AuthorHandler инкапсулирует gRPC‑клиент для Author Service.
type AuthorHandler struct {
	Client authorpb.AuthorServiceClient
}

// NewAuthorHandler создает новый экземпляр AuthorHandler.
func NewAuthorHandler(client authorpb.AuthorServiceClient) *AuthorHandler {
	return &AuthorHandler{
		Client: client,
	}
}

// CreateAuthorHandler обрабатывает POST /authors/create и вызывает метод CreateAuthor.
func (h *AuthorHandler) CreateAuthorHandler(c *gin.Context) {
	var req authorpb.CreateAuthorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Неверный запрос",
			"errors":  err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.CreateAuthor(ctx, &req)
	if err != nil {
		st, _ := status.FromError(err)
		httpStatus := http.StatusInternalServerError
		if st.Code() == codes.InvalidArgument {
			httpStatus = http.StatusBadRequest
		}
		c.JSON(httpStatus, gin.H{"message": st.Message()})
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

// GetAuthorHandler обрабатывает GET /authors/:author_id и вызывает метод GetAuthor.
func (h *AuthorHandler) GetAuthorHandler(c *gin.Context) {
	authorID := c.Param("author_id")
	if authorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "author_id не указан"})
		return
	}

	req := &authorpb.GetAuthorRequest{
		AuthorId: authorID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.GetAuthor(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		httpStatus := http.StatusInternalServerError
		if st.Code() == codes.InvalidArgument {
			httpStatus = http.StatusBadRequest
		} else if st.Code() == codes.NotFound {
			httpStatus = http.StatusNotFound
		}
		c.JSON(httpStatus, gin.H{"message": st.Message()})
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

// GetAllAuthorsHandler обрабатывает GET /authors и вызывает метод GetAll.
func (h *AuthorHandler) GetAllAuthorsHandler(c *gin.Context) {
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

	req := &authorpb.GetAllRequest{
		PageNumber: int32(pageNumber),
		PageSize:   int32(pageSize),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.GetAll(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": st.Message()})
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

// UpdateAuthorHandler обрабатывает PUT /authors/:author_id.
// Значение author_id извлекается из URL, остальные поля берутся из JSON тела запроса.
func (h *AuthorHandler) UpdateAuthorHandler(c *gin.Context) {
	authorID := c.Param("author_id")
	if authorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "author_id не указан"})
		return
	}

	var req authorpb.UpdateAuthorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Неверный запрос", "errors": err.Error()})
		return
	}
	// Устанавливаем author_id из URL.
	req.AuthorId = authorID

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.UpdateAuthor(ctx, &req)
	if err != nil {
		st, _ := status.FromError(err)
		httpStatus := http.StatusInternalServerError
		if st.Code() == codes.InvalidArgument {
			httpStatus = http.StatusBadRequest
		} else if st.Code() == codes.NotFound {
			httpStatus = http.StatusNotFound
		}
		c.JSON(httpStatus, gin.H{"message": st.Message()})
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

// DeleteAuthorHandler обрабатывает DELETE /authors/:author_id и вызывает метод DeleteAuthor.
func (h *AuthorHandler) DeleteAuthorHandler(c *gin.Context) {
	authorID := c.Param("author_id")
	if authorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "author_id не указан"})
		return
	}

	req := &authorpb.DeleteAuthorRequest{
		AuthorId: authorID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.DeleteAuthor(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": st.Message()})
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
