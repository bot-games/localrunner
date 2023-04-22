//go:generate gostatic2lib -path static/dist/ -package localrunner -out static.go
package localrunner

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"path"
	"strings"

	manager "github.com/bot-games/game-manager"
	"github.com/bot-games/localrunner/api"
	"github.com/bot-games/localrunner/storage"

	"github.com/go-qbit/rpc/openapi"
)

var (
	addr = flag.String("addr", "localhost:10000", "The address of the server")
)

func Start(gameManager *manager.GameManager, storage *storage.Storage) {
	flag.Parse()

	mux := http.NewServeMux()
	gameApi := gameManager.GetGameApi()
	staticHandler := NewHTTPHandler()

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if staticHandler.GetFile(r.URL.Path) == nil {
			_, file := path.Split(r.URL.Path)
			if strings.IndexByte(file, '.') == -1 {
				r.URL.Path = "/"
			}
		}

		staticHandler.ServeHTTP(w, r)
	}))
	mux.Handle("/game/", http.StripPrefix("/game", gameApi))
	mux.Handle("/api/", http.StripPrefix("/api", api.New(storage)))
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		swagger := gameApi.GetSwagger(r.Context())
		swagger.Servers = []openapi.Server{{Url: "/game/"}}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(swagger); err != nil {
			log.Fatal(err)
		}
	})

	mux.Handle("/player/", http.StripPrefix("/player", gameApi.GetPlayerHandler()))

	log.Printf("Starting server on http://%s", *addr)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal(err)
	}
}
