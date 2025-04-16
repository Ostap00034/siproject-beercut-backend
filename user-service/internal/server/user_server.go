package server

import (
	"context"
	"strconv"

	"github.com/Ostap00034/siproject-beercut-backend/user-service/ent"
	"github.com/Ostap00034/siproject-beercut-backend/user-service/ent/user"
	userpb "github.com/Ostap00034/siproject-beercut-backend/user-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userServer struct {
	userpb.UnimplementedUserServiceServer
	db *ent.Client
}

func NewUserServer(dbClient *ent.Client) *userServer {
	return &userServer{db: dbClient}
}

// CreateUser создает нового пользователя.
func (s *userServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	// Здесь можно добавить валидацию входных данных.
	if req.GetEmail() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Email обязателен")
	}
	if req.GetFullName() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "ФИО обязательно для заполнения")
	}
	allowedRoles := map[string]bool{
		"DIRECTOR": true,
		"ADMIN":    true,
		"RESTORER": true,
		"EMPLOYEE": true,
	}
	if !allowedRoles[req.GetRole()] {
		return nil, status.Errorf(codes.InvalidArgument, "Неверное значение роли")
	}

	// Создаем пользователя. (В реальном приложении здесь следует хешировать пароль.)
	u, err := s.db.User.
		Create().
		SetEmail(req.GetEmail()).
		SetPasswordHash(req.GetPassword()).
		SetFullName(req.GetFullName()).
		SetRole(user.Role(req.GetRole())).
		Save(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка создания пользователя: %v", err)
	}

	return &userpb.CreateUserResponse{
		UserId: strconv.Itoa(u.ID),
	}, nil
}

// GetUser возвращает данные пользователя по его user_id.
func (s *userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	id, err := strconv.Atoi(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат user_id")
	}

	u, err := s.db.User.Query().Where(user.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Пользователь не найден")
	}

	return &userpb.GetUserResponse{
		User: &userpb.UserData{
			Id:           strconv.Itoa(u.ID),
			Email:        u.Email,
			FullName:     u.FullName,
			Role:         string(u.Role),
			PasswordHash: u.PasswordHash,
		},
	}, nil
}

// GetUserByEmail возвращает данные пользователя по email.
func (s *userServer) GetUserByEmail(ctx context.Context, req *userpb.GetUserByEmailRequest) (*userpb.GetUserResponse, error) {
	u, err := s.db.User.
		Query().
		Where(user.EmailEQ(req.GetEmail())).
		Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Пользователь с email '%s' не найден", req.GetEmail())
	}

	return &userpb.GetUserResponse{
		User: &userpb.UserData{
			Id:           strconv.Itoa(u.ID),
			Email:        u.Email,
			FullName:     u.FullName,
			Role:         string(u.Role),
			PasswordHash: u.PasswordHash,
		},
	}, nil
}

// UpdateUser обновляет данные пользователя.
func (s *userServer) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	id, err := strconv.Atoi(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Неверный формат user_id")
	}

	u, err := s.db.User.Query().Where(user.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Пользователь не найден")
	}

	if req.GetEmail() != "" {
		u, err = u.Update().SetEmail(req.GetEmail()).Save(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка обновления email: %v", err)
		}
	}
	if req.GetFullName() != "" {
		u, err = u.Update().SetFullName(req.GetFullName()).Save(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка обновления ФИО: %v", err)
		}
	}
	if req.GetRole() != "" {
		allowedRoles := map[string]bool{
			"DIRECTOR": true,
			"ADMIN":    true,
			"RESTORER": true,
			"EMPLOYEE": true,
		}
		if !allowedRoles[req.GetRole()] {
			return nil, status.Errorf(codes.InvalidArgument, "Неверное значение роли")
		}
		u, err = u.Update().SetRole(user.Role(req.GetRole())).Save(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка обновления роли: %v", err)
		}
	}
	if req.GetPassword() != "" {
		if len(req.GetPassword()) < 6 {
			return nil, status.Errorf(codes.InvalidArgument, "Пароль должен содержать минимум 6 символов")
		}
		u, err = u.Update().SetPasswordHash(req.GetPassword()).Save(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка обновления пароля: %v", err)
		}
	}

	updatedUser := &userpb.UserData{
		Id:           strconv.Itoa(u.ID),
		Email:        u.Email,
		FullName:     u.FullName,
		Role:         string(u.Role),
		PasswordHash: u.PasswordHash,
	}

	return &userpb.UpdateUserResponse{
		User: updatedUser,
	}, nil
}

// GetAllUsers возвращает список всех пользователей.
func (s *userServer) GetAllUsers(ctx context.Context, req *userpb.GetAllUsersRequest) (*userpb.GetAllUsersResponse, error) {
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

	total, err := s.db.User.Query().Count(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка получения общего количества жанров: %v", err)
	}

	// Вычисляем количество страниц (округляем в большую сторону).
	totalPages := (total + int(pageSize) - 1) / int(pageSize)

	users, err := s.db.User.Query().Offset(int(offset)).Limit(int(pageSize)).All(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка получения списка пользователей: %v", err)
	}

	var userList []*userpb.UserData
	for _, u := range users {
		userList = append(userList, &userpb.UserData{
			Id:           strconv.Itoa(u.ID),
			Email:        u.Email,
			FullName:     u.FullName,
			Role:         string(u.Role),
			PasswordHash: u.PasswordHash,
		})
	}

	return &userpb.GetAllUsersResponse{
		Users: userList,
		Total:       int32(total),
		TotalPages:  int32(totalPages),
	}, nil
}
