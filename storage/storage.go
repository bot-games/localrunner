package storage

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"

	manager "github.com/bot-games/game-manager"
)

type Storage struct {
	games    []*Game
	gamesMap map[uuid.UUID]*Game
	mtx      sync.RWMutex
}

type Game struct {
	Ts       time.Time
	Info     *manager.GameInfo
	Winner   uint8
	Timeout  uint8
	Finished time.Time
	Ticks    []manager.Tick
}

func New() *Storage {
	return &Storage{
		gamesMap: map[uuid.UUID]*Game{},
	}
}

func (s *Storage) CreateGame(ctx context.Context, info *manager.GameInfo) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	g := &Game{
		Ts:   time.Now(),
		Info: info,
	}

	s.games = append(s.games, g)
	s.gamesMap[info.Uuid] = g

	return nil
}

func (s *Storage) SaveGame(ctx context.Context, uuid uuid.UUID, winner, timeout uint8, finished time.Time, ticks []manager.Tick) error {
	g := s.GetGame(uuid)
	if g == nil {
		return fmt.Errorf("invalid game")
	}

	g.Winner = winner
	g.Timeout = timeout
	g.Finished = finished
	g.Ticks = ticks

	return nil
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

func (s *Storage) GetGame(uuid uuid.UUID) *Game {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	return s.gamesMap[uuid]
}

func (s *Storage) GetGames() []*Game {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	return s.games
}
