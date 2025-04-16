package server

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/Ostap00034/siproject-beercut-backend/movementhistory-service/ent"
	"github.com/Ostap00034/siproject-beercut-backend/movementhistory-service/ent/movementhistory"
	movementhistorypb "github.com/Ostap00034/siproject-beercut-backend/movementhistory-service/proto"
	picturepb "github.com/Ostap00034/siproject-beercut-backend/picture-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// movementHistoryServer реализует интерфейс MovementHistoryService.
type movementHistoryServer struct {
	movementhistorypb.UnimplementedMovementHistoryServiceServer
	db            *ent.Client
	pictureClient picturepb.PictureServiceClient
}

// NewMovementHistoryServer создаёт новый экземпляр movementHistoryServer.
func NewMovementHistoryServer(dbClient *ent.Client, pictureClient picturepb.PictureServiceClient) *movementHistoryServer {
	return &movementHistoryServer{
		db:            dbClient,
		pictureClient: pictureClient,
	}
}

// CreateMovementHistory создаёт новую запись истории перемещений.
func (s *movementHistoryServer) CreateMovementHistory(ctx context.Context, req *movementhistorypb.CreateMovementHistoryRequest) (*movementhistorypb.CreateMovementHistoryResponse, error) {
	// Проверка обязательных полей.
	if strings.TrimSpace(req.GetUserId()) == "" {
		return nil, status.Errorf(codes.InvalidArgument, "user_id не может быть пустым")
	}
	if strings.TrimSpace(req.GetPictureId()) == "" {
		return nil, status.Errorf(codes.InvalidArgument, "picture_id не может быть пустым")
	}

	// Дополнительная проверка для enum-полей "from" и "to".
	allowedValues := map[string]bool{
		"IN_STORAGE":     true,
		"IN_EXHIBITION":  true,
		"IN_RESTORATION": true,
	}

	if !allowedValues[req.GetFrom()] {
		return nil, status.Errorf(codes.InvalidArgument, "Недопустимое значение поля from: %s", req.GetFrom())
	}
	if !allowedValues[req.GetTo()] {
		return nil, status.Errorf(codes.InvalidArgument, "Недопустимое значение поля to: %s", req.GetTo())
	}

	// Создание записи в базе данных.
	mh, err := s.db.MovementHistory.
		Create().
		SetUserID(req.GetUserId()).
		SetPictureID(req.GetPictureId()).
		SetFrom(movementhistory.From(req.GetFrom())).
		SetTo(movementhistory.To(req.GetTo())).
		Save(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка создания записи: %v", err)
	}

	data := &movementhistorypb.MovementHistoryData{
		Id:        strconv.Itoa(mh.ID),
		UserId:    mh.UserID,
		PictureId: mh.PictureID,
		From:      mh.From.String(),
		To:        mh.To.String(),
		CreatedAt: mh.CreatedAt.Format(time.RFC3339),
	}

	return &movementhistorypb.CreateMovementHistoryResponse{
		Movementhistory: data,
		Message:         "Запись истории перемещений успешно создана",
	}, nil
}

// GetMovementHistoryByPictureId возвращает запись истории перемещений по идентификатору.
func (s *movementHistoryServer) GetMovementHistorysByPictureId(ctx context.Context, req *movementhistorypb.GetMovementHistorysByPictureIdRequest) (*movementhistorypb.GetMovementHistorysByPictureIdResponse, error) {
	pictureId := req.GetPictureId()

	var pictureReq picturepb.GetPictureRequest

	pictureReq.PictureId = req.GetPictureId()

	_, err := s.pictureClient.GetPicture(ctx, &pictureReq)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Картина не найдена")
	}

	pageNumber := req.GetPageNumber()
	pageSize := req.GetPageSize()
	if pageNumber <= 0 {
		pageNumber = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (pageNumber - 1) * pageSize

	mhs, err := s.db.MovementHistory.Query().Offset(int(offset)).Limit(int(pageSize)).Where(movementhistory.PictureIDEQ(pictureId)).All(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка получения записей: %v", err)
	}

	total, err := s.db.MovementHistory.Query().Count(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка получения общего количества записей: %v", err)
	}

	totalPages := (total + int(pageSize) - 1) / int(pageSize)
	var list []*movementhistorypb.MovementHistoryData
	for _, mh := range mhs {
		list = append(list, &movementhistorypb.MovementHistoryData{
			Id:        strconv.Itoa(mh.ID),
			UserId:    mh.UserID,
			PictureId: mh.PictureID,
			From:      mh.From.String(),
			To:        mh.To.String(),
			CreatedAt: mh.CreatedAt.Format(time.RFC3339),
		})
	}

	return &movementhistorypb.GetMovementHistorysByPictureIdResponse{
		Movementhistorys: list,
		Total:            int32(total),
		TotalPages:       int32(totalPages),
	}, nil
}

