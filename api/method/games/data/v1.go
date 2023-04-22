package data

import (
	"context"
	"time"

	"github.com/go-qbit/rpc"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"

	"github.com/bot-games/localrunner/api/username"
)

type reqV1 struct {
	Id string `json:"id"`
}

type GameV1 struct {
	Ts           time.Time    `json:"ts"`
	Participants []GameUserV1 `json:"participants" field:"game_users"`
	Winner       uint8        `json:"winner"`
	Options      []byte       `json:"options"`
	Ticks        []TickV1     `json:"ticks" field:"game_ticks"`
}

type GameUserV1 struct {
	Id       uint32 `json:"-"`
	Score    uint32 `json:"score"`
	NewScore uint32 `json:"new_score"`
	User     UserV1 `json:"user"`
}

type UserV1 struct {
	Id        uint32 `json:"id"`
	GhLogin   string `json:"gh_login"`
	Name      string `json:"name" field:"fullname"`
	AvatarUrl string `json:"avatar_url"`
}

type TickV1 struct {
	Tick    uint16     `json:"tick"`
	State   []byte     `json:"state"`
	Actions []ActionV1 `json:"actions" field:"game_tick_actions"`
}

type ActionV1 struct {
	Data   []byte `json:"data"`
	UserId uint32 `json:"user" field:"fk_user_id"`
}

var errorsV1 struct {
	InvalidGameId rpc.ErrorFunc `desc:"Invalid game id"`
}

func (m *Method) ErrorsV1() interface{} {
	return &errorsV1
}

func (m *Method) V1(ctx context.Context, r *reqV1) (*GameV1, error) {
	gameUuid, err := uuid.Parse(r.Id)
	if err != nil {
		return nil, errorsV1.InvalidGameId("Invalid game Id")
	}

	g := m.storage.GetGame(gameUuid)
	if g == nil || g.Finished.IsZero() {
		return nil, errorsV1.InvalidGameId("The game is not finished yet")
	}

	participants := make([]GameUserV1, len(g.Info.Uids))
	for i, uid := range g.Info.Uids {
		participants[i] = GameUserV1{
			Id: uid,
			User: UserV1{
				Id:   uid,
				Name: username.UserName(uid),
			},
		}
	}

	bOptions, err := proto.Marshal(g.Info.Options)
	if err != nil {
		return nil, err
	}

	ticks := make([]TickV1, len(g.Ticks))
	for i, t := range g.Ticks {
		ticks[i].Tick = uint16(i)

		ticks[i].Actions = make([]ActionV1, len(t.Actions))
		for j, a := range t.Actions {
			bAction, err := proto.Marshal(a.Action)
			if err != nil {
				return nil, err
			}
			ticks[i].Actions[j] = ActionV1{
				Data:   bAction,
				UserId: a.Uid,
			}
		}

		ticks[i].State, err = proto.Marshal(t.State)
		if err != nil {
			return nil, err
		}
	}

	return &GameV1{
		Ts:           g.Ts,
		Participants: participants,
		Winner:       g.Winner,
		Options:      bOptions,
		Ticks:        ticks,
	}, nil
}
