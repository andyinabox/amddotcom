package webservice

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

func (s *Service) Serve(ctx context.Context, port int) (err error) {

	// check that it exists
	_, err = os.Stat(s.cnf.WebRoot)
	if err != nil {
		return
	}

	fileServer := http.FileServer(http.Dir(s.cnf.WebRoot))

	http.Handle("/", fileServer)

	fmt.Printf("Server started at http://localhost:%d\n", port)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
