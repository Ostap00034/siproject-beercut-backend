package server

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	authorpb "github.com/Ostap00034/siproject-beercut-backend/author-service/proto"
	exhibitionpb "github.com/Ostap00034/siproject-beercut-backend/exhibition-service/proto"
	genrepb "github.com/Ostap00034/siproject-beercut-backend/genre-service/proto"
	movementhistorypb "github.com/Ostap00034/siproject-beercut-backend/movementhistory-service/proto"
	"github.com/Ostap00034/siproject-beercut-backend/picture-service/ent"
	"github.com/Ostap00034/siproject-beercut-backend/picture-service/ent/picture"
	picturepb "github.com/Ostap00034/siproject-beercut-backend/picture-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// checkDuplicates проверяет, что в срезе идентификаторов отсутствуют дубликаты.
// fieldName используется для формирования сообщения об ошибке (например, "авторах" или "жанрах").
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

func convertGenres(genres []*genrepb.GenreData) []*picturepb.GenreData {
	result := make([]*picturepb.GenreData, 0, len(genres))
	for _, g := range genres {
		result = append(result, &picturepb.GenreData{
			Id:          g.Id,
			Name:        g.Name,
			Description: g.Description,
			CreatedAt:   g.CreatedAt,
		})
	}
	return result
}

func convertAuthors(authors []*authorpb.AuthorData) []*picturepb.AuthorData {
	result := make([]*picturepb.AuthorData, 0, len(authors))
	for _, g := range authors {
		result = append(result, &picturepb.AuthorData{
			Id:          g.Id,
			FullName:    g.FullName,
			DateOfBirth: g.DateOfBirth,
			DateOfDeath: g.DateOfDeath,
			CreatedAt:   g.CreatedAt,
		})
	}
	return result
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

// pictureServer реализует интерфейс PictureServiceServer.
type pictureServer struct {
	picturepb.UnimplementedPictureServiceServer
	db                    *ent.Client
	authorClient          authorpb.AuthorServiceClient
	genreClient           genrepb.GenreServiceClient
	exhibitionClient      exhibitionpb.ExhibitionServiceClient
	movementhistoryClient movementhistorypb.MovementHistoryServiceClient
}

func NewPictureServer(dbClient *ent.Client, authorClient authorpb.AuthorServiceClient, genreClient genrepb.GenreServiceClient, exhibitionClient exhibitionpb.ExhibitionServiceClient, movementhistoryClient movementhistorypb.MovementHistoryServiceClient) *pictureServer {
	return &pictureServer{
		db:                    dbClient,
		authorClient:          authorClient,
		genreClient:           genreClient,
		exhibitionClient:      exhibitionClient,
		movementhistoryClient: movementhistoryClient,
	}
}

// normalizeLocation приводит значение локации к ожидаемому формату.
func normalizeLocation(loc string) (string, error) {
	loc = strings.TrimSpace(loc)
	if loc == "" {
		return "IN_STORAGE", nil
	}
	upper := strings.ToUpper(loc)
	switch upper {
	case "IN_STORAGE", "IN_EXHIBITION", "IN_RESTORATION":
		return upper, nil
	default:
		return "", fmt.Errorf("Неверный формат локации: %s", loc)
	}
}

// LoadAuthorsAndGenres загружает данные авторов и жанров для списка идентификаторов.
// Возвращает срезы заполненных объектов и актуальные списки id.
func (s *pictureServer) LoadAuthorsAndGenres(ctx context.Context, authorsIds []string, genresIds []string) ([]*picturepb.AuthorData, []*picturepb.GenreData, []string, []string, error) {
	validAuthorsIDs := []string{}
	authorsList := []*authorpb.AuthorData{}
	for _, id := range authorsIds {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		resp, err := s.authorClient.GetAuthor(ctx, &authorpb.GetAuthorRequest{AuthorId: id})
		if err != nil {
			// Пропускаем несуществующих авторов.
			continue
		}
		validAuthorsIDs = append(validAuthorsIDs, id)
		authorsList = append(authorsList, resp.Author)
	}

	validGenresIDs := []string{}
	genresList := []*genrepb.GenreData{}
	for _, id := range genresIds {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		resp, err := s.genreClient.GetGenre(ctx, &genrepb.GetGenreRequest{GenreId: id})
		if err != nil {
			continue
		}
		validGenresIDs = append(validGenresIDs, id)
		genresList = append(genresList, resp.Genre)
	}

	return convertAuthors(authorsList), convertGenres(genresList), validAuthorsIDs, validGenresIDs, nil
}

// CreatePicture создает новую картину с проверкой дубликатов, нормализацией локации и загрузкой авторов/жанров.
func (s *pictureServer) CreatePicture(ctx context.Context, req *picturepb.CreatePictureRequest) (*picturepb.CreatePictureResponse, error) {
	if strings.TrimSpace(req.GetName()) == "" {
		errDetail := &picturepb.ErrorResponse{
			Message: "Ошибка валидации",
			Errors:  map[string]string{"name": "Название картины обязательно для заполнения"},
		}
		st, err2 := status.New(codes.InvalidArgument, "Название отсутствует").WithDetails(errDetail)
		if err2 != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка добавления деталей: %v", err2)
		}
		return nil, st.Err()
	}
	if strings.TrimSpace(req.GetDateOfPainting()) == "" {
		errDetail := &picturepb.ErrorResponse{
			Message: "Ошибка валидации",
			Errors:  map[string]string{"date_of_painting": "Дата написания картины обязательна"},
		}
		st, err2 := status.New(codes.InvalidArgument, "Дата написания отсутствует").WithDetails(errDetail)
		if err2 != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка добавления деталей: %v", err2)
		}
		return nil, st.Err()
	}
	// Проверка дубликатов в списках авторов и жанров.
	if len(req.GetAuthorsIds()) > 0 {
		if err := checkDuplicates(req.GetAuthorsIds(), "авторах"); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		for _, authorID := range req.GetAuthorsIds() {
			_, err := s.authorClient.GetAuthor(ctx, &authorpb.GetAuthorRequest{AuthorId: authorID})
			if err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "Автор с id %s не найден: %v", authorID, err)
			}
		}
	}
	if len(req.GetGenresIds()) > 0 {
		if err := checkDuplicates(req.GetGenresIds(), "жанрах"); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		for _, genreID := range req.GetGenresIds() {
			_, err := s.genreClient.GetGenre(ctx, &genrepb.GetGenreRequest{GenreId: genreID})
			if err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "Жанр с id %s не найден: %v", genreID, err)
			}
		}
	}

	normLoc, err := normalizeLocation(req.GetLocation())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	// Сохраняем картину.
	p, err := s.db.Picture.
		Create().
		SetName(req.GetName()).
		SetDateOfPainting(req.GetDateOfPainting()).
		SetLocation(picture.Location(normLoc)).
		SetCost(req.GetCost()).
		SetExhibitionID(req.GetExhibitionId()).
		SetAuthorsIds(req.GetAuthorsIds()).
		SetGenresIds(req.GetGenresIds()).
		Save(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, status.Errorf(codes.InvalidArgument, "Картина с таким названием уже существует")
		}
		return nil, status.Errorf(codes.Internal, "Ошибка создания картины: %v", err)
	}

	// Загружаем полные данные авторов и жанров, удаляем несуществующие.
	authors, genres, validAuthorsIds, validGenresIds, err := s.LoadAuthorsAndGenres(ctx, p.AuthorsIds, p.GenresIds)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка загрузки авторов и жанров: %v", err)
	}

	return &picturepb.CreatePictureResponse{
		Picture: &picturepb.PictureData{
			Id:             strconv.Itoa(p.ID),
			Name:           p.Name,
			DateOfPainting: p.DateOfPainting,
			GenresIds:      validGenresIds,
			AuthorsIds:     validAuthorsIds,
			ExhibitionId:   p.ExhibitionID,
			Cost:           p.Cost,
			Location:       p.Location.String(),
			CreatedAt:      p.CreatedAt.String(),
			Genres:         genres,
			Authors:        authors,
		},
		Message: "Картина успешно создана",
	}, nil
}

