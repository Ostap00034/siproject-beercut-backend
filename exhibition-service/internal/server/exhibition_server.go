package server

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Ostap00034/siproject-beercut-backend/exhibition-service/ent"
	"github.com/Ostap00034/siproject-beercut-backend/exhibition-service/ent/exhibition"
	exhibitionpb "github.com/Ostap00034/siproject-beercut-backend/exhibition-service/proto"
	picturepb "github.com/Ostap00034/siproject-beercut-backend/picture-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// checkDuplicates проверяет, что в срезе идентификаторов отсутствуют дубликаты.
func checkDuplicates(ids []string, fieldName string) error {
	seen := make(map[string]struct{})
	for _, id := range ids {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		if _, exists := seen[id]; exists {
			return fmt.Errorf("Найдены дубликаты в %s: %s", fieldName, id)
		}
		seen[id] = struct{}{}
	}
	return nil
}

// removeDuplicates возвращает срез строк без дубликатов.
func removeDuplicates(ids []string) []string {
	seen := make(map[string]struct{})
	result := []string{}
	for _, id := range ids {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		if _, exists := seen[id]; !exists {
			seen[id] = struct{}{}
			result = append(result, id)
		}
	}
	return result
}

// sliceDiff возвращает элементы, которые присутствуют в oldSlice, но отсутствуют в newSlice.
func sliceDiff(oldSlice, newSlice []string) []string {
	diff := []string{}
	newMap := make(map[string]struct{})
	for _, s := range newSlice {
		newMap[s] = struct{}{}
	}
	for _, s := range oldSlice {
		if _, exists := newMap[s]; !exists {
			diff = append(diff, s)
		}
	}
	return diff
}

// convertGenres преобразует срез GenreData из picture-сервиса (тип picturepb.GenreData) в тип для выставки (exhibitionpb.GenreData).
func convertGenres(genres []*picturepb.GenreData) []*exhibitionpb.GenreData {
	result := make([]*exhibitionpb.GenreData, 0, len(genres))
	for _, g := range genres {
		result = append(result, &exhibitionpb.GenreData{
			Id:          g.Id,
			Name:        g.Name,
			Description: g.Description,
			CreatedAt:   g.CreatedAt,
		})
	}
	return result
}

// convertAuthors преобразует срез AuthorData из picture-сервиса (тип picturepb.AuthorData) в тип для выставки (exhibitionpb.AuthorData).
func convertAuthors(authors []*picturepb.AuthorData) []*exhibitionpb.AuthorData {
	result := make([]*exhibitionpb.AuthorData, 0, len(authors))
	for _, a := range authors {
		result = append(result, &exhibitionpb.AuthorData{
			Id:          a.Id,
			FullName:    a.FullName,
			DateOfBirth: a.DateOfBirth,
			DateOfDeath: a.DateOfDeath,
			CreatedAt:   a.CreatedAt,
		})
	}
	return result
}

// convertPictures преобразует срез PictureData из picture-сервиса в тип для выставки (exhibitionpb.PictureData).
func convertPictures(pictures []*picturepb.PictureData) []*exhibitionpb.PictureData {
	result := make([]*exhibitionpb.PictureData, 0, len(pictures))
	for _, p := range pictures {
		result = append(result, &exhibitionpb.PictureData{
			Id:             p.Id,
			Name:           p.Name,
			DateOfPainting: p.DateOfPainting,
			AuthorsIds:     p.AuthorsIds,
			GenresIds:      p.GenresIds,
			Authors:        convertAuthors(p.Authors),
			Genres:         convertGenres(p.Genres),
			ExhibitionId:   p.ExhibitionId,
			Cost:           p.Cost,
			Location:       p.Location,
			CreatedAt:      p.CreatedAt,
		})
	}
	return result
}

