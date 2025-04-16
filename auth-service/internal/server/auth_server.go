package server

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/Ostap00034/siproject-beercut-backend/auth-service/ent"
	"github.com/Ostap00034/siproject-beercut-backend/auth-service/ent/token"
	authpb "github.com/Ostap00034/siproject-beercut-backend/auth-service/proto"

	// Клиент User Service для вызова методов получения пользователя по email
	userpb "github.com/Ostap00034/siproject-beercut-backend/user-service/proto"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// isValidEmail проверяет корректность email с помощью регулярного выражения.
func isValidEmail(email string) bool {
	const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

type authServer struct {
	authpb.UnimplementedAuthServiceServer
	db         *ent.Client
	userClient userpb.UserServiceClient // Клиент для вызова User Service.
}

// NewAuthServer создаёт новый экземпляр authServer с заданными зависимостями.
func NewAuthServer(dbClient *ent.Client, userClient userpb.UserServiceClient) *authServer {
	return &authServer{
		db:         dbClient,
		userClient: userClient,
	}
}

// RegisterUser регистрирует пользователя: проводит валидацию,
// вызывает User Service для создания пользователя и сохраняет сгенерированный токен в базе Auth Service.
func (s *authServer) RegisterUser(ctx context.Context, req *authpb.RegisterUserRequest) (*authpb.RegisterUserResponse, error) {
	// Валидация email.
	if !isValidEmail(req.GetEmail()) {
		errDetail := &authpb.ErrorResponse{
			Message: "Ошибка валидации",
			Errors:  map[string]string{"email": "Некорректная почта"},
		}
		st, err2 := status.New(codes.InvalidArgument, "Неверный формат email").WithDetails(errDetail)
		if err2 != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка добавления деталей: %v", err2)
		}
		return nil, st.Err()
	}
	// Проверка ФИО.
	if strings.TrimSpace(req.GetFullName()) == "" {
		errDetail := &authpb.ErrorResponse{
			Message: "Ошибка валидации",
			Errors:  map[string]string{"full_name": "ФИО обязательно для заполнения"},
		}
		st, err2 := status.New(codes.InvalidArgument, "ФИО отсутствует").WithDetails(errDetail)
		if err2 != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка добавления деталей: %v", err2)
		}
		return nil, st.Err()
	}
	// Проверка роли.
	allowedRoles := map[string]bool{
		"DIRECTOR": true,
		"ADMIN":    true,
		"RESTORER": true,
		"EMPLOYEE": true,
	}
	if req.GetRole() == "" || !allowedRoles[req.GetRole()] {
		errDetail := &authpb.ErrorResponse{
			Message: "Ошибка валидации",
			Errors: map[string]string{
				"role": "Неверное или отсутствующее значение роли. Допустимые: DIRECTOR, ADMIN, RESTORER, EMPLOYEE",
			},
		}
		st, err2 := status.New(codes.InvalidArgument, "Неверное значение роли").WithDetails(errDetail)
		if err2 != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка добавления деталей: %v", err2)
		}
		return nil, st.Err()
	}
	// Проверка пароля.
	if len(req.GetPassword()) < 6 {
		errDetail := &authpb.ErrorResponse{
			Message: "Ошибка валидации",
			Errors:  map[string]string{"password": "Пароль должен содержать минимум 6 символов"},
		}
		st, err2 := status.New(codes.InvalidArgument, "Пароль слишком короткий").WithDetails(errDetail)
		if err2 != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка добавления деталей: %v", err2)
		}
		return nil, st.Err()
	}

	// Хеширование пароля.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка хеширования: %v", err)
	}

	// Вызов User Service для создания пользователя.
	userResp, err := s.userClient.CreateUser(ctx, &userpb.CreateUserRequest{
		Email:    req.GetEmail(),
		FullName: req.GetFullName(),
		Role:     req.GetRole(),
		// Передаём хешированный пароль (логика хеширования может быть повторена в User Service, если потребуется)
		Password: string(hashedPassword),
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return nil, status.Errorf(codes.AlreadyExists, "Пользователь с такой почтой уже существует")
		}
		return nil, status.Errorf(codes.Internal, "Ошибка создания пользователя в User Service: %v", err)
	}

	userID := userResp.GetUserId()

	// Генерация PAT токена.
	tokenStr := uuid.New().String()
	expiresAt := time.Now().Add(24 * time.Hour)
	_, err = s.db.Token.
		Create().
		SetToken(tokenStr).
		SetExpiresAt(expiresAt).
		SetUserID(userID).
		Save(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка создания токена: %v", err)
	}

	return &authpb.RegisterUserResponse{
		UserId:  userID,
		Message: "Пользователь успешно зарегистрирован",
	}, nil
}

