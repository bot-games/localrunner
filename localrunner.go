//go:generate gostatic2lib -path swagger/ -package localrunner -out swagger.go
package localrunner

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/go-qbit/rpc/openapi"
)

var (
	addr = flag.String("addr", "localhost:10000", "The address of the server")
)

type GameRpc interface {
	http.Handler
	GetSwagger(ctx context.Context) *openapi.OpenApi
}

func Start(gameApi GameRpc) {
	flag.Parse()

	mux := http.NewServeMux()

	mux.Handle("/", NewHTTPHandler())
	mux.Handle("/api/", http.StripPrefix("/api", gameApi))
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		swagger := gameApi.GetSwagger(r.Context())
		swagger.Servers = []openapi.Server{{Url: "/api/"}}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(swagger); err != nil {
			log.Fatal(err)
		}
	})

	log.Printf("Starting server on http://%s", *addr)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal(err)
	}
}
