package repositories

import (
	"telebot-template/internal/database/models"

	"gorm.io/gorm"
)

var UserRepo *UserRepository

func InitUserRepo(db *gorm.DB) {
	UserRepo = &UserRepository{db: db}
}

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Where(&models.User{ID: id}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByTelegramID(telegramID int64) (*models.User, error) {
	var user models.User
	if err := r.db.Where(&models.User{TelegramID: telegramID}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Updates only the non-zero fields of the user object
func (r *UserRepository) Update(id uint, user *models.User) error {
	return r.db.Model(&models.User{ID: id}).Updates(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{ID: id}).Error
}
