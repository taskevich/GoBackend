package service

import (
	"main/internal/dto"
	"main/internal/models"
	"main/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserById(id uint) (*models.User, error) {
	return s.userRepo.GetUserById(id)
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.userRepo.GetUsers()
}

func (s *UserService) DeleteUserById(id uint) error {
	return s.userRepo.DeleteUserById(id)
}

func (s *UserService) AddUser(newUser *dto.CreateUserRequest) (uint, error) {
	return s.userRepo.AddUser(newUser)
}
