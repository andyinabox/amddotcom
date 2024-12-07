package webservice

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

func (s *Service) Serve(ctx context.Context, port int) (chan<- struct{}, error) {

	ctx, cancel := context.WithCancelCause(ctx)

	refresh := make(chan struct{})

	// check that it exists
	_, err := os.Stat(s.cnf.WebRoot)
	if err != nil {
		return nil, err
	}

	fileServer := http.FileServer(http.Dir(s.cnf.WebRoot))

	http.Handle("/", fileServer)

	fmt.Printf("Server started at http://localhost:%d\n", port)

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
