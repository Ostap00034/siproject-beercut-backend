package server

import (
	"context"
	"strconv"
	"strings"

	"github.com/Ostap00034/siproject-beercut-backend/genre-service/ent"
	"github.com/Ostap00034/siproject-beercut-backend/genre-service/ent/genre"
	genrepb "github.com/Ostap00034/siproject-beercut-backend/genre-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type genreServer struct {
	genrepb.UnimplementedGenreServiceServer
	db *ent.Client
}

func NewGenreServer(dbClient *ent.Client) *genreServer {
	return &genreServer{
		db: dbClient,
	}
}

func (s *genreServer) GetGenre(ctx context.Context, req *genrepb.GetGenreRequest) (*genrepb.GetGenreResponse, error) {
	id, err := strconv.Atoi(req.GetGenreId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат genre_id")
	}

	g, err := s.db.Genre.Query().Where(genre.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Жанр не найден")
	}

	return &genrepb.GetGenreResponse{
		Genre: &genrepb.GenreData{
			Id:          strconv.Itoa(g.ID),
			Name:        g.Name,
			Description: g.Description,
			CreatedAt:   g.CreatedAt.String(),
		},
	}, nil
}

func (s *genreServer) CreateGenre(ctx context.Context, req *genrepb.CreateGenreRequest) (*genrepb.CreateGenreResponse, error) {
	if strings.TrimSpace(req.GetName()) == "" {
		errDetail := &genrepb.ErrorResponse{
			Message: "Ошибка валидации",
			Errors:  map[string]string{"name": "Название обязательно для заполнения"},
		}
		st, err2 := status.New(codes.InvalidArgument, "Название отсутствует").WithDetails(errDetail)
		if err2 != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка добавления деталей: %v", err2)
		}
		return nil, st.Err()
	}

	genre, err := s.db.Genre.
		Create().
		SetName(req.GetName()).
		SetDescription(req.GetDescription()).
		Save(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, status.Errorf(codes.InvalidArgument, "Жанр с таким названием уже существует")
		}
		return nil, status.Errorf(codes.Internal, "Ошибка создания жанра: %v", err)
	}

	return &genrepb.CreateGenreResponse{
		Genre: &genrepb.GenreData{
			Id:          strconv.Itoa(genre.ID),
			Name:        genre.Name,
			Description: genre.Description,
			CreatedAt:   genre.CreatedAt.String(),
		},
		Message: "Жанр успешно создан",
	}, nil
}

// GetAll реализует пагинацию на основе параметров page_number и page_size.
// Если параметры не заданы (<=0), используются значения по умолчанию.
func (s *genreServer) GetAll(ctx context.Context, req *genrepb.GetAllRequest) (*genrepb.GetAllResponse, error) {
	// Получаем пагинационные параметры. Если они не заданы, используем значения по умолчанию.
	pageNumber := req.GetPageNumber()
	pageSize := req.GetPageSize()
	if pageNumber <= 0 {
		pageNumber = 1
	}
	if pageSize <= 0 {
		pageSize = 10 // значение по умолчанию
	}
	offset := (pageNumber - 1) * pageSize

	genres, err := s.db.Genre.Query().Offset(int(offset)).Limit(int(pageSize)).All(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка получения жанров: %v", err)
	}

	total, err := s.db.Genre.Query().Count(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка получения общего количества жанров: %v", err)
	}

	// Вычисляем количество страниц (округляем в большую сторону).
	totalPages := (total + int(pageSize) - 1) / int(pageSize)

	var genresList []*genrepb.GenreData
	for _, g := range genres {
		genresList = append(genresList, &genrepb.GenreData{
			Id:          strconv.Itoa(g.ID),
			Name:        g.Name,
			Description: g.Description,
			CreatedAt:   g.CreatedAt.String(),
		})
	}

	return &genrepb.GetAllResponse{
		Genres:     genresList,
		Total:      int32(total),
		TotalPages: int32(totalPages),
	}, nil
}

func (s *genreServer) UpdateGenre(ctx context.Context, req *genrepb.UpdateGenreRequest) (*genrepb.UpdateGenreResponse, error) {
	id, err := strconv.Atoi(req.GetGenreId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат genre_id")
	}

	g, err := s.db.Genre.Query().Where(genre.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Жанр не найден")
	}

	if req.GetName() != "" {
		g, err = g.Update().SetName(req.GetName()).Save(ctx)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				return nil, status.Errorf(codes.InvalidArgument, "Жанр с таким названием уже существует")
			}
			return nil, status.Errorf(codes.Internal, "Ошибка обновления названия: %v", err)
		}
	}

	if req.GetDescription() != "" {
		g, err = g.Update().SetDescription(req.GetDescription()).Save(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка обновления описания: %v", err)
		}
	}

	updatedGenre := &genrepb.GenreData{
		Id:          strconv.Itoa(g.ID),
		Name:        g.Name,
		Description: g.Description,
		CreatedAt:   g.CreatedAt.String(),
	}

	return &genrepb.UpdateGenreResponse{
		Genre: updatedGenre,
	}, nil
}

func (s *genreServer) DeleteGenre(ctx context.Context, req *genrepb.DeleteGenreRequest) (*genrepb.DeleteGenreResponse, error) {
	id, err := strconv.Atoi(req.GetGenreId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат genre_id")
	}

	_, err = s.db.Genre.Query().Where(genre.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Жанр не найден")
	}

	_, err = s.db.Genre.Delete().Where(genre.IDEQ(id)).Exec(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка при удалении жанра")
	}

	return &genrepb.DeleteGenreResponse{
		Message: "Жанр успешно удален",
	}, nil
}
