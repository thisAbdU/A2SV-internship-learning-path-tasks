package usecase

import (
	"context"
	"task-management-api/domain/entities"
	"task-management-api/domain/model"
)

type UserUsecase struct {
	userRepository entities.UserRepository
}

func NewUserUsecase(userRepository entities.UserRepository) entities.UserUsecase {
    return &UserUsecase{
        userRepository: userRepository,
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


func (uc *UserUsecase) UpdateUser(ctx context.Context, id string, updatedUser entities.User) error {
	err := uc.userRepository.UpdateUser(context.Background(), id, updatedUser)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUsecase) DeleteUser(ctx context.Context, id string) error {
	err := uc.userRepository.DeleteUser(ctx,id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUsecase) CreateUser(ctx context.Context, newUser model.UserCreate) error {
	_, err := uc.userRepository.CreateUser(ctx, newUser)
	if err != nil {
		return err
	}
	return nil
}

