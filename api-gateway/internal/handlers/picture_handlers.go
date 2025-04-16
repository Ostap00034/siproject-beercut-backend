package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	picturepb "github.com/Ostap00034/siproject-beercut-backend/picture-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// PictureHandler инкапсулирует gRPC-клиент для Picture Service.
type PictureHandler struct {
	Client picturepb.PictureServiceClient
}

// NewPictureHandler создаёт новый экземпляр PictureHandler.
func NewPictureHandler(client picturepb.PictureServiceClient) *PictureHandler {
	return &PictureHandler{
		Client: client,
	}
}

// CreatePictureHandler обрабатывает POST /pictures/create и вызывает метод CreatePicture.
func (h *PictureHandler) CreatePictureHandler(c *gin.Context) {
	var req picturepb.CreatePictureRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Неверный запрос",
			"errors":  err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.CreatePicture(ctx, &req)
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

// GetPictureHandler обрабатывает GET /pictures/:picture_id и вызывает метод GetPicture.
func (h *PictureHandler) GetPictureHandler(c *gin.Context) {
	pictureID := c.Param("picture_id")
	if pictureID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "picture_id не указан"})
		return
	}

	req := &picturepb.GetPictureRequest{
		PictureId: pictureID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.GetPicture(ctx, req)
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

// GetAllPicturesHandler обрабатывает GET /pictures и вызывает метод GetAll.
func (h *PictureHandler) GetAllPicturesHandler(c *gin.Context) {
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

	req := &picturepb.GetAllRequest{
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

// UpdatePictureHandler обрабатывает PUT /pictures/:picture_id.
func (h *PictureHandler) UpdatePictureHandler(c *gin.Context) {
	pictureID := c.Param("picture_id")
	if pictureID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "picture_id не указан"})
		return
	}

	var req picturepb.UpdatePictureRequest
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

	// Передаем picture_id из URL в запрос.
	req.PictureId = pictureID
	req.UserId = userId.(string)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.UpdatePicture(ctx, &req)
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

// DeletePictureHandler обрабатывает DELETE /pictures/:picture_id.
func (h *PictureHandler) DeletePictureHandler(c *gin.Context) {
	pictureID := c.Param("picture_id")
	if pictureID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "picture_id не указан"})
		return
	}

	req := &picturepb.DeletePictureRequest{
		PictureId: pictureID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.DeletePicture(ctx, req)
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
