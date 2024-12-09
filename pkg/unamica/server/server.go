package server

import (
	"fmt"
	"net/http"

	"github.com/amddotcom/pkg/unamica/server/connectiongroup"
)

type Config struct {
	WebRoot string // path to directory to be served
	Port    int64
}

type Server struct {
	cnf *Config
	srv http.Server
	cg  *connectiongroup.ConnectionGroup
}

func New(cnf *Config) *Server {

	s := &Server{
		cnf: cnf,
		cg:  connectiongroup.New(),
	}

	mux := http.NewServeMux()

	// filesystem server
	mux.Handle("/", http.FileServer(http.Dir(s.cnf.WebRoot)))

	// live reload handler
	mux.HandleFunc("/_livereload", s.getLiveReloadHandler())

	s.srv = http.Server{
		Addr:    fmt.Sprintf("localhost:%d", cnf.Port),
		Handler: mux,
	}

	return s
}
