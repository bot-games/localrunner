package list

import (
	"context"

	"github.com/bot-games/localrunner/storage"
)

type Method struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *Method {
	return &Method{
		storage: storage,
	}
}

func (m *Method) Caption(ctx context.Context) string {
	return `Games list`
}

func (m *Method) Description(ctx context.Context) string {
	return `Returns games list`
}
