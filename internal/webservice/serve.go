package webservice

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
)

func (s *Service) Serve(ctx context.Context, port int) (chan<- string, error) {

	ctx, cancel := context.WithCancelCause(ctx)

	refresh := make(chan string)

	// check that it exists
	_, err := os.Stat(s.cnf.WebRoot)
	if err != nil {
		return nil, err
	}

	fileServer := http.FileServer(http.Dir(s.cnf.WebRoot))

	http.Handle("/", fileServer)

	// events handler
	http.HandleFunc("/_livereload", func(w http.ResponseWriter, r *http.Request) {
		log.Debug("got livereload handshake from client")

		// Set CORS headers to allow all origins. You may want to restrict this to specific origins in a production environment.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Type")
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Simulate sending events (you can replace this with real data)
		for refreshMsg := range refresh {
			log.Debug("send refresh event", "message", refreshMsg)
			fmt.Fprintf(w, "data: %s\n\n", refreshMsg)
			w.(http.Flusher).Flush()
		}

	})

	log.Info(fmt.Sprintf("server started at http://localhost:%d", port))

	cleanup := func() {
		close(refresh)
	}

	// server goroutine
	// will cancel context if there's an error starting the server
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
				if err != nil {
					cancel(err)
					return
				}
			}
		}
	}()

	// refresh goroutine
	// this one is responsible for cleanup if the context is cancelled
	go func() {
		for {
			select {
			case <-refresh:
				// send refresh SSR
			case <-ctx.Done():
				cleanup()
				return
			}
		}
	}()

	return refresh, nil
}
