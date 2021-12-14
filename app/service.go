package main

import (
	"context"
)

type userService struct {
	userRepository *userRepository
}

func NewUserService(userRepository *userRepository) *userService {
	return &userService{userRepository: userRepository}
}

func (service *userService) CreateUser(ctx context.Context, user User) (User, error) {
	return service.userRepository.CreateUser(ctx, user)
}

func (service *userService) GetUser(ctx context.Context, id int64) (User, error) {
	return service.userRepository.GetUser(ctx, id)
}

func (service *userService) UpdateUser(ctx context.Context, id int64, user User) (User, error) {
	return service.userRepository.UpdateUser(ctx, id, user)
}

func (service *userService) DeleteUser(ctx context.Context, id int64) error {
	return service.userRepository.DeleteUser(ctx, id)
}
