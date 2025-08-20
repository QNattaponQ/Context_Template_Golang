package user

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id uint) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *User) error {
	// Check if email already exists
	var existingUser User
	if err := r.db.WithContext(ctx).Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	} else if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("error checking email uniqueness: %w", err)
	}

	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (*User, error) {
	var user User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user with ID %d not found", id)
		}
		return nil, fmt.Errorf("error finding user: %w", err)
	}
	return &user, nil
}
