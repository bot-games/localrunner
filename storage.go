package localrunner

import (
	"context"
	"strconv"
	"time"

	"github.com/google/uuid"

	manager "github.com/bot-games/game-manager"
)

type Storage struct{}

func (s *Storage) CreateGame(ctx context.Context, info *manager.GameInfo) error {
	panic("implement me")
}

func (s *Storage) SaveGame(ctx context.Context, uuid uuid.UUID, winner, timeout uint8, finished time.Time, ticks []manager.Tick) error {
	panic("implement me")
}

func (s *Storage) GetGameId(ctx context.Context, uuid uuid.UUID) (uint32, error) {
	return 0, nil
}

func (s *Storage) GetUserByToken(ctx context.Context, token string) (*manager.User, error) {
	id, err := strconv.ParseUint(token, 10, 32)
	if err != nil {
		return nil, err
	}

	return &manager.User{
		Id:    uint32(id),
		Score: 0,
	}, nil
}