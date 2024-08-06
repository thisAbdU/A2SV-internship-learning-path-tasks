package usecase

import (
	"context"
	"task-management-api/config"
	"task-management-api/domain/entities"
)

type UserUsecase struct {
	userRepository entities.UserRepository
	Environment   *config.Environment
}

func NewUserUsecase(environment *config.Environment, userRepository entities.UserRepository) *UserUsecase {
    return &UserUsecase{
        userRepository: userRepository,
        Environment:    environment,
    }
}

func (uc *UserUsecase) GetUsers(ctx context.Context, param string) ([]*entities.User, error) {
    users, err := uc.userRepository.GetUser(ctx, param)
    if err != nil {
        return nil, err
    }
    return users, nil
}


func (uc *UserUsecase) GetUserByID(ctx context.Context, id string) (*entities.User, error) {
    user, err := uc.userRepository.GetUserByID(ctx, id)
    if err != nil {
        return nil, err
    }
    return user, nil
}


func (uc *UserUsecase) UpdateUser(id string, updatedUser entities.User) error {
	err := uc.userRepository.UpdateUser(context.Background(), id, updatedUser)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUsecase) DeleteUser(id string) error {
	err := uc.userRepository.DeleteUser(context.Background(),id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUsecase) CreateUser(newUser entities.User) error {
	_, err := uc.userRepository.CreateUser(context.Background(), newUser)
	if err != nil {
		return err
	}
	return nil
}