// checkPicturesIds проверяет список pictures_ids: отсутствие дубликатов и существование каждой картины.
func (s *exhibitionServer) checkPicturesIds(ctx context.Context, pics []string) error {
	seen := make(map[string]struct{})
	for _, id := range pics {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		if _, exists := seen[id]; exists {
			return fmt.Errorf("В списке картин обнаружены дубликаты: %s", id)
		}
		seen[id] = struct{}{}
		// Проверяем существование картины через PictureService.
		_, err := s.pictureClient.GetPicture(ctx, &picturepb.GetPictureRequest{PictureId: id})
		if err != nil {
			return fmt.Errorf("Картина с id %s не существует", id)
		}
	}
	return nil
}

// LoadPictures загружает данные картин через PictureService по списку pictures_ids.
// Возвращает: срез преобразованных PictureData для выставки и актуальный список существующих id.
func (s *exhibitionServer) LoadPictures(ctx context.Context, picturesIds []string) ([]*exhibitionpb.PictureData, []string, error) {
	validPicturesIds := []string{}
	picturesList := []*picturepb.PictureData{}
	for _, id := range picturesIds {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		resp, err := s.pictureClient.GetPicture(ctx, &picturepb.GetPictureRequest{PictureId: id})
		if err != nil {
			// Пропускаем несуществующие картины.
			continue
		}
		validPicturesIds = append(validPicturesIds, id)
		picturesList = append(picturesList, resp.Picture)
	}
	return convertPictures(picturesList), validPicturesIds, nil
}

// normalizeStatus приводит значение поля status к ожидаемому формату.
// Допустимые значения: OPENED, CLOSED. Если строка пуста, возвращается "CLOSED".
func normalizeStatus(st string) (string, error) {
	st = strings.TrimSpace(st)
	if st == "" {
		return "CLOSED", nil
	}
	upper := strings.ToUpper(st)
	if upper == "OPENED" || upper == "CLOSED" {
		return upper, nil
	}
	return "", fmt.Errorf("Неверный формат статуса: %s", st)
}

// exhibitionServer реализует интерфейс ExhibitionService.
type exhibitionServer struct {
	exhibitionpb.UnimplementedExhibitionServiceServer
	db            *ent.Client
	pictureClient picturepb.PictureServiceClient
}

func NewExhibitionServer(dbClient *ent.Client, pictureClient picturepb.PictureServiceClient) *exhibitionServer {
	return &exhibitionServer{
		db:            dbClient,
		pictureClient: pictureClient,
	}
}

// CreateExhibition создает новую выставку.
func (s *exhibitionServer) CreateExhibition(ctx context.Context, req *exhibitionpb.CreateExhibitionRequest) (*exhibitionpb.CreateExhibitionResponse, error) {
	if strings.TrimSpace(req.GetName()) == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Название выставки обязательно для заполнения")
	}
	if len(req.GetPicturesIds()) > 0 {
		if err := s.checkPicturesIds(ctx, req.GetPicturesIds()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
	}
	normStatus, err := normalizeStatus(req.GetStatus())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	exh, err := s.db.Exhibition.
		Create().
		SetName(req.GetName()).
		SetDescription(req.GetDescription()).
		SetPicturesIds(req.GetPicturesIds()).
		SetStatus(exhibition.Status(normStatus)).
		Save(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, status.Errorf(codes.InvalidArgument, "Выставка с таким названием уже существует")
		}
		return nil, status.Errorf(codes.Internal, "Ошибка создания выставки: %v", err)
	}

	return &exhibitionpb.CreateExhibitionResponse{
		Exhibition: &exhibitionpb.ExhibitionData{
			Id:          strconv.Itoa(exh.ID),
			Name:        exh.Name,
			Description: exh.Description,
			PicturesIds: exh.PicturesIds,
			Status:      exh.Status.String(),
			CreatedAt:   exh.CreatedAt.String(),
		},
		Message: "Выставка успешно создана",
	}, nil
}

