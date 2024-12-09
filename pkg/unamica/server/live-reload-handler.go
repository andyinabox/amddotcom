package server

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
)

func (s *Server) getLiveReloadHandler(refresh <-chan string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.liveReloadConnections += 1
		defer func() {
			s.liveReloadConnections -= 1
		}()

		ctx := r.Context()

		log.Debugf("got livereload handshake from client, %d connections open", s.liveReloadConnections)

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
