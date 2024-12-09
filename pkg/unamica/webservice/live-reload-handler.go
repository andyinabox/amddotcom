package webservice

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
)

func (*Service) getLiveReloadHandler(refresh <-chan string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		log.Debug("got livereload handshake from client")

		// Set CORS headers to allow all origins. You may want to restrict this to specific origins in a production environment.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Type")
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		for {
			select {
			case refreshMsg := <-refresh:
				log.Debug("send refresh event", "message", refreshMsg)
				fmt.Fprintf(w, "data: %s\n\n", refreshMsg)
				w.(http.Flusher).Flush()
			case <-ctx.Done():
				log.Debug("context cancelled, closing livereload connection")
				return
			}
		}
	}
}