// GetPicture возвращает данные картины, включая связанные авторов и жанры, и очищает несуществующие id.
func (s *pictureServer) GetPicture(ctx context.Context, req *picturepb.GetPictureRequest) (*picturepb.GetPictureResponse, error) {
	id, err := strconv.Atoi(req.GetPictureId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат picture_id")
	}
	p, err := s.db.Picture.Query().Where(picture.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Картина не найдена")
	}

	// Фильтруем несуществующих авторов и жанры.
	authors, genres, validAuthorsIds, validGenresIds, err := s.LoadAuthorsAndGenres(ctx, p.AuthorsIds, p.GenresIds)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка загрузки авторов и жанров: %v", err)
	}

	return &picturepb.GetPictureResponse{
		Picture: &picturepb.PictureData{
			Id:             strconv.Itoa(p.ID),
			Name:           p.Name,
			DateOfPainting: p.DateOfPainting,
			GenresIds:      validGenresIds,
			AuthorsIds:     validAuthorsIds,
			ExhibitionId:   p.ExhibitionID,
			Cost:           p.Cost,
			Location:       p.Location.String(),
			CreatedAt:      p.CreatedAt.String(),
			Genres:         genres,
			Authors:        authors,
		},
	}, nil
}

// GetAll возвращает список картин с пагинацией.
func (s *pictureServer) GetAll(ctx context.Context, req *picturepb.GetAllRequest) (*picturepb.GetAllResponse, error) {
	pageNumber := req.GetPageNumber()
	pageSize := req.GetPageSize()
	if pageNumber <= 0 {
		pageNumber = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (pageNumber - 1) * pageSize

	pictures, err := s.db.Picture.Query().Offset(int(offset)).Limit(int(pageSize)).All(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка получения картин: %v", err)
	}

	total, err := s.db.Picture.Query().Count(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка подсчета картин: %v", err)
	}
	totalPages := (total + int(pageSize) - 1) / int(pageSize)

	var pictureList []*picturepb.PictureData
	for _, p := range pictures {
		// Фильтруем несуществующих авторов и жанры для каждой картины.
		authors, genres, validAuthorsIds, validGenresIds, err := s.LoadAuthorsAndGenres(ctx, p.AuthorsIds, p.GenresIds)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка загрузки авторов и жанров: %v", err)
		}
		pictureList = append(pictureList, &picturepb.PictureData{
			Id:             strconv.Itoa(p.ID),
			Name:           p.Name,
			DateOfPainting: p.DateOfPainting,
			GenresIds:      validGenresIds,
			AuthorsIds:     validAuthorsIds,
			ExhibitionId:   p.ExhibitionID,
			Cost:           p.Cost,
			Location:       p.Location.String(),
			CreatedAt:      p.CreatedAt.String(),
			Genres:         genres,
			Authors:        authors,
		})
	}

	return &picturepb.GetAllResponse{
		Pictures:   pictureList,
		Total:      int32(total),
		TotalPages: int32(totalPages),
	}, nil
}