// GetExhibition возвращает данные выставки по её exhibition_id.
func (s *exhibitionServer) GetExhibition(ctx context.Context, req *exhibitionpb.GetExhibitionRequest) (*exhibitionpb.GetExhibitionResponse, error) {
	id, err := strconv.Atoi(req.GetExhibitionId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат exhibition_id")
	}
	exh, err := s.db.Exhibition.Query().Where(exhibition.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Выставка не найдена")
	}
	return &exhibitionpb.GetExhibitionResponse{
		Exhibition: &exhibitionpb.ExhibitionData{
			Id:          strconv.Itoa(exh.ID),
			Name:        exh.Name,
			Description: exh.Description,
			PicturesIds: exh.PicturesIds,
			Status:      exh.Status.String(),
			CreatedAt:   exh.CreatedAt.String(),
		},
	}, nil
}

// GetAll возвращает список выставок с пагинацией. В каждом элементе удаляются дубликаты в поле PicturesIds,
// а затем загружаются подробные данные картин.
func (s *exhibitionServer) GetAll(ctx context.Context, req *exhibitionpb.GetAllRequest) (*exhibitionpb.GetAllResponse, error) {
	pageNumber := req.GetPageNumber()
	pageSize := req.GetPageSize()
	if pageNumber <= 0 {
		pageNumber = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (pageNumber - 1) * pageSize

	total, err := s.db.Exhibition.Query().Count(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка получения общего количества выставок: %v", err)
	}
	totalPages := (total + int(pageSize) - 1) / int(pageSize)

	exhs, err := s.db.Exhibition.Query().Offset(int(offset)).Limit(int(pageSize)).All(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка получения выставок: %v", err)
	}

	var exhibitionsList []*exhibitionpb.ExhibitionData
	for _, e := range exhs {
		dedupedPics := removeDuplicates(e.PicturesIds)
		pics, validPics, err := s.LoadPictures(ctx, dedupedPics)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка загрузки картин для выставки: %v", err)
		}
		exhData := &exhibitionpb.ExhibitionData{
			Id:          strconv.Itoa(e.ID),
			Name:        e.Name,
			Description: e.Description,
			PicturesIds: validPics,
			Pictures:    pics,
			Status:      e.Status.String(),
			CreatedAt:   e.CreatedAt.String(),
		}
		exhibitionsList = append(exhibitionsList, exhData)
	}
	return &exhibitionpb.GetAllResponse{
		Exhibitions: exhibitionsList,
		Total:       int32(total),
		TotalPages:  int32(totalPages),
	}, nil
}

