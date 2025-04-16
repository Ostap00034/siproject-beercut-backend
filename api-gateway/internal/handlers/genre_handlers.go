package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	genrepb "github.com/Ostap00034/siproject-beercut-backend/genre-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// GenreHandler инкапсулирует gRPC‑клиент для Genre Service.
type GenreHandler struct {
	Client genrepb.GenreServiceClient
}

// NewGenreHandler создает новый экземпляр GenreHandler.
func NewGenreHandler(client genrepb.GenreServiceClient) *GenreHandler {
	return &GenreHandler{
		Client: client,
	}
}

// CreateGenreHandler обрабатывает POST /genres/create и вызывает метод CreateGenre.
func (h *GenreHandler) CreateGenreHandler(c *gin.Context) {
	var req genrepb.CreateGenreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Неверный запрос",
			"errors":  err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.CreateGenre(ctx, &req)
	if err != nil {
		st, _ := status.FromError(err)
		httpStatus := http.StatusInternalServerError
		if st.Code() == codes.InvalidArgument {
			httpStatus = http.StatusBadRequest
		}
		c.JSON(httpStatus, gin.H{
			"message": st.Message(),
		})
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

// GetGenreHandler обрабатывает GET /genres/:genre_id и вызывает метод GetGenre.
func (h *GenreHandler) GetGenreHandler(c *gin.Context) {
	genreID := c.Param("genre_id")
	if genreID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "genre_id не указан"})
		return
	}

	req := &genrepb.GetGenreRequest{
		GenreId: genreID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.GetGenre(ctx, req)
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

// GetAllGenresHandler обрабатывает GET /genres и ожидает параметры пагинации из query-параметров:
// page_number и page_size.
func (h *GenreHandler) GetAllGenresHandler(c *gin.Context) {
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

	req := &genrepb.GetAllRequest{
		PageNumber: int32(pageNumber),
		PageSize:   int32(pageSize),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.GetAll(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		httpStatus := 500
		if st.Code() == codes.InvalidArgument {
			httpStatus = 400
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

// UpdateGenreHandler обрабатывает PUT /genres/:genre_id и вызывает метод UpdateGenre.
// ID жанра извлекается из параметра URL, а остальные поля передаются в теле запроса.
func (h *GenreHandler) UpdateGenreHandler(c *gin.Context) {
	genreID := c.Param("genre_id")
	if genreID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "genre_id не указан"})
		return
	}

	var req genrepb.UpdateGenreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Неверный запрос", "errors": err.Error()})
		return
	}
	// Передаем genre_id из URL в запрос.
	req.GenreId = genreID

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.UpdateGenre(ctx, &req)
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

// DeleteGenreHandler обрабатывает DELETE /genres/:genre_id и вызывает метод DeleteGenre.
func (h *GenreHandler) DeleteGenreHandler(c *gin.Context) {
	genreID := c.Param("genre_id")
	if genreID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "genre_id не указан"})
		return
	}

	req := &genrepb.DeleteGenreRequest{
		GenreId: genreID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.DeleteGenre(ctx, req)
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
