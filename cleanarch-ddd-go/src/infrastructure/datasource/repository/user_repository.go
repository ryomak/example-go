package repository

import (
	"context"
	"map-friend/src/domain/repository"
	"map-friend/src/infrastructure/datasource/database"

	"map-friend/src/domain/model"
)

func NewIUserRepository(dbm database.DBM) repository.IUserRepository {
	return &userRepository{dbm}
}

type userRepository struct {
	dbm database.DBM
}

func (u *userRepository) Save(ctx context.Context, user *model.User) (*model.User, error) {
	
}

func (u *userRepository) GetByNames(ctx context.Context, userName, roomName string) (*model.User, error) {
	
}

func (u *userRepository) GetListByRoom(ctx context.Context, roomName string) ([]*model.User, error) {
	
}
