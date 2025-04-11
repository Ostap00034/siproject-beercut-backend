// internal/server/auth_server.go
package server

import (
	"context"
	"regexp"
	"strconv"
	"time"

	"github.com/Ostap00034/siproject-beercut-backend/auth-service/ent"
	"github.com/Ostap00034/siproject-beercut-backend/auth-service/ent/token"
	"github.com/Ostap00034/siproject-beercut-backend/auth-service/ent/user"
	authpb "github.com/Ostap00034/siproject-beercut-backend/auth-service/proto"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// isValidEmail проверяет корректность адреса электронной почты с помощью регулярного выражения.
func isValidEmail(email string) bool {
	const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

type authServer struct {
	authpb.UnimplementedAuthServiceServer
	db *ent.Client
}

// NewAuthServer создаёт новый экземпляр authServer с переданным ent-клиентом.
func NewAuthServer(dbClient *ent.Client) *authServer {
	return &authServer{db: dbClient}
}

// RegisterUser регистрирует нового пользователя после проверки входных данных.
func (s *authServer) RegisterUser(ctx context.Context, req *authpb.RegisterUserRequest) (*authpb.RegisterUserResponse, error) {
	// Проверка формата email.
	if !isValidEmail(req.GetEmail()) {
		errDetail := &authpb.ErrorResponse{
			Message: "Ошибка валидации",
			Errors: map[string]string{
				"email": "Некорректная почта",
			},
		}
		st, err2 := status.New(codes.InvalidArgument, "Неверный формат email").WithDetails(errDetail)
		if err2 != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка добавления деталей ошибки: %v", err2)
		}
		return nil, st.Err()
	}

	// Проверка наличия ФИО.
	if req.GetFullName() == "" {
		errDetail := &authpb.ErrorResponse{
			Message: "Ошибка валидации",
			Errors: map[string]string{
				"full_name": "ФИО обязательно для заполнения",
			},
		}
		st, err2 := status.New(codes.InvalidArgument, "ФИО отсутствует").WithDetails(errDetail)
		if err2 != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка добавления деталей ошибки: %v", err2)
		}
		return nil, st.Err()
	}

	// Проверка наличия и корректности роли.
	allowedRoles := map[string]bool{"DIRECTOR": true, "ADMIN": true, "RESTORER": true, "EMPLOYEE": true}
	if req.GetRole() == "" || !allowedRoles[req.GetRole()] {
		errDetail := &authpb.ErrorResponse{
			Message: "Ошибка валидации",
			Errors: map[string]string{
				"role": "Неверное или отсутствующее значение роли. Допустимые значения: DIRECTOR, ADMIN, RESTORER, EMPLOYEE",
			},
		}
		st, err2 := status.New(codes.InvalidArgument, "Неверное значение роли").WithDetails(errDetail)
		if err2 != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка добавления деталей ошибки: %v", err2)
		}
		return nil, st.Err()
	}

	// Проверка минимальной длины пароля.
	if len(req.GetPassword()) < 6 {
		errDetail := &authpb.ErrorResponse{
			Message: "Ошибка валидации",
			Errors: map[string]string{
				"password": "Пароль должен содержать минимум 6 символов",
			},
		}
		st, err2 := status.New(codes.InvalidArgument, "Пароль слишком короткий").WithDetails(errDetail)
		if err2 != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка добавления деталей ошибки: %v", err2)
		}
		return nil, st.Err()
	}

	// Хеширование пароля.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка хеширования пароля: %v", err)
	}

	u, err := s.db.User.
		Create().
		SetEmail(req.GetEmail()).
		SetPasswordHash(string(hashedPassword)).
		SetFullName(req.GetFullName()).
		SetRole(user.Role(req.GetRole())). // Приведение входящей строки к типу user.Role.
		Save(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка создания пользователя: %v", err)
	}

	return &authpb.RegisterUserResponse{
		UserId:  strconv.Itoa(u.ID),
		Message: "Пользователь успешно зарегистрирован",
	}, nil
}

// Login осуществляет аутентификацию пользователя по email и паролю, генерирует и сохраняет PAT токен.
func (s *authServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	// Проверка формата email.
	if !isValidEmail(req.GetEmail()) {
		errDetail := &authpb.ErrorResponse{
			Message: "Ошибка валидации",
			Errors: map[string]string{
				"email": "Некорректная почта",
			},
		}
		st, err2 := status.New(codes.InvalidArgument, "Неверный формат email").WithDetails(errDetail)
		if err2 != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка добавления деталей ошибки: %v", err2)
		}
		return nil, st.Err()
	}

	// Поиск пользователя по email.
	u, err := s.db.User.
		Query().
		Where(user.EmailEQ(req.GetEmail())).
		Only(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Пользователь не найден")
	}

	// Сравнение паролей.
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.GetPassword())); err != nil {
		errDetail := &authpb.ErrorResponse{
			Message: "Ошибка аутентификации",
			Errors: map[string]string{
				"password": "Неправильный пароль",
			},
		}
		st, err2 := status.New(codes.Unauthenticated, "Неправильный пароль").WithDetails(errDetail)
		if err2 != nil {
			return nil, status.Errorf(codes.Internal, "Ошибка добавления деталей ошибки: %v", err2)
		}
		return nil, st.Err()
	}

	// Генерация PAT токена.
	tokenStr := uuid.New().String()
	expiresAt := time.Now().Add(24 * time.Hour)

	_, err = s.db.Token.
		Create().
		SetToken(tokenStr).
		SetRole(string(u.Role)). // Приведение enum к строке.
		SetExpiresAt(expiresAt).
		SetUser(u).
		Save(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Ошибка создания токена: %v", err)
	}

	return &authpb.LoginResponse{
		Token:     tokenStr,
		ExpiresAt: expiresAt.Format(time.RFC3339),
		UserId:    strconv.Itoa(u.ID),
		Role:      string(u.Role),
	}, nil
}

// ValidateToken проверяет действительность PAT токена и возвращает данные пользователя.
func (s *authServer) ValidateToken(ctx context.Context, req *authpb.ValidateTokenRequest) (*authpb.ValidateTokenResponse, error) {
	t, err := s.db.Token.
		Query().
		Where(token.TokenEQ(req.GetToken())).
		WithUser(). // Подгружает связанные данные пользователя.
		Only(ctx)
	if err != nil {
		return &authpb.ValidateTokenResponse{Valid: false}, status.Errorf(codes.NotFound, "Токен не найден")
	}

	if time.Now().After(t.ExpiresAt) {
		return &authpb.ValidateTokenResponse{Valid: false}, status.Errorf(codes.Unauthenticated, "Токен просрочен")
	}

	return &authpb.ValidateTokenResponse{
		Valid:  true,
		UserId: strconv.Itoa(t.Edges.User.ID),
		Role:   string(t.Role),
	}, nil
}

// DeleteToken отзваляет (удаляет) PAT токен.
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
