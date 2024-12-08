package webservice

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
)

func (s *Service) Serve(ctx context.Context) (chan<- string, error) {

	// check that it exists
	_, err := os.Stat(s.cnf.WebRoot)
	if err != nil {
		return nil, err
	}

	// create refresh channel
	refresh := make(chan string)

	// filesystem server
	http.Handle("/", http.FileServer(http.Dir(s.cnf.WebRoot)))

	// live reload handler
	http.HandleFunc("/_livereload", s.getLiveReloadHandler(refresh))

	// server goroutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(refresh)
				return
			default:
				log.Info(fmt.Sprintf("starting server at http://localhost:%d", s.cnf.Port))
				err := http.ListenAndServe(fmt.Sprintf(":%d", s.cnf.Port), nil)
				if err != nil {
					log.Fatal("error starting server", "error", err)
					return
				}
			}
		}
	}()

	return refresh, nil
}