// Login осуществляет аутентификацию, обращаясь к User Service для получения данных о пользователе по email,
// затем сравнивает пароль и, если все ок, генерирует PAT токен, сохраняет его и возвращает.
func (s *authServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	// Проверка формата email.
	if !isValidEmail(req.GetEmail()) {
		errDetail := &authpb.ErrorResponse{
			Message: "Ошибка валидации",
			Errors:  map[string]string{"email": "Некорректная почта"},
		}
		st, err2 := status.New(codes.InvalidArgument, "Неверный формат email").WithDetails(errDetail)
		if err2 != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка добавления деталей: %v", err2)
		}
		return nil, st.Err()
	}

	fmt.Println(req)

	// Получаем данные пользователя через User Service по email.
	userResp, err := s.userClient.GetUserByEmail(ctx, &userpb.GetUserByEmailRequest{
		Email: req.GetEmail(),
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Пользователь не найден")
	}

	// Сравнение паролей: предполагается, что userResp.User содержит поле PasswordHash.
	if err := bcrypt.CompareHashAndPassword([]byte(userResp.User.PasswordHash), []byte(req.GetPassword())); err != nil {
		errDetail := &authpb.ErrorResponse{
			Message: "Ошибка аутентификации",
			Errors:  map[string]string{"password": "Неправильный пароль"},
		}
		st, err2 := status.New(codes.Unauthenticated, "Неправильный пароль").WithDetails(errDetail)
		if err2 != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка добавления деталей: %v", err2)
		}
		return nil, st.Err()
	}

	// Генерация PAT токена.
	tokenStr := uuid.New().String()
	expiresAt := time.Now().Add(24 * time.Hour)
	_, err = s.db.Token.
		Create().
		SetToken(tokenStr).
		SetExpiresAt(expiresAt).
		// Здесь сохраняем userID, полученный от User Service.
		SetUserID(userResp.User.Id).
		Save(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка создания токена: %v", err)
	}

	return &authpb.LoginResponse{
		Token:     tokenStr,
		ExpiresAt: expiresAt.Format(time.RFC3339),
		UserId:    userResp.User.Id,
		Role:      userResp.User.Role,
	}, nil
}

// ValidateToken проверяет действительность токена.
func (s *authServer) ValidateToken(ctx context.Context, req *authpb.ValidateTokenRequest) (*authpb.ValidateTokenResponse, error) {
	t, err := s.db.Token.
		Query().
		Where(token.TokenEQ(req.GetToken())).
		Only(ctx)
	if err != nil {
		return &authpb.ValidateTokenResponse{Valid: false}, status.Errorf(codes.NotFound, "Токен не найден")
	}

	if time.Now().After(t.ExpiresAt) {
		return &authpb.ValidateTokenResponse{Valid: false}, status.Errorf(codes.Unauthenticated, "Токен просрочен")
	}

	userResp, err := s.userClient.GetUser(ctx, &userpb.GetUserRequest{
		UserId: t.UserID,
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Пользователь не найден")
	}

	return &authpb.ValidateTokenResponse{
		Valid:  true,
		UserId: t.UserID,
		Role:   userResp.User.Role,
	}, nil
}

// DeleteToken удаляет токен.
func (s *authServer) DeleteToken(ctx context.Context, req *authpb.DeleteTokenRequest) (*authpb.DeleteTokenResponse, error) {
	_, err := s.db.Token.
		Delete().
		Where(token.TokenEQ(req.GetToken())).
		Exec(ctx)
	if err != nil {
		return &authpb.DeleteTokenResponse{Success: false}, status.Errorf(codes.Internal, "Ошибка удаления токена: %v", err)
	}
	return &authpb.DeleteTokenResponse{Success: true}, nil
}
