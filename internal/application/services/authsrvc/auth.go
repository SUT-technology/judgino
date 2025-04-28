package authsrvc

import (
	"context"
	"fmt"
	"time"

	"github.com/SUT-technology/judgino/internal/domain/dto"
	"github.com/SUT-technology/judgino/internal/domain/entity"
	"github.com/SUT-technology/judgino/internal/domain/model"
	"github.com/SUT-technology/judgino/internal/domain/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthSrvc struct {
	db         repository.Pool
	secret_key string
}

func NewAuthSrvc(db repository.Pool, secret_key string) AuthSrvc {
	return AuthSrvc{
		db:         db,
		secret_key: secret_key,
	}
}

func generateToken(userID uint, isAdmin bool, secret_key string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours
	claims := &model.JWTClaims{
		UserID:  int64(userID),
		IsAdmin: isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret_key))
}
func (c AuthSrvc) Login(ctx context.Context, loginRequest dto.LoginRequest) (dto.AuthResponse, error) {

	var (
		user *entity.User
		err  error
	)

	queryFuncFindUser := func(r *repository.Repo) error {
		user, err = r.Tables.Users.GetUserByUsername(ctx, loginRequest.Username)
		if err != nil {
			return fmt.Errorf("find user by username: %w", err)
		}
		return nil
	}
	err = c.db.Query(ctx, queryFuncFindUser)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return dto.AuthResponse{Error: model.InvalidPassword}, err
	}

	// Generate token
	token, err := generateToken(user.ID, user.IsAdmin(), c.secret_key)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	return dto.AuthResponse{Token: token}, nil
}
func (c AuthSrvc) Signup(ctx context.Context, signupRequest dto.SignupRequest) (dto.AuthResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("generating password: %w", err)
	}

	user := &entity.User{
		FirstName: signupRequest.FirstName,
		LastName:  signupRequest.LastName,
		Phone:     signupRequest.Phone,
		Email:     signupRequest.Email,
		Username:  signupRequest.Username,
		Role:      "user",
		Password:  string(hashedPassword),
	}

	queryFuncFindUser := func(r *repository.Repo) error {
		err = r.Tables.Users.CreateUser(ctx, user)
		if err != nil {
			return fmt.Errorf("create user: %w", err)
		}
		return nil
	}
	err = c.db.Query(ctx, queryFuncFindUser)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	// Generate token
	token, err := generateToken(user.ID, user.IsAdmin(), c.secret_key)
	if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("generating token: %w : %v", err, c.secret_key)
	}

	return dto.AuthResponse{Token: token}, nil
}
