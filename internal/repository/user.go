package repository

import (
	"gorm.io/gorm"
	"main/internal/dto"
	"main/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserById(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	r.db.Preload("Role").Find(&users)
	return users, nil
}

func (r *UserRepository) ExportUsers() {

}

func (r *UserRepository) DeleteUserById(id uint) error {
	if err := r.db.Delete(models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) AddUser(newUser *dto.CreateUserRequest) (uint, error) {
	user := models.User{Login: newUser.Login, Password: newUser.Password, RoleId: newUser.RoleId}
	result := r.db.Create(&user)
	if err := result.Error; err != nil {
		return 0, err
	}
	return user.Id, nil
}
