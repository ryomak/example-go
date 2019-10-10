package usecase

import (
	"context"
	"map-friend/src/domain/model"
	"map-friend/src/domain/repository"
)

type IRoomUseCase interface {
	GetRoom(context.Context, string, *model.User) (*model.Room, error)
}

type roomUseCase struct {
	transaction repository.ITransactionRepository
	roomRepo    repository.IRoomRepository
	userRepo    repository.IUserRepository
}

func NewIRoomUseCase(
	t repository.ITransactionRepository,
	room repository.IRoomRepository,
	user repository.IUserRepository) IRoomUseCase {
	return &roomUseCase{t, room, user}
}

func (r *roomUseCase) GetRoom(ctx context.Context, roomName string, user *model.User) (*model.Room, error) {

	ctx, err := r.transaction.Begin(ctx)
	if err != nil {
		r.transaction.Rollback(ctx)
		return nil, err
	}

	room, err := r.roomRepo.Save(ctx, roomName)
	if err != nil {
		r.transaction.Rollback(ctx)
		return nil, err
	}

	if _, err := r.transaction.Commit(ctx); err != nil {
		r.transaction.Rollback(ctx)
		return nil, err
	}
	return room, nil
}

