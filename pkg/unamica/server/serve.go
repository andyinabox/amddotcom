package server

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

func (s *Server) Serve(ctx context.Context) (chan<- string, error) {

	// check that it exists
	_, err := os.Stat(s.cnf.WebRoot)
	if err != nil {
		return nil, err
	}

	// create refresh channel
	refresh := make(chan string)

	// server goroutine
	go func() {
		log.Info(fmt.Sprintf("starting server at%s", s.srv.Addr))
		err := s.srv.ListenAndServe()
		if err != nil {
			log.Fatal("error starting server", "error", err)
			return
		}
	}()

	// message reiver goroutine
	go func() {
		for {
			select {
			case msg := <-refresh:
				log.Debug("recieved refresh event", "message", msg)
				s.cg.Send(msg)
			case <-ctx.Done():
				s.srv.Shutdown(ctx)
				close(refresh)
				return
			}
		}
	}()

	return refresh, nil
}
