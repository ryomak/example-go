package repository

import (
	"context"
	"map-friend/src/domain/repository"
	"map-friend/src/infrastructure/datasource/database"
	"time"

	"map-friend/src/domain/model"
)

func NewIRoomRepository(dbm database.DBM) repository.IRoomRepository {
	return &roomRepository{dbm}
}

type roomRepository struct {
	dbm database.DBM
}

func (r *roomRepository) Save(ctx context.Context, roomName string) (*model.Room, error) {

}

func (r *roomRepository) Get(ctx context.Context, name string) (*model.Room, error) {

}

