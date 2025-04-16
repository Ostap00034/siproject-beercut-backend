package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	movementhistorypb "github.com/Ostap00034/siproject-beercut-backend/movementhistory-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

type MovementHistoryHandler struct {
	Client movementhistorypb.MovementHistoryServiceClient
}

func NewMovementHistoryServer(client movementhistorypb.MovementHistoryServiceClient) *MovementHistoryHandler {
	return &MovementHistoryHandler{
		Client: client,
	}
}

func (h *MovementHistoryHandler) CreateMovementHistoryHandler(c *gin.Context) {
	var req movementhistorypb.CreateMovementHistoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Неверный запрос",
			"errors":  err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.CreateMovementHistory(ctx, &req)
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

func (h *MovementHistoryHandler) GetMovementHistorysByPictureIdHandler(c *gin.Context) {
	pictureID := c.Param("picture_id")
	if pictureID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "picture_id не указан"})
		return
	}

	req := &movementhistorypb.GetMovementHistorysByPictureIdRequest{
		PictureId: pictureID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.GetMovementHistorysByPictureId(ctx, req)
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

func (h *MovementHistoryHandler) GetMovementHistoryHandler(c *gin.Context) {
	movementhistoryID := c.Param("movementhistory_id")
	if movementhistoryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "movementhistory_id не указан"})
		return
	}

	req := &movementhistorypb.GetMovementHistoryRequest{
		MovementhistoryId: movementhistoryID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.GetMovementHistory(ctx, req)
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

func (h *MovementHistoryHandler) GetAllMovementHistoryHandler(c *gin.Context) {
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

	req := &movementhistorypb.GetAllRequest{
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

func (h *MovementHistoryHandler) DeleteMovementHistoryHandler(c *gin.Context) {
	movementhistoryID := c.Param("movementhistory_id")
	if movementhistoryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "movementhistory_id не указан"})
		return
	}

	req := &movementhistorypb.DeleteMovementHistoryRequest{
		MovementhistoryId: movementhistoryID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := h.Client.DeleteMovementHistory(ctx, req)
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
