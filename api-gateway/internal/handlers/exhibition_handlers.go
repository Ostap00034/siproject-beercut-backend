package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	exhibitionpb "github.com/Ostap00034/siproject-beercut-backend/exhibition-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// ExhibitionHandler инкапсулирует gRPC-клиент для Exhibition Service.
type ExhibitionHandler struct {
	Client exhibitionpb.ExhibitionServiceClient
}

// NewExhibitionHandler создаёт новый экземпляр ExhibitionHandler.
func NewExhibitionHandler(client exhibitionpb.ExhibitionServiceClient) *ExhibitionHandler {
	return &ExhibitionHandler{
		Client: client,
	}
}

// CreateExhibitionHandler обрабатывает POST /exhibitions/create и вызывает метод CreateExhibition.
func (h *ExhibitionHandler) CreateExhibitionHandler(c *gin.Context) {
	var req exhibitionpb.CreateExhibitionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Неверный запрос",
			"errors":  err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.CreateExhibition(ctx, &req)
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

// GetExhibitionHandler обрабатывает GET /exhibitions/:exhibition_id и вызывает метод GetExhibition.
func (h *ExhibitionHandler) GetExhibitionHandler(c *gin.Context) {
	exhibitionID := c.Param("exhibition_id")
	if exhibitionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "exhibition_id не указан"})
		return
	}

	req := &exhibitionpb.GetExhibitionRequest{
		ExhibitionId: exhibitionID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.GetExhibition(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		httpStatus := http.StatusInternalServerError
		switch st.Code() {
		case codes.InvalidArgument:
			httpStatus = http.StatusBadRequest
		case codes.NotFound:
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

// GetAllExhibitionsHandler обрабатывает GET /exhibitions и вызывает метод GetAll.
func (h *ExhibitionHandler) GetAllExhibitionsHandler(c *gin.Context) {
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

	req := &exhibitionpb.GetAllRequest{
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

// UpdateExhibitionHandler обрабатывает PUT /exhibitions/:exhibition_id и вызывает метод UpdateExhibition.
// Идентификатор выставки извлекается из URL, остальные поля — из тела запроса.
func (h *ExhibitionHandler) UpdateExhibitionHandler(c *gin.Context) {
	exhibitionID := c.Param("exhibition_id")
	if exhibitionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "exhibition_id не указан"})
		return
	}

	var req exhibitionpb.UpdateExhibitionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Неверный запрос",
			"errors":  err.Error(),
		})
		return
	}
	userId, exist := c.Get("user_id")
	if !exist {
		status.Errorf(codes.InvalidArgument, "Ошибка получения user_id")
	}

	// Передаем exhibition_id из URL в запрос.
	req.ExhibitionId = exhibitionID
	req.UserId = userId.(string)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.UpdateExhibition(ctx, &req)
	if err != nil {
		st, _ := status.FromError(err)
		httpStatus := http.StatusInternalServerError
		switch st.Code() {
		case codes.InvalidArgument:
			httpStatus = http.StatusBadRequest
		case codes.NotFound:
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

// DeleteExhibitionHandler обрабатывает DELETE /exhibitions/:exhibition_id и вызывает метод DeleteExhibition.
func (h *ExhibitionHandler) DeleteExhibitionHandler(c *gin.Context) {
	exhibitionID := c.Param("exhibition_id")
	if exhibitionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "exhibition_id не указан"})
		return
	}

	req := &exhibitionpb.DeleteExhibitionRequest{
		ExhibitionId: exhibitionID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.DeleteExhibition(ctx, req)
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
