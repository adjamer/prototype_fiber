package repositories

import (
	"prototype-fiber/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entities.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryImpl) GetByID(id uuid.UUID) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Update(user *entities.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.User{}, id).Error
}

func (r *UserRepositoryImpl) List(offset, limit int) ([]*entities.User, error) {
	var users []*entities.User
	err := r.db.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}