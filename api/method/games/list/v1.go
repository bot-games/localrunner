package list

import (
	"context"
	"time"

	"github.com/bot-games/localrunner/api/username"
)

type reqV1 struct{}

type GameV1 struct {
	Id           string     `json:"id"`
	Debug        bool       `json:"debug,omitempty"`
	Ts           time.Time  `json:"ts"`
	Finished     *time.Time `json:"finished,omitempty"`
	Participants []UserV1   `json:"participants"`
	Winner       uint8      `json:"winner"`
	Timeout      uint8      `json:"timeout"`
}

type UserV1 struct {
	Name string `json:"name"`
}

func (m *Method) V1(ctx context.Context, r *reqV1) ([]GameV1, error) {
	games := m.storage.GetGames()
	res := make([]GameV1, len(games))

	for i, g := range games {
		res[i] = GameV1{
			Id:           g.Info.Uuid.String(),
			Debug:        g.Info.Debug,
			Ts:           g.Ts,
			Participants: make([]UserV1, len(g.Info.Uids)),
			Winner:       g.Winner,
			Timeout:      g.Timeout,
		}

		if !g.Finished.IsZero() {
			res[i].Finished = &g.Finished
		}
		for j, uid := range g.Info.Uids {
			res[i].Participants[j].Name = username.UserName(uid)
		}
	}

	return res, nil
}