// UpdatePicture обновляет картину. Если location равна IN_EXHIBITION, проверяется существование выставки,
// а также выполняется фильтрация авторов и жанров.
func (s *pictureServer) UpdatePicture(ctx context.Context, req *picturepb.UpdatePictureRequest) (*picturepb.UpdatePictureResponse, error) {
	// Преобразуем picture_id из строки в число.
	id, err := strconv.Atoi(req.GetPictureId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат picture_id")
	}

	// Получаем текущую запись картины.
	p, err := s.db.Picture.Query().Where(picture.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Картина не найдена")
	}

	fromLocation := p.Location.String()

	// Сохраняем старые значения location и exhibition_id для последующего сравнения.
	oldLoc := p.Location.String()
	oldExhibitionID := p.ExhibitionID

	// Обновляем поля, если они переданы в запросе.
	if strings.TrimSpace(req.GetName()) != "" {
		p, err = p.Update().SetName(req.GetName()).Save(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка обновления названия: %v", err)
		}
	}

	if strings.TrimSpace(req.GetDateOfPainting()) != "" {
		p, err = p.Update().SetDateOfPainting(req.GetDateOfPainting()).Save(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка обновления даты написания: %v", err)
		}
	}

	if len(req.GetAuthorsIds()) > 0 {
		if err := checkDuplicates(req.GetAuthorsIds(), "авторах"); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		for _, authorID := range req.GetAuthorsIds() {
			_, err := s.authorClient.GetAuthor(ctx, &authorpb.GetAuthorRequest{AuthorId: authorID})
			if err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "Автор с id %s не найден: %v", authorID, err)
			}
		}
		p, err = p.Update().SetAuthorsIds(req.GetAuthorsIds()).Save(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка обновления списка авторов: %v", err)
		}
	}

	if len(req.GetGenresIds()) > 0 {
		if err := checkDuplicates(req.GetGenresIds(), "жанрах"); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		for _, genreID := range req.GetGenresIds() {
			_, err := s.genreClient.GetGenre(ctx, &genrepb.GetGenreRequest{GenreId: genreID})
			if err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "Жанр с id %s не найден: %v", genreID, err)
			}
		}
		p, err = p.Update().SetGenresIds(req.GetGenresIds()).Save(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка обновления списка жанров: %v", err)
		}
	}

	if req.GetCost() != 0 {
		p, err = p.Update().SetCost(req.GetCost()).Save(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка обновления стоимости: %v", err)
		}
	}

	// Обновляем поле location (и exhibition_id, если требуется).
	if strings.TrimSpace(req.GetLocation()) != "" {
		normLoc, err := normalizeLocation(req.GetLocation())
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}

		_, err = s.movementhistoryClient.CreateMovementHistory(ctx, &movementhistorypb.CreateMovementHistoryRequest{
			PictureId: req.PictureId,
			UserId:    req.UserId,
			From:      fromLocation,
			To:        req.Location,
		})
		if err != nil {
			status.Errorf(codes.Internal, "При создании истории перемещения произошла ошибка")
		}

		if normLoc == "IN_EXHIBITION" {
			if strings.TrimSpace(req.GetExhibitionId()) == "" {
				return nil, status.Errorf(codes.InvalidArgument, "При установке локации IN_EXHIBITION обязательно должен быть указан exhibition_id")
			}
			// Проверяем наличие выставки.
			_, err := s.exhibitionClient.GetExhibition(ctx, &exhibitionpb.GetExhibitionRequest{ExhibitionId: req.GetExhibitionId()})
			if err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "Выставка с id %s не найдена", req.GetExhibitionId())
			}
			// Обновляем картину: устанавливаем location = IN_EXHIBITION и exhibition_id.
			p, err = p.Update().
				SetLocation(picture.Location(normLoc)).
				SetExhibitionID(req.GetExhibitionId()).
				Save(ctx)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Ошибка обновления location: %v", err)
			}
		} else {
			// Если новая локация не IN_EXHIBITION, сбрасываем exhibition_id.
			p, err = p.Update().
				SetLocation(picture.Location(normLoc)).
				ClearExhibitionID().
				Save(ctx)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Ошибка обновления location: %v", err)
			}
		}
	}

	// После обновления получаем новое значение location.
	newLoc := p.Location.String()

	// Если старая локация была IN_EXHIBITION, а новая уже не IN_EXHIBITION, обновляем выставку:
	// удаляем ID этой картины из списка PicturesIds у выставки, указанной в старом exhibition_id.
	if oldLoc == "IN_EXHIBITION" && newLoc != "IN_EXHIBITION" && oldExhibitionID != "" {
		exhResp, err := s.exhibitionClient.GetExhibition(ctx, &exhibitionpb.GetExhibitionRequest{ExhibitionId: oldExhibitionID})
		if err == nil {
			picIDStr := strconv.Itoa(p.ID)
			newPics := []string{}
			for _, id := range exhResp.Exhibition.PicturesIds {
				if id != picIDStr {
					newPics = append(newPics, id)
				}
			}
			_, err = s.exhibitionClient.UpdateExhibition(ctx, &exhibitionpb.UpdateExhibitionRequest{
				ExhibitionId: oldExhibitionID,
				PicturesIds:  newPics,
			})
			if err != nil {
				fmt.Printf("Ошибка обновления выставки при удалении id картинки: %v\n", err)
			}
		}
	}

	// Если новая локация стала IN_EXHIBITION, обновляем выставку:
	// добавляем ID текущей картины в список, если его там ещё нет.
	if newLoc == "IN_EXHIBITION" {
		if p.ExhibitionID != "" {
			exhResp, err := s.exhibitionClient.GetExhibition(ctx, &exhibitionpb.GetExhibitionRequest{ExhibitionId: p.ExhibitionID})
			if err == nil {
				picIDStr := strconv.Itoa(p.ID)
				if !contains(exhResp.Exhibition.PicturesIds, picIDStr) {
					newPics := append(exhResp.Exhibition.PicturesIds, picIDStr)
					_, err = s.exhibitionClient.UpdateExhibition(ctx, &exhibitionpb.UpdateExhibitionRequest{
						ExhibitionId: p.ExhibitionID,
						PicturesIds:  newPics,
					})
					if err != nil {
						fmt.Printf("Ошибка обновления выставки при добавлении id картинки: %v\n", err)
					}
				}
			}
		}
	}

	// Загружаем связанные данные авторов и жанров.
	authors, genres, validAuthorsIds, validGenresIds, err := s.LoadAuthorsAndGenres(ctx, p.AuthorsIds, p.GenresIds)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка загрузки авторов и жанров: %v", err)
	}

	updatedPicture := &picturepb.PictureData{
		Id:             strconv.Itoa(p.ID),
		Name:           p.Name,
		DateOfPainting: p.DateOfPainting,
		GenresIds:      validGenresIds,
		AuthorsIds:     validAuthorsIds,
		ExhibitionId:   p.ExhibitionID,
		Cost:           p.Cost,
		Location:       p.Location.String(),
		CreatedAt:      p.CreatedAt.String(),
		Genres:         genres,
		Authors:        authors,
	}

	return &picturepb.UpdatePictureResponse{
		Picture: updatedPicture,
	}, nil
}

// DeletePicture удаляет картину по её picture_id.
func (s *pictureServer) DeletePicture(ctx context.Context, req *picturepb.DeletePictureRequest) (*picturepb.DeletePictureResponse, error) {
	id, err := strconv.Atoi(req.GetPictureId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат picture_id")
	}
	_, err = s.db.Picture.Query().Where(picture.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Картина не найдена")
	}
	if _, err := s.db.Picture.Delete().Where(picture.IDEQ(id)).Exec(ctx); err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка при удалении картины: %v", err)
	}
	return &picturepb.DeletePictureResponse{
		Message: "Картина успешно удалена",
	}, nil
}