// GetMovementHistory возвращает запись истории перемещений по идентификатору.
func (s *movementHistoryServer) GetMovementHistory(ctx context.Context, req *movementhistorypb.GetMovementHistoryRequest) (*movementhistorypb.GetMovementHistoryResponse, error) {
	id, err := strconv.Atoi(req.GetMovementhistoryId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат movementhistory_id")
	}

	mh, err := s.db.MovementHistory.Query().Where(movementhistory.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Запись истории перемещений не найдена")
	}

	data := &movementhistorypb.MovementHistoryData{
		Id:        strconv.Itoa(mh.ID),
		UserId:    mh.UserID,
		PictureId: mh.PictureID,
		From:      mh.From.String(),
		To:        mh.To.String(),
		CreatedAt: mh.CreatedAt.Format(time.RFC3339),
	}

	return &movementhistorypb.GetMovementHistoryResponse{
		Movementhistory: data,
	}, nil
}

// GetAll возвращает список записей истории перемещений с пагинацией.
func (s *movementHistoryServer) GetAll(ctx context.Context, req *movementhistorypb.GetAllRequest) (*movementhistorypb.GetAllResponse, error) {
	pageNumber := req.GetPageNumber()
	pageSize := req.GetPageSize()
	if pageNumber <= 0 {
		pageNumber = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (pageNumber - 1) * pageSize

	mhs, err := s.db.MovementHistory.Query().Offset(int(offset)).Limit(int(pageSize)).All(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка получения записей: %v", err)
	}

	total, err := s.db.MovementHistory.Query().Count(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка получения общего количества записей: %v", err)
	}

	totalPages := (total + int(pageSize) - 1) / int(pageSize)
	var list []*movementhistorypb.MovementHistoryData
	for _, mh := range mhs {
		list = append(list, &movementhistorypb.MovementHistoryData{
			Id:        strconv.Itoa(mh.ID),
			UserId:    mh.UserID,
			PictureId: mh.PictureID,
			From:      mh.From.String(),
			To:        mh.To.String(),
			CreatedAt: mh.CreatedAt.Format(time.RFC3339),
		})
	}

	return &movementhistorypb.GetAllResponse{
		Movementhistorys: list,
		Total:            int32(total),
		TotalPages:       int32(totalPages),
	}, nil
}

// DeleteMovementHistory удаляет запись истории перемещений по идентификатору.
func (s *movementHistoryServer) DeleteMovementHistory(ctx context.Context, req *movementhistorypb.DeleteMovementHistoryRequest) (*movementhistorypb.DeleteMovementHistoryResponse, error) {
	id, err := strconv.Atoi(req.GetMovementhistoryId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат movementhistory_id")
	}

	// Проверяем, существует ли запись.
	_, err = s.db.MovementHistory.Query().Where(movementhistory.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Запись истории перемещений не найдена")
	}

	_, err = s.db.MovementHistory.Delete().Where(movementhistory.IDEQ(id)).Exec(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка удаления записи: %v", err)
	}

	return &movementhistorypb.DeleteMovementHistoryResponse{
		Message: "Запись истории перемещений успешно удалена",
	}, nil
}
