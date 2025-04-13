package server

import (
	"context"
	"strconv"
	"strings"

	"github.com/Ostap00034/siproject-beercut-backend/author-service/ent"
	"github.com/Ostap00034/siproject-beercut-backend/author-service/ent/author"
	authorpb "github.com/Ostap00034/siproject-beercut-backend/author-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authorServer struct {
	authorpb.UnimplementedAuthorServiceServer
	db *ent.Client
}

func NewAuthorServer(dbClient *ent.Client) *authorServer {
	return &authorServer{
		db: dbClient,
	}
}

// GetAuthor возвращает данные автора по его author_id.
func (s *authorServer) GetAuthor(ctx context.Context, req *authorpb.GetAuthorRequest) (*authorpb.GetAuthorResponse, error) {
	id, err := strconv.Atoi(req.GetAuthorId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат author_id")
	}

	a, err := s.db.Author.Query().Where(author.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Автор не найден")
	}

	return &authorpb.GetAuthorResponse{
		Author: &authorpb.AuthorData{
			Id:          strconv.Itoa(a.ID),
			FullName:    a.FullName,
			DateOfBirth: a.DateOfBirth,
			DateOfDeath: a.DateOfDeath, // Если поле не установлено, оно может быть пустым
			CreatedAt:   a.CreatedAt.String(),
		},
	}, nil
}

// CreateAuthor создает нового автора. Дата смерти является необязательным полем.
func (s *authorServer) CreateAuthor(ctx context.Context, req *authorpb.CreateAuthorRequest) (*authorpb.CreateAuthorResponse, error) {
	if strings.TrimSpace(req.GetFullName()) == "" {
		errDetail := &authorpb.ErrorResponse{
			Message: "Ошибка валидации",
			Errors:  map[string]string{"full_name": "ФИО обязательно для заполнения"},
		}
		st, err2 := status.New(codes.InvalidArgument, "ФИО отсутствует").WithDetails(errDetail)
		if err2 != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка добавления деталей: %v", err2)
		}
		return nil, st.Err()
	}

	if strings.TrimSpace(req.GetDateOfBirth()) == "" {
		errDetail := &authorpb.ErrorResponse{
			Message: "Ошибка валидации",
			Errors:  map[string]string{"date_of_birth": "Дата рождения обязательна для заполнения"},
		}
		st, err2 := status.New(codes.InvalidArgument, "Дата рождения отсутствует").WithDetails(errDetail)
		if err2 != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка добавления деталей: %v", err2)
		}
		return nil, st.Err()
	}

	creator := s.db.Author.Create().
		SetFullName(req.GetFullName()).
		SetDateOfBirth(req.GetDateOfBirth())

	// Если дата смерти передана (не пустая строка), устанавливаем её.
	if strings.TrimSpace(req.GetDateOfDeath()) != "" {
		creator = creator.SetDateOfDeath(req.GetDateOfDeath())
	}

	a, err := creator.Save(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка создания автора: %v", err)
	}

	return &authorpb.CreateAuthorResponse{
		Author: &authorpb.AuthorData{
			Id:          strconv.Itoa(a.ID),
			FullName:    a.FullName,
			DateOfBirth: a.DateOfBirth,
			DateOfDeath: a.DateOfDeath,
			CreatedAt:   a.CreatedAt.String(),
		},
		Message: "Автор успешно создан",
	}, nil
}

// GetAll возвращает список всех авторов.
func (s *authorServer) GetAll(ctx context.Context, req *authorpb.GetAllRequest) (*authorpb.GetAllResponse, error) {
	authors, err := s.db.Author.Query().All(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка получения авторов: %v", err)
	}

	var authorsList []*authorpb.AuthorData
	for _, a := range authors {
		authorsList = append(authorsList, &authorpb.AuthorData{
			Id:          strconv.Itoa(a.ID),
			FullName:    a.FullName,
			DateOfBirth: a.DateOfBirth,
			DateOfDeath: a.DateOfDeath,
			CreatedAt:   a.CreatedAt.String(),
		})
	}

	return &authorpb.GetAllResponse{
		Authors: authorsList,
	}, nil
}

// UpdateAuthor обновляет данные автора. Если поле DateOfDeath передано, оно обновляется, иначе оставляется прежним.
func (s *authorServer) UpdateAuthor(ctx context.Context, req *authorpb.UpdateAuthorRequest) (*authorpb.UpdateAuthorResponse, error) {
	id, err := strconv.Atoi(req.GetAuthorId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат author_id")
	}

	a, err := s.db.Author.Query().Where(author.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Автор не найден")
	}

	// Обновляем ФИО, если передано.
	if req.GetFullName() != "" {
		a, err = a.Update().SetFullName(req.GetFullName()).Save(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка обновления ФИО: %v", err)
		}
	}
	// Обновляем дату рождения, если передана.
	if req.GetDateOfBirth() != "" {
		a, err = a.Update().SetDateOfBirth(req.GetDateOfBirth()).Save(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка обновления даты рождения: %v", err)
		}
	}
	// Обновляем дату смерти, если передана (иначе не меняем текущее значение).
	if req.GetDateOfDeath() != "" {
		a, err = a.Update().SetDateOfDeath(req.GetDateOfDeath()).Save(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка обновления даты смерти: %v", err)
		}
	}

	updatedAuthor := &authorpb.AuthorData{
		Id:          strconv.Itoa(a.ID),
		FullName:    a.FullName,
		DateOfBirth: a.DateOfBirth,
		DateOfDeath: a.DateOfDeath,
		CreatedAt:   a.CreatedAt.String(),
	}

	return &authorpb.UpdateAuthorResponse{
		Author: updatedAuthor,
	}, nil
}

// DeleteAuthor удаляет автора по его author_id.
func (s *authorServer) DeleteAuthor(ctx context.Context, req *authorpb.DeleteAuthorRequest) (*authorpb.DeleteAuthorResponse, error) {
	id, err := strconv.Atoi(req.GetAuthorId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат author_id")
	}

	_, err = s.db.Author.Query().Where(author.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Автор не найден")
	}

	_, err = s.db.Author.Delete().Where(author.IDEQ(id)).Exec(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка при удалении автора: %v", err)
	}

	return &authorpb.DeleteAuthorResponse{
		Message: "Автор успешно удален",
	}, nil
}
