package api

import (
	"net/http"

	"github.com/go-qbit/rpc"
	"github.com/go-qbit/rpc/client/typescript"

	mGamesData "github.com/bot-games/localrunner/api/method/games/data"
	mGamesList "github.com/bot-games/localrunner/api/method/games/list"
	"github.com/bot-games/localrunner/storage"
)

func New(storage *storage.Storage) http.Handler {
	gamesRpc := rpc.New("github.com/bot-games/localrunner/api/method")

	if err := gamesRpc.RegisterMethods(
		mGamesData.New(storage),
		mGamesList.New(storage),
	); err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gamesRpc)

	mux.HandleFunc("/index.ts", typescript.New(gamesRpc, "github.com/bot-games/localrunner/api/method"))

	return mux
}
