// internal/server/auth_server.go
package server

import (
	"context"
	"errors"
	"time"

	"github.com/Ostap00034/siproject-beercut-backend/auth-service/ent"
	"github.com/Ostap00034/siproject-beercut-backend/auth-service/ent/token"
	"github.com/Ostap00034/siproject-beercut-backend/auth-service/ent/user"
	"github.com/Ostap00034/siproject-beercut-backend/auth-service/proto"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authServer struct {
	proto.UnimplementedAuthServiceServer
	db *ent.Client
}

// NewAuthServer создаёт новый экземпляр authServer.
func NewAuthServer(dbClient *ent.Client) *authServer {
	return &authServer{db: dbClient}
}

// RegisterUser: регистрация нового пользователя.
func (s *authServer) RegisterUser(ctx context.Context, req *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error) {
	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u, err := s.db.User.
		Create().
		SetEmail(req.GetEmail()).
		SetPasswordHash(string(hashedPassword)).
		SetFullName(req.GetFullName()).
		SetRole(req.GetRole()). // Обратите внимание: здесь проверка на допустимые значения role осуществляется на уровне схемы
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &proto.RegisterUserResponse{
		UserId:  u.ID.String(),
		Message: "User registered successfully",
	}, nil
}

// Login: аутентификация по email и паролю, генерация PAT токена.
func (s *authServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	u, err := s.db.User.
		Query().
		Where(user.EmailEQ(req.GetEmail())). // поиск по email
		Only(ctx)
	if err != nil {
		return nil, err
	}

	// Сравнение пароля
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.GetPassword())); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Генерация PAT токена
	tokenStr := uuid.New().String()
	expiresAt := time.Now().Add(24 * time.Hour)

	// Создание записи токена
	_, err = s.db.Token.
		Create().
		SetToken(tokenStr).
		SetRole(u.Role).
		SetExpiresAt(expiresAt).
		SetUser(u).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{
		Token:     tokenStr,
		ExpiresAt: expiresAt.Format(time.RFC3339),
		UserId:    u.ID.String(),
		Role:      u.Role,
	}, nil
}

// ValidateToken: проверка валидности PAT токена.
func (s *authServer) ValidateToken(ctx context.Context, req *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error) {
	t, err := s.db.Token.
		Query().
		Where(token.TokenEQ(req.GetToken())).
		Only(ctx)
	if err != nil {
		return &proto.ValidateTokenResponse{Valid: false}, err
	}

	if time.Now().After(t.ExpiresAt) {
		return &proto.ValidateTokenResponse{Valid: false}, nil
	}

	return &proto.ValidateTokenResponse{
		Valid:  true,
		UserId: t.Edges.User.ID.String(),
		Role:   t.Role,
	}, nil
}

// DeleteToken: отзыв (удаление) PAT токена.
func (s *authServer) DeleteToken(ctx context.Context, req *proto.DeleteTokenRequest) (*proto.DeleteTokenResponse, error) {
	err := s.db.Token.
		Delete().
		Where(token.TokenEQ(req.GetToken())).
		Exec(ctx)
	if err != nil {
		return &proto.DeleteTokenResponse{Success: false}, err
	}
	return &proto.DeleteTokenResponse{Success: true}, nil
}