// UpdateExhibition обновляет данные выставки.
// Если в запросе передан список PicturesIds, производится проверка на дубликаты и существование картин.
// При обновлении:
//   - Если новый список PicturesIds пустой, то для всех ранее привязанных картин обновляются их данные: устанавливается location = IN_STORAGE и exhibition_id сбрасывается.
//   - Если новый список не пустой, то для каждой картинки из нового списка вызывается PictureService.UpdatePicture для установки location = IN_EXHIBITION и exhibition_id = req.ExhibitionId.
func (s *exhibitionServer) UpdateExhibition(ctx context.Context, req *exhibitionpb.UpdateExhibitionRequest) (*exhibitionpb.UpdateExhibitionResponse, error) {
	// Преобразуем exhibition_id из строки в число.
	id, err := strconv.Atoi(req.GetExhibitionId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат exhibition_id")
	}

	// Получаем текущую запись выставки.
	exh, err := s.db.Exhibition.Query().Where(exhibition.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Выставка не найдена")
	}

	// Сохраняем старый список PicturesIds.
	oldPics := exh.PicturesIds

	if strings.TrimSpace(req.GetName()) != "" {
		exh, err = exh.Update().SetName(req.GetName()).Save(ctx)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				return nil, status.Errorf(codes.InvalidArgument, "Выставка с таким названием уже существует")
			}
			return nil, status.Errorf(codes.Internal, "Ошибка обновления имени: %v", err)
		}
	}
	if req.GetDescription() != "" {
		exh, err = exh.Update().SetDescription(req.GetDescription()).Save(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка обновления описания: %v", err)
		}
	}

	// Обновляем список PicturesIds.
	if req.GetPicturesIds() != nil {
		// Проверяем наличие дубликатов и существование картин.
		if err := s.checkPicturesIds(ctx, req.GetPicturesIds()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		exh, err = exh.Update().SetPicturesIds(req.GetPicturesIds()).Save(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка обновления списка картин: %v", err)
		}
	}

	if req.GetStatus() != "" {
		normStatus, err := normalizeStatus(req.GetStatus())
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		exh, err = exh.Update().SetStatus(exhibition.Status(normStatus)).Save(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка обновления статуса: %v", err)
		}
	}

	// Если новый список PicturesIds задан и пустой, то:
	// Обновляем все картинки из старого списка, устанавливая для них location = IN_STORAGE и exhibition_id = "".
	if req.GetPicturesIds() == nil && len(req.GetPicturesIds()) == 0 {
		for _, picID := range oldPics {
			updPicReq := &picturepb.UpdatePictureRequest{
				PictureId:    picID,
				Location:     "IN_STORAGE",
				ExhibitionId: "",
				UserId:       req.UserId,
			}
			_, err := s.pictureClient.UpdatePicture(ctx, updPicReq)
			if err != nil {
				fmt.Printf("Ошибка обновления картинки (удалена из выставки) с id %s: %v\n", picID, err)
			}
		}
		exh, err = exh.Update().SetPicturesIds([]string{}).Save(ctx)
		if err != nil {
			fmt.Printf("Ошибка обновления выставки %v", err)
		}
	} else if req.GetPicturesIds() != nil && len(req.GetPicturesIds()) > 0 {
		// Если новый список не пустой, для каждой картинки из нового списка:
		// устанавливаем location = IN_EXHIBITION и exhibition_id = req.ExhibitionId.
		newPics := req.GetPicturesIds()
		for _, picID := range newPics {
			updPicReq := &picturepb.UpdatePictureRequest{
				PictureId:    picID,
				Location:     "IN_EXHIBITION",
				ExhibitionId: req.GetExhibitionId(),
			}
			_, err := s.pictureClient.UpdatePicture(ctx, updPicReq)
			if err != nil {
				fmt.Printf("Ошибка обновления картинки с id %s: %v\n", picID, err)
			}
		}
	}

	// Загружаем связанные картинки для выставки.
	pics, validPics, err := s.LoadPictures(ctx, exh.PicturesIds)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка загрузки картин для выставки: %v", err)
	}

	return &exhibitionpb.UpdateExhibitionResponse{
		Exhibition: &exhibitionpb.ExhibitionData{
			Id:          strconv.Itoa(exh.ID),
			Name:        exh.Name,
			Description: exh.Description,
			PicturesIds: validPics,
			Pictures:    pics,
			Status:      exh.Status.String(),
			CreatedAt:   exh.CreatedAt.String(),
		},
	}, nil
}

// DeleteExhibition удаляет выставку по её exhibition_id.
func (s *exhibitionServer) DeleteExhibition(ctx context.Context, req *exhibitionpb.DeleteExhibitionRequest) (*exhibitionpb.DeleteExhibitionResponse, error) {
	id, err := strconv.Atoi(req.GetExhibitionId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат exhibition_id")
	}
	_, err = s.db.Exhibition.Query().Where(exhibition.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Выставка не найдена")
	}
	_, err = s.db.Exhibition.Delete().Where(exhibition.IDEQ(id)).Exec(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка при удалении выставки: %v", err)
	}
	return &exhibitionpb.DeleteExhibitionResponse{
		Message: "Выставка успешно удалена",
	}, nil
}
