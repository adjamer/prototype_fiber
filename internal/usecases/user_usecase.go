package usecases

import (
	"errors"
	"time"

	"prototype-fiber/internal/domain/entities"
	"prototype-fiber/pkg/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserUseCase struct {
	userRepo   entities.UserRepository
	jwtConfig  config.JWTConfig
}

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Phone     string `json:"phone"`
}

type AuthResponse struct {
	Token string        `json:"token"`
	User  *entities.User `json:"user"`
}

func NewUserUseCase(userRepo entities.UserRepository, jwtConfig config.JWTConfig) *UserUseCase {
	return &UserUseCase{
		userRepo:  userRepo,
		jwtConfig: jwtConfig,
	}
}

func (uc *UserUseCase) Register(req *RegisterRequest) (*AuthResponse, error) {
	// Check if user already exists
	existingUser, _ := uc.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Create new user
	user := &entities.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Role:      entities.RoleCustomer,
		IsActive:  true,
	}

	if err := user.HashPassword(req.Password); err != nil {
		return nil, err
	}

	if err := uc.userRepo.Create(user); err != nil {
		return nil, err
	}

	token, err := uc.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

func (uc *UserUseCase) Login(req *AuthRequest) (*AuthResponse, error) {
	user, err := uc.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !user.CheckPassword(req.Password) {
		return nil, errors.New("invalid credentials")
	}

	if !user.IsActive {
		return nil, errors.New("account is deactivated")
	}

	token, err := uc.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

func (uc *UserUseCase) GetProfile(userID uuid.UUID) (*entities.User, error) {
	return uc.userRepo.GetByID(userID)
}

func (uc *UserUseCase) UpdateProfile(userID uuid.UUID, updates map[string]interface{}) (*entities.User, error) {
	user, err := uc.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if firstName, ok := updates["first_name"].(string); ok {
		user.FirstName = firstName
	}
	if lastName, ok := updates["last_name"].(string); ok {
		user.LastName = lastName
	}
	if phone, ok := updates["phone"].(string); ok {
		user.Phone = phone
	}

	if err := uc.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) generateToken(user *entities.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(uc.jwtConfig.Secret))
}